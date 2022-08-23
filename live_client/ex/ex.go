package ex

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"io"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/jfyne/live"
	"github.com/pelletier/go-toml/v2"
	"github.com/runar-rkmedia/gotally/database"
	"github.com/runar-rkmedia/gotally/generated"
	tally "github.com/runar-rkmedia/gotally/tallylogic"
)

// Model of our thermostat.
type GameModel struct {
	Template          *tally.GameTemplate
	Hints             map[string]tally.Hint
	Error             string
	HintButtonCounter int
	UserName          string
	SelfVotes         map[string]int
	tally.Game
}

var (
	startedAt time.Time
)

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
	m := GameModel{SelfVotes: map[string]int{}}
	fmt.Println("created new model")
	m.Template = template
	game, err := tally.NewGame(mode, template)
	if err != nil {
		log.Printf("creating new game failed: %v", err)
		m.Error = err.Error()
	} else {
		m.Error = ""
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
		log.Println("no session-id, new game")
		mode := tally.GameModeTemplate
		if len(generatedTemplates) > 0 {
			mode := tally.GameModeRandomChallenge
			i := tally.NewRandomizer().Intn(len(generatedTemplates))
			m = NewGameModel(mode, &generatedTemplates[i])
		} else {
			m = NewGameModel(mode, &tally.ChallengeGames[0])
		}
		if m.SelfVotes == nil {
			m.SelfVotes = map[string]int{}
		}
		if m.UserName != "" {
			if votes, err := db.GetVotesForBoardByUserName(m.UserName); err == nil {
				for k, v := range votes {
					m.SelfVotes[k] = v.FunVote
					fmt.Println("user vote", k, v.FunVote)
				}
			} else {
				log.Println("faield to retrieve votes", err)
			}
		} else {
			fmt.Println("no username")
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
func restart(ctx context.Context, s live.Socket, p live.Params) (interface{}, error) {
	model := NewThermoModel(s)
	if model.Template == nil {
		return model, fmt.Errorf("currently, only tempated games can be restarted")
	}

	fmt.Println("ex", model.Rules.GameMode, model.Template)
	v := NewGameModel(tally.GameModeRandomChallenge, model.Template)
	model.Game = v.Game
	model.Error = v.Error
	model.Hints = v.Hints
	return model, nil
}

var db database.DB

func init() {
	d, err := database.NewDatabase("")
	if err != nil {
		panic(err)
	}
	db = d

}

func setUserName(ctx context.Context, s live.Socket, p live.Params) (interface{}, error) {

	userName := strings.TrimSpace(p.String("username"))
	fmt.Println("estuser", userName)
	ex := NewThermoModel(s)
	if userName == "" {
		return ex, fmt.Errorf("no username")
	}
	ex.UserName = userName
	if votes, err := db.GetVotesForBoardByUserName(ex.UserName); err == nil {
		log.Printf("\nFound %d votes user %s", len(votes), ex.UserName)
		for k, v := range votes {
			ex.SelfVotes[k] = v.FunVote
			fmt.Println("user vote", k, v.FunVote)
		}
	} else {
		log.Println("faield to retrieve votes", err)
	}

	return ex, nil
}
func vote(ctx context.Context, s live.Socket, p live.Params) (interface{}, error) {
	voteValue := p.Int("vote")
	userName := p.String("username")
	fmt.Println("username???", userName, voteValue, p)
	ex := NewThermoModel(s)
	if userName == "" {
		userName = ex.UserName
	} else {
		ex.UserName = userName
	}
	if userName == "" {
		return ex, fmt.Errorf("username is required")
	}
	if voteValue == 0 {
		return ex, fmt.Errorf("no value")
	}
	if ex.Template == nil {
		return ex, fmt.Errorf("only templated games can be voted on")
	}
	sessID := getSesssionId(s)
	vote, err := db.VoteForBoard(ex.Template.ID, sessID, userName, int(voteValue))
	if err != nil {
		return ex, err
	}
	ex.SelfVotes[vote.ID] = voteValue
	for k, v := range ex.SelfVotes {
		fmt.Println("User has votes", k, v)

	}

	return ex, nil
}
func newGame(ctx context.Context, s live.Socket, p live.Params) (interface{}, error) {
	model := NewThermoModel(s)
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
	v := NewGameModel(mode, template)
	model.Game = v.Game
	model.Hints = v.Hints
	model.Error = v.Error
	model.Template = v.Template

	sess := getSesssionId(s)
	cache.SetGame(sess, model)
	return model, nil
}
func getHint(ctx context.Context, s live.Socket, p live.Params) (interface{}, error) {
	model := NewThermoModel(s)
	model.Hints = model.Game.GetHint()
	if model.Template != nil && false {
		solver := tally.NewBruteSolver(tally.SolveOptions{
			MaxDepth:     0,
			MaxVisits:    0,
			MinMoves:     0,
			MaxMoves:     20,
			MaxSolutions: 1,
		})
		solutions, err := solver.SolveGame(model.Game)
		if err != nil {
			log.Println("failed to solve game", err)
		} else if len(solutions) > 0 {
			// model.Hints = map[string]tally.Hint{}
			// for _, s := range solutions[0].History {
			// 	model.Hints["s"] = tally.Hint{
			// 	}
			// }
		}
	}
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
		// if gen.GeneratorOptions.Rows != 3 {
		// 	return nil
		// }
		var description string
		if len(gen.Solutions) > 0 {
			description = fmt.Sprintf("Get at least one cell to a value of %d. This game can be solved in %d moves, with the highest cell at %d", gen.GeneratorOptions.TargetCellValue, gen.Solutions[0].Moves, gen.Solutions[0].HighestCellValue)

		}
		template := tally.NewGameTemplate(gen.Hash, gen.Name, description, gen.GeneratorOptions.Columns, gen.GeneratorOptions.Rows).
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
	startedAt = time.Now()
	err := ReadGeneratedBoardsFromDisk()
	if err != nil {
		log.Printf("failed to read generated files: %s", err.Error())
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
			"startedAt":     startedAt,
			"data":          data,
			"templateGames": &tally.ChallengeGames,
		}

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
	h.HandleEvent("restart", restart)
	h.HandleEvent("vote", vote)
	h.HandleEvent("set-username", setUserName)

	http.Handle("/", live.NewHttpHandler(cookieStore, h))

	// This serves the JS needed to make live work.
	http.Handle("/live.js", live.Javascript{})
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("starting... on port %s\n", port)
	http.ListenAndServe(":"+port, nil)
}
