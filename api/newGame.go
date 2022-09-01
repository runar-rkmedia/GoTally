package api

import (
	"context"
	"errors"
	"fmt"
	"math/rand"

	"github.com/bufbuild/connect-go"
	model "github.com/runar-rkmedia/gotally/gen/proto/tally/v1"
	"github.com/runar-rkmedia/gotally/live_client/ex"
	logic "github.com/runar-rkmedia/gotally/tallylogic"
)

var (
	ErrSliceUnderflow = errors.New("Slice index underflow")
	ErrSliceOverFlow  = errors.New("Slice index overflow")
)

func withinSlice(length, index int) (int, error) {
	if index < 0 {
		return -(index % length), ErrSliceUnderflow
	}
	if index >= length {
		return (index % length), ErrSliceOverFlow
	}
	return index, nil
}

func (s *TallyServer) NewGame(
	ctx context.Context,
	req *connect.Request[model.NewGameRequest],
) (*connect.Response[model.NewGameResponse], error) {
	session := UserStateFromContext(ctx)
	var mode logic.GameMode
	var template *logic.GameTemplate

	switch req.Msg.Mode {
	case model.GameMode_GAME_MODE_RANDOM_CHALLENGE:
		mode = logic.GameModeRandomChallenge
		fmt.Println(len(ex.GeneratedTemplates))
		index := rand.Intn(len(ex.GeneratedTemplates))
		template = &ex.GeneratedTemplates[index]
	case model.GameMode_GAME_MODE_TUTORIAL:
		mode = logic.GameModeTemplate
		if _i, ok := req.Msg.Variant.(*model.NewGameRequest_LevelIndex); ok {
			i := int(_i.LevelIndex)
			if i < 0 {
				return nil, fmt.Errorf("invalid levelindex, must be positive got %d", i)
			}
			if i >= len(logic.ChallengeGames) {
				return nil, fmt.Errorf("invalid levelindex, got %d, must be lower than %d", i, len(logic.ChallengeGames))
			}
			template = &logic.ChallengeGames[i]
		} else {
			template = &logic.ChallengeGames[0]
		}

	}

	game, err := logic.NewGame(mode, template)
	if err != nil {
		return nil, fmt.Errorf("failed to created game: %w", err)
	}
	session.Game = game
	session.GamesStarted++
	session.GameSnapshotAtStart = game.Copy()
	Store.SetUserState(session)
	response := &model.NewGameResponse{
		Board: toModalBoard(&session.Game),
		Score: session.Game.Score(),
		Moves: int64(session.Game.Moves()),
	}
	res := connect.NewResponse(response)
	return res, nil
}
