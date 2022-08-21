package ex

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"io"
	"io/fs"
	"net/http"
	"os"
	"path"
	"sync"

	"github.com/jfyne/live"
	"github.com/pelletier/go-toml/v2"
	"github.com/runar-rkmedia/gotally/generated"
	tally "github.com/runar-rkmedia/gotally/tallylogic"
)

// Model of our thermostat.
type GameModel struct {
	Hints             map[string]tally.Hint
	HintButtonCounter int
	tally.Game
}

type stupidcache struct {
	games map[string]*GameModel
	sync.RWMutex
}

func (c *stupidcache) GetGame(s string) *GameModel {
	c.RLock()
	defer c.RUnlock()
	return c.games[s]
}
func (c *stupidcache) SetGame(s string, game *GameModel) {
	c.Lock()
	defer c.Unlock()
	c.games[s] = game
}

var (
	cache stupidcache = stupidcache{
		games: map[string]*GameModel{},
	}
	cookieStore = live.NewCookieStore("cookie", []byte("eeeee"))
)

func getSesssionId(s live.Socket) string {
	if session, ok := s.Session()["_lsid"]; ok {
		return session.(string)
	}
	return ""
}

func NewGameModel(mode tally.GameMode, template *tally.GameTemplate) *GameModel {
	m := GameModel{}
	game, err := tally.NewGame(mode, template)
	if err != nil {
		panic("Starting game failed: " + err.Error())
	}
	m.Game = game
	return &m

}

func NewThermoModel(s live.Socket) *GameModel {
	m, ok := s.Assigns().(*GameModel)
	if !ok {
		sessionID := getSesssionId(s)
		ex := cache.GetGame(sessionID)
		if ex != nil {
			return ex
		}
		mode := tally.GameModeTemplate
		if len(generatedTemplates) > 0 {
			mode := tally.GameModeRandomChallenge
			i := tally.NewRandomizer().Intn(len(generatedTemplates))
			m = NewGameModel(mode, &generatedTemplates[i])
		} else {
			m = NewGameModel(mode, &tally.ChallengeGames[0])
		}
		cache.SetGame(sessionID, m)
	}
	return m
}

// thermoMount initialises the thermostat state. Data returned in the mount function will
// automatically be assigned to the socket.
func thermoMount(ctx context.Context, s live.Socket) (interface{}, error) {
	return NewThermoModel(s), nil
}

// swipe on the temp down event, decrease the thermostat temperature by .1 C.
func swipe(ctx context.Context, s live.Socket, p live.Params) (interface{}, error) {
	model := NewThermoModel(s)
	changed := model.Swipe(tally.SwipeDirection(p.String("dir")))
	if changed {
		model.Hints = map[string]tally.Hint{}
	}
	return model, nil
}
func selectCell(ctx context.Context, s live.Socket, p live.Params) (interface{}, error) {
	model := NewThermoModel(s)
	index := p.Int("i")
	selection := model.SelectedCells()
	if len(selection) > 0 && selection[len(selection)-1] == index {
		ok := model.EvaluateSelection()
		if ok {
			model.Hints = map[string]tally.Hint{}
		}
	} else {
		model.SelectCell(index)
	}
	return model, nil
}
func newGame(ctx context.Context, s live.Socket, p live.Params) (interface{}, error) {
	mode := tally.GameMode(p.Int("mode"))
	var template *tally.GameTemplate
	if mode == tally.GameModeRandomChallenge {
		if len(generatedTemplates) == 0 {
			// panic("no games")
			return nil, fmt.Errorf("could not find any generated games")
		}
		i := tally.NewRandomizer().Intn(len(generatedTemplates) - 1)
		template = &generatedTemplates[i]
	}
	if mode == tally.GameModeTemplate {
		i := p.Int("template")
		template = &tally.ChallengeGames[i]
	}
	model := NewGameModel(mode, template)

	sess := getSesssionId(s)
	cache.SetGame(sess, model)
	return model, nil
}
func getHint(ctx context.Context, s live.Socket, p live.Params) (interface{}, error) {
	model := NewThermoModel(s)
	model.Hints = model.Game.GetHint()
	model.HintButtonCounter++
	return model, nil
}

var generatedTemplates []tally.GameTemplate

func ReadGeneratedBoardsFromDisk() error {
	// generatorDir := path.Join("./generated")
	// generatorDir := generated.GenDir
	err := fs.WalkDir(generated.GenDir, ".", func(p string, info fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		b, err := fs.ReadFile(generated.GenDir, p)
		if err != nil {
			return err
		}
		var gen tally.GeneratedGame
		err = toml.Unmarshal(b, &gen)
		if err != nil {
			return err
		}
		template := tally.NewGameTemplate(gen.Hash, path.Base(p), "", 4, 4).
			SetGoalCheckerLargestValue(int64(gen.GeneratorOptions.TargetCellValue)).
			SetMaxMoves(gen.GeneratorOptions.MaxMoves).
			SetStartingLayout(gen.Cells...)

		generatedTemplates = append(generatedTemplates, *template)
		return nil
	})
	if err != nil {
		return err
	}
	return nil

}

// Example shows a simple temperature control using the
// "live-click" event.
func Example() {
	err := ReadGeneratedBoardsFromDisk()
	if err != nil {
		fmt.Println(fmt.Errorf("failed to read generated files: %w", err))
	}

	// Setup the handler.
	h := live.NewHandler()

	// Mount function is called on initial HTTP load and then initial web
	// socket connection. This should be used to create the initial state,
	// the socket Connected func will be true if the mount call is on a web
	// socket connection.
	h.HandleMount(thermoMount)
	tmpl := template.New("index")
	tmpl.Parse(tmpltIndexHtml)

	h.HandleRender(func(ctx context.Context, data *live.RenderContext) (io.Reader, error) {
		var buf bytes.Buffer
		d := map[string]interface{}{
			"data":          data,
			"templateGames": &tally.ChallengeGames,
		}
		d["data"] = data

		if err := tmpl.Execute(&buf, d); err != nil {
			return nil, err
		}
		return &buf, nil
	})

	// This handles the `live-click="temp-up"` button. First we load the model from
	// the socket, increment the temperature, and then return the new state of the
	// model. Live will now calculate the diff between the last time it rendered and now,
	// produce a set of diffs and push them to the browser to update.
	h.HandleEvent("swipe", swipe)
	h.HandleEvent("new-game", newGame)
	h.HandleEvent("select-cell", selectCell)
	h.HandleEvent("get-hint", getHint)

	http.Handle("/", live.NewHttpHandler(cookieStore, h))

	// This serves the JS needed to make live work.
	http.Handle("/live.js", live.Javascript{})
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Printf("starting... on port %s\n", port)
	http.ListenAndServe(":"+port, nil)
}
