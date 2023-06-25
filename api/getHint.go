package api

import (
	"context"
	"fmt"
	"sort"
	"time"

	"github.com/bufbuild/connect-go"
	model "github.com/runar-rkmedia/gotally/gen/proto/tally/v1"
	"github.com/runar-rkmedia/gotally/tallylogic"
)

func intsTouInt32s(ints []int) []uint32 {
	out := make([]uint32, len(ints))
	for i := 0; i < len(ints); i++ {
		out[i] = uint32(ints[i])
	}
	return out
}

const MaxUint = ^uint(0)
const MaxInt = int(MaxUint >> 1)

func (s *TallyServer) GetHint(
	ctx context.Context,
	req *connect.Request[model.GetHintRequest],
) (*connect.Response[model.GetHintResponse], error) {
	// There are usually a lot of hints available,
	// but some hints leads to more fun games than others
	// In general, multiplication is more fun than addition
	// and swiping is a bit dull.
	// Long hints seem like magic, so we prefer shorter hints.
	// However, too short hints are also boring
	// TODO: introduce a weighted hint and solution-sorter
	session := ContextGetUserState(ctx)

	response := &model.GetHintResponse{
		// Instruction: []*model.Instruction{},
	}
	// Get a single hint. Does not look ahead to do swipes etc.
	if session.Game.Rules.GameMode == tallylogic.GameModeRandom {
		session.Game.GetCombineHints(func(path []int, method tallylogic.EvalMethod) bool {

			indexes := make([]uint32, len(path))
			for i := 0; i < len(path); i++ {
				indexes[i] = uint32(path[i])
			}
			response.Instructions = []*model.Instruction{{
				InstructionOneof: &model.Instruction_Combine{
					Combine: &model.Indexes{
						Index: indexes,
					},
				},
			}}

			return true
		})
		if len(response.Instructions) > 0 {
			response.Method = model.HintMethod_HINT_METHOD_FALLBACK
			res := connect.NewResponse(response)
			return res, nil
		}
		hints := session.GetHint()
		if len(hints) > 0 {

			s.l.Debug().
				Bool("deep", false).
				Int("hintsFound", len(hints)).
				Msg("Returning hints")
			// TODO: sort these better.
			best := map[string]tallylogic.Hint{}
			for k, h := range hints {
				if h.Method != tallylogic.EvalMethodProduct {
					continue
				}
				best[k] = h
				response.Instructions = toModelHint(best)
				response.Method = *model.HintMethod_HINT_METHOD_DEPTH_FIRST.Enum()
				return connect.NewResponse(response), nil
			}
			for k, h := range hints {
				best[k] = h
				response.Instructions = toModelHint(best)
				return connect.NewResponse(response), nil
			}
		}
	}
	// Deeper hint, looking ahead to find better hints, attempting to solve the game if possible.
	// h := tallylogic.NewHintCalculator(session.Game, session.Game, session.Game)
	hintMethod, games, err := tallylogic.SolveGame(tallylogic.SolveOptions{
		MaxDepth:     10,
		MaxVisits:    6000,
		MinMoves:     0,
		MaxMoves:     10,
		MaxSolutions: 1,
		MaxTime:      time.Second * 10,
	}, session.Game, nil)
	switch hintMethod {
	case tallylogic.HintMethodBreadthFirst:
		response.Method = model.HintMethod_HINT_METHOD_BREADTH_FIRST
	case tallylogic.HintMethodDepthFirst:
		response.Method = model.HintMethod_HINT_METHOD_DEPTH_FIRST
	}
	if err != nil {

		session.Game.GetCombineHints(func(path []int, method tallylogic.EvalMethod) bool {

			indexes := make([]uint32, len(path))
			for i := 0; i < len(path); i++ {
				indexes[i] = uint32(path[i])
			}
			response.Instructions = []*model.Instruction{{
				InstructionOneof: &model.Instruction_Combine{
					Combine: &model.Indexes{
						Index: indexes,
					},
				},
			}}

			return true
		})
		s.l.Error().
			Err(err).
			Int("instruction-count", len(response.Instructions)).
			Msg("Failed to generate hint, returning fallback hint (if any)")
		if len(response.Instructions) > 0 {
			response.Method = model.HintMethod_HINT_METHOD_FALLBACK
			res := connect.NewResponse(response)
			return res, nil
		}
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("Failed to generate hint"))
	}
	response.Instructions = make([]*model.Instruction, 1)
	s.l.Debug().
		Int("solutions", len(games)).
		Msg("Solver returned solutions")
	if len(games) == 0 {
		return connect.NewResponse(response), nil
	}
	if req.Msg.MaxLength == 0 {
		req.Msg.MaxLength = 1
	}
	if req.Msg.HintPreference == model.HintPreference_HINT_PREFERENCE_UNSPECIFIED {
		if session.Game.Rules.GameMode == tallylogic.GameModeRandom {
			req.Msg.HintPreference = model.HintPreference_HINT_PREFERENCE_FIRST_COMBINE
		} else {
			req.Msg.HintPreference = model.HintPreference_HINT_PREFERENCE_SHORT
		}
	}
	historyOffset := session.Game.History.Length()
	switch req.Msg.HintPreference {
	case model.HintPreference_HINT_PREFERENCE_HIGHEST_SCORE:
		sort.Slice(games, func(i, j int) bool {
			return games[i].Score() < games[j].Score()
		})
	case model.HintPreference_HINT_PREFERENCE_SHORT:
		sort.Slice(games, func(i, j int) bool {
			return games[i].History.Length() < games[j].History.Length()
		})
	case model.HintPreference_HINT_PREFERENCE_MINIMUM_SWIPES:
		sort.Slice(games, func(i, j int) bool {
			var swipeI int
			var swipeJ int
			games[i].History.Iterate(
				func(dir tallylogic.SwipeDirection, i int) error {
					if i < historyOffset {
						return nil
					}
					swipeI++
					return nil
				},
				func(path []int, i int) error { return nil },
				func(helper tallylogic.Helper, i int) error { return nil },
			)
			games[j].History.Iterate(
				func(dir tallylogic.SwipeDirection, i int) error {
					if i < historyOffset {
						return nil
					}
					swipeJ++
					return nil
				},
				func(path []int, i int) error { return nil },
				func(helper tallylogic.Helper, i int) error { return nil },
			)
			return swipeI < swipeJ
		})
	case model.HintPreference_HINT_PREFERENCE_FIRST_COMBINE:
		sort.Slice(games, func(i, j int) bool {
			var combineIndexI int
			var combineIndexJ int
			// TODO
			games[i].History.Iterate(
				func(dir tallylogic.SwipeDirection, i int) error { return nil },
				func(path []int, i int) error {
					if i < historyOffset {
						return nil
					}
					combineIndexI = i
					return fmt.Errorf("stop")
				},
				func(helper tallylogic.Helper, i int) error { return nil },
			)
			games[j].History.Iterate(
				func(dir tallylogic.SwipeDirection, i int) error { return nil },
				func(path []int, i int) error {
					if i < historyOffset {
						return nil
					}
					combineIndexJ = i
					return fmt.Errorf("stop")
				},
				func(helper tallylogic.Helper, i int) error { return nil },
			)
			if combineIndexI == combineIndexJ {
				return games[i].History.Length() < games[j].History.Length()
			}
			return combineIndexI < combineIndexJ
		})
	case model.HintPreference_HINT_PREFERENCE_MINIMUM_SWIPES_TO_COMBINE_RATIO:
		sort.Slice(games, func(i, j int) bool {
			var swipeI float32
			var swipeJ float32
			var combineI float32
			var combineJ float32
			games[i].History.Iterate(
				func(dir tallylogic.SwipeDirection, i int) error {
					if i < historyOffset {
						return nil
					}
					swipeI++
					return nil
				},
				func(path []int, i int) error {
					if i < historyOffset {
						return nil
					}
					combineI++
					return nil
				},
				func(helper tallylogic.Helper, i int) error { return nil },
			)
			games[j].History.Iterate(
				func(dir tallylogic.SwipeDirection, i int) error {
					if i < historyOffset {
						return nil
					}
					swipeJ++
					return nil
				},
				func(path []int, i int) error {
					if i < historyOffset {
						return nil
					}
					combineJ++
					return nil
				},
				func(helper tallylogic.Helper, i int) error { return nil },
			)
			ratioI := swipeI / combineI
			ratioJ := swipeJ / combineI
			return ratioI < ratioJ
		})
	}
	bestInstructions := games[0].History
	var length int = bestInstructions.Length() - historyOffset
	if req.Msg.MaxLength > 0 && req.Msg.MaxLength < uint32(length) {
		length = int(req.Msg.MaxLength)
	}
	fmt.Println("length", length, historyOffset, bestInstructions.Describe())
	ins, err := toModelInstruction(bestInstructions, historyOffset)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to map instructions: %#v", err))
	}
	response.Instructions = ins
	res := connect.NewResponse(response)
	return res, nil
}
