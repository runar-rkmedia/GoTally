// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: proto/tally/v1/board.proto

package tallyv1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/runar-rkmedia/gotally/gen/proto/tally/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// BoardServiceName is the fully-qualified name of the BoardService service.
	BoardServiceName = "tally.v1.BoardService"
)

// BoardServiceClient is a client for the tally.v1.BoardService service.
type BoardServiceClient interface {
	NewGame(context.Context, *connect_go.Request[v1.NewGameRequest]) (*connect_go.Response[v1.NewGameResponse], error)
	NewGameFromTemplate(context.Context, *connect_go.Request[v1.NewGameFromTemplateRequest]) (*connect_go.Response[v1.NewGameFromTemplateResponse], error)
	GetHint(context.Context, *connect_go.Request[v1.GetHintRequest]) (*connect_go.Response[v1.GetHintResponse], error)
	RestartGame(context.Context, *connect_go.Request[v1.RestartGameRequest]) (*connect_go.Response[v1.RestartGameResponse], error)
	GetSession(context.Context, *connect_go.Request[v1.GetSessionRequest]) (*connect_go.Response[v1.GetSessionResponse], error)
	SwipeBoard(context.Context, *connect_go.Request[v1.SwipeBoardRequest]) (*connect_go.Response[v1.SwipeBoardResponse], error)
	CombineCells(context.Context, *connect_go.Request[v1.CombineCellsRequest]) (*connect_go.Response[v1.CombineCellsResponse], error)
	GenerateGame(context.Context, *connect_go.Request[v1.GenerateGameRequest]) (*connect_go.Response[v1.GenerateGameResponse], error)
	VoteBoard(context.Context, *connect_go.Request[v1.VoteBoardRequest]) (*connect_go.Response[v1.VoteBoardResponse], error)
	GetGameChallenges(context.Context, *connect_go.Request[v1.GetGameChallengesRequest]) (*connect_go.Response[v1.GetGameChallengesResponse], error)
	CreateGameChallenge(context.Context, *connect_go.Request[v1.CreateGameChallengeRequest]) (*connect_go.Response[v1.CreateGameChallengeResponse], error)
}

// NewBoardServiceClient constructs a client for the tally.v1.BoardService service. By default, it
// uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewBoardServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) BoardServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &boardServiceClient{
		newGame: connect_go.NewClient[v1.NewGameRequest, v1.NewGameResponse](
			httpClient,
			baseURL+"/tally.v1.BoardService/NewGame",
			opts...,
		),
		newGameFromTemplate: connect_go.NewClient[v1.NewGameFromTemplateRequest, v1.NewGameFromTemplateResponse](
			httpClient,
			baseURL+"/tally.v1.BoardService/NewGameFromTemplate",
			opts...,
		),
		getHint: connect_go.NewClient[v1.GetHintRequest, v1.GetHintResponse](
			httpClient,
			baseURL+"/tally.v1.BoardService/GetHint",
			opts...,
		),
		restartGame: connect_go.NewClient[v1.RestartGameRequest, v1.RestartGameResponse](
			httpClient,
			baseURL+"/tally.v1.BoardService/RestartGame",
			opts...,
		),
		getSession: connect_go.NewClient[v1.GetSessionRequest, v1.GetSessionResponse](
			httpClient,
			baseURL+"/tally.v1.BoardService/GetSession",
			opts...,
		),
		swipeBoard: connect_go.NewClient[v1.SwipeBoardRequest, v1.SwipeBoardResponse](
			httpClient,
			baseURL+"/tally.v1.BoardService/SwipeBoard",
			opts...,
		),
		combineCells: connect_go.NewClient[v1.CombineCellsRequest, v1.CombineCellsResponse](
			httpClient,
			baseURL+"/tally.v1.BoardService/CombineCells",
			opts...,
		),
		generateGame: connect_go.NewClient[v1.GenerateGameRequest, v1.GenerateGameResponse](
			httpClient,
			baseURL+"/tally.v1.BoardService/GenerateGame",
			opts...,
		),
		voteBoard: connect_go.NewClient[v1.VoteBoardRequest, v1.VoteBoardResponse](
			httpClient,
			baseURL+"/tally.v1.BoardService/VoteBoard",
			opts...,
		),
		getGameChallenges: connect_go.NewClient[v1.GetGameChallengesRequest, v1.GetGameChallengesResponse](
			httpClient,
			baseURL+"/tally.v1.BoardService/GetGameChallenges",
			opts...,
		),
		createGameChallenge: connect_go.NewClient[v1.CreateGameChallengeRequest, v1.CreateGameChallengeResponse](
			httpClient,
			baseURL+"/tally.v1.BoardService/CreateGameChallenge",
			opts...,
		),
	}
}

// boardServiceClient implements BoardServiceClient.
type boardServiceClient struct {
	newGame             *connect_go.Client[v1.NewGameRequest, v1.NewGameResponse]
	newGameFromTemplate *connect_go.Client[v1.NewGameFromTemplateRequest, v1.NewGameFromTemplateResponse]
	getHint             *connect_go.Client[v1.GetHintRequest, v1.GetHintResponse]
	restartGame         *connect_go.Client[v1.RestartGameRequest, v1.RestartGameResponse]
	getSession          *connect_go.Client[v1.GetSessionRequest, v1.GetSessionResponse]
	swipeBoard          *connect_go.Client[v1.SwipeBoardRequest, v1.SwipeBoardResponse]
	combineCells        *connect_go.Client[v1.CombineCellsRequest, v1.CombineCellsResponse]
	generateGame        *connect_go.Client[v1.GenerateGameRequest, v1.GenerateGameResponse]
	voteBoard           *connect_go.Client[v1.VoteBoardRequest, v1.VoteBoardResponse]
	getGameChallenges   *connect_go.Client[v1.GetGameChallengesRequest, v1.GetGameChallengesResponse]
	createGameChallenge *connect_go.Client[v1.CreateGameChallengeRequest, v1.CreateGameChallengeResponse]
}

// NewGame calls tally.v1.BoardService.NewGame.
func (c *boardServiceClient) NewGame(ctx context.Context, req *connect_go.Request[v1.NewGameRequest]) (*connect_go.Response[v1.NewGameResponse], error) {
	return c.newGame.CallUnary(ctx, req)
}

// NewGameFromTemplate calls tally.v1.BoardService.NewGameFromTemplate.
func (c *boardServiceClient) NewGameFromTemplate(ctx context.Context, req *connect_go.Request[v1.NewGameFromTemplateRequest]) (*connect_go.Response[v1.NewGameFromTemplateResponse], error) {
	return c.newGameFromTemplate.CallUnary(ctx, req)
}

// GetHint calls tally.v1.BoardService.GetHint.
func (c *boardServiceClient) GetHint(ctx context.Context, req *connect_go.Request[v1.GetHintRequest]) (*connect_go.Response[v1.GetHintResponse], error) {
	return c.getHint.CallUnary(ctx, req)
}

// RestartGame calls tally.v1.BoardService.RestartGame.
func (c *boardServiceClient) RestartGame(ctx context.Context, req *connect_go.Request[v1.RestartGameRequest]) (*connect_go.Response[v1.RestartGameResponse], error) {
	return c.restartGame.CallUnary(ctx, req)
}

// GetSession calls tally.v1.BoardService.GetSession.
func (c *boardServiceClient) GetSession(ctx context.Context, req *connect_go.Request[v1.GetSessionRequest]) (*connect_go.Response[v1.GetSessionResponse], error) {
	return c.getSession.CallUnary(ctx, req)
}

// SwipeBoard calls tally.v1.BoardService.SwipeBoard.
func (c *boardServiceClient) SwipeBoard(ctx context.Context, req *connect_go.Request[v1.SwipeBoardRequest]) (*connect_go.Response[v1.SwipeBoardResponse], error) {
	return c.swipeBoard.CallUnary(ctx, req)
}

// CombineCells calls tally.v1.BoardService.CombineCells.
func (c *boardServiceClient) CombineCells(ctx context.Context, req *connect_go.Request[v1.CombineCellsRequest]) (*connect_go.Response[v1.CombineCellsResponse], error) {
	return c.combineCells.CallUnary(ctx, req)
}

// GenerateGame calls tally.v1.BoardService.GenerateGame.
func (c *boardServiceClient) GenerateGame(ctx context.Context, req *connect_go.Request[v1.GenerateGameRequest]) (*connect_go.Response[v1.GenerateGameResponse], error) {
	return c.generateGame.CallUnary(ctx, req)
}

// VoteBoard calls tally.v1.BoardService.VoteBoard.
func (c *boardServiceClient) VoteBoard(ctx context.Context, req *connect_go.Request[v1.VoteBoardRequest]) (*connect_go.Response[v1.VoteBoardResponse], error) {
	return c.voteBoard.CallUnary(ctx, req)
}

// GetGameChallenges calls tally.v1.BoardService.GetGameChallenges.
func (c *boardServiceClient) GetGameChallenges(ctx context.Context, req *connect_go.Request[v1.GetGameChallengesRequest]) (*connect_go.Response[v1.GetGameChallengesResponse], error) {
	return c.getGameChallenges.CallUnary(ctx, req)
}

// CreateGameChallenge calls tally.v1.BoardService.CreateGameChallenge.
func (c *boardServiceClient) CreateGameChallenge(ctx context.Context, req *connect_go.Request[v1.CreateGameChallengeRequest]) (*connect_go.Response[v1.CreateGameChallengeResponse], error) {
	return c.createGameChallenge.CallUnary(ctx, req)
}

// BoardServiceHandler is an implementation of the tally.v1.BoardService service.
type BoardServiceHandler interface {
	NewGame(context.Context, *connect_go.Request[v1.NewGameRequest]) (*connect_go.Response[v1.NewGameResponse], error)
	NewGameFromTemplate(context.Context, *connect_go.Request[v1.NewGameFromTemplateRequest]) (*connect_go.Response[v1.NewGameFromTemplateResponse], error)
	GetHint(context.Context, *connect_go.Request[v1.GetHintRequest]) (*connect_go.Response[v1.GetHintResponse], error)
	RestartGame(context.Context, *connect_go.Request[v1.RestartGameRequest]) (*connect_go.Response[v1.RestartGameResponse], error)
	GetSession(context.Context, *connect_go.Request[v1.GetSessionRequest]) (*connect_go.Response[v1.GetSessionResponse], error)
	SwipeBoard(context.Context, *connect_go.Request[v1.SwipeBoardRequest]) (*connect_go.Response[v1.SwipeBoardResponse], error)
	CombineCells(context.Context, *connect_go.Request[v1.CombineCellsRequest]) (*connect_go.Response[v1.CombineCellsResponse], error)
	GenerateGame(context.Context, *connect_go.Request[v1.GenerateGameRequest]) (*connect_go.Response[v1.GenerateGameResponse], error)
	VoteBoard(context.Context, *connect_go.Request[v1.VoteBoardRequest]) (*connect_go.Response[v1.VoteBoardResponse], error)
	GetGameChallenges(context.Context, *connect_go.Request[v1.GetGameChallengesRequest]) (*connect_go.Response[v1.GetGameChallengesResponse], error)
	CreateGameChallenge(context.Context, *connect_go.Request[v1.CreateGameChallengeRequest]) (*connect_go.Response[v1.CreateGameChallengeResponse], error)
}

// NewBoardServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewBoardServiceHandler(svc BoardServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/tally.v1.BoardService/NewGame", connect_go.NewUnaryHandler(
		"/tally.v1.BoardService/NewGame",
		svc.NewGame,
		opts...,
	))
	mux.Handle("/tally.v1.BoardService/NewGameFromTemplate", connect_go.NewUnaryHandler(
		"/tally.v1.BoardService/NewGameFromTemplate",
		svc.NewGameFromTemplate,
		opts...,
	))
	mux.Handle("/tally.v1.BoardService/GetHint", connect_go.NewUnaryHandler(
		"/tally.v1.BoardService/GetHint",
		svc.GetHint,
		opts...,
	))
	mux.Handle("/tally.v1.BoardService/RestartGame", connect_go.NewUnaryHandler(
		"/tally.v1.BoardService/RestartGame",
		svc.RestartGame,
		opts...,
	))
	mux.Handle("/tally.v1.BoardService/GetSession", connect_go.NewUnaryHandler(
		"/tally.v1.BoardService/GetSession",
		svc.GetSession,
		opts...,
	))
	mux.Handle("/tally.v1.BoardService/SwipeBoard", connect_go.NewUnaryHandler(
		"/tally.v1.BoardService/SwipeBoard",
		svc.SwipeBoard,
		opts...,
	))
	mux.Handle("/tally.v1.BoardService/CombineCells", connect_go.NewUnaryHandler(
		"/tally.v1.BoardService/CombineCells",
		svc.CombineCells,
		opts...,
	))
	mux.Handle("/tally.v1.BoardService/GenerateGame", connect_go.NewUnaryHandler(
		"/tally.v1.BoardService/GenerateGame",
		svc.GenerateGame,
		opts...,
	))
	mux.Handle("/tally.v1.BoardService/VoteBoard", connect_go.NewUnaryHandler(
		"/tally.v1.BoardService/VoteBoard",
		svc.VoteBoard,
		opts...,
	))
	mux.Handle("/tally.v1.BoardService/GetGameChallenges", connect_go.NewUnaryHandler(
		"/tally.v1.BoardService/GetGameChallenges",
		svc.GetGameChallenges,
		opts...,
	))
	mux.Handle("/tally.v1.BoardService/CreateGameChallenge", connect_go.NewUnaryHandler(
		"/tally.v1.BoardService/CreateGameChallenge",
		svc.CreateGameChallenge,
		opts...,
	))
	return "/tally.v1.BoardService/", mux
}

// UnimplementedBoardServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedBoardServiceHandler struct{}

func (UnimplementedBoardServiceHandler) NewGame(context.Context, *connect_go.Request[v1.NewGameRequest]) (*connect_go.Response[v1.NewGameResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("tally.v1.BoardService.NewGame is not implemented"))
}

func (UnimplementedBoardServiceHandler) NewGameFromTemplate(context.Context, *connect_go.Request[v1.NewGameFromTemplateRequest]) (*connect_go.Response[v1.NewGameFromTemplateResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("tally.v1.BoardService.NewGameFromTemplate is not implemented"))
}

func (UnimplementedBoardServiceHandler) GetHint(context.Context, *connect_go.Request[v1.GetHintRequest]) (*connect_go.Response[v1.GetHintResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("tally.v1.BoardService.GetHint is not implemented"))
}

func (UnimplementedBoardServiceHandler) RestartGame(context.Context, *connect_go.Request[v1.RestartGameRequest]) (*connect_go.Response[v1.RestartGameResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("tally.v1.BoardService.RestartGame is not implemented"))
}

func (UnimplementedBoardServiceHandler) GetSession(context.Context, *connect_go.Request[v1.GetSessionRequest]) (*connect_go.Response[v1.GetSessionResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("tally.v1.BoardService.GetSession is not implemented"))
}

func (UnimplementedBoardServiceHandler) SwipeBoard(context.Context, *connect_go.Request[v1.SwipeBoardRequest]) (*connect_go.Response[v1.SwipeBoardResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("tally.v1.BoardService.SwipeBoard is not implemented"))
}

func (UnimplementedBoardServiceHandler) CombineCells(context.Context, *connect_go.Request[v1.CombineCellsRequest]) (*connect_go.Response[v1.CombineCellsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("tally.v1.BoardService.CombineCells is not implemented"))
}

func (UnimplementedBoardServiceHandler) GenerateGame(context.Context, *connect_go.Request[v1.GenerateGameRequest]) (*connect_go.Response[v1.GenerateGameResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("tally.v1.BoardService.GenerateGame is not implemented"))
}

func (UnimplementedBoardServiceHandler) VoteBoard(context.Context, *connect_go.Request[v1.VoteBoardRequest]) (*connect_go.Response[v1.VoteBoardResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("tally.v1.BoardService.VoteBoard is not implemented"))
}

func (UnimplementedBoardServiceHandler) GetGameChallenges(context.Context, *connect_go.Request[v1.GetGameChallengesRequest]) (*connect_go.Response[v1.GetGameChallengesResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("tally.v1.BoardService.GetGameChallenges is not implemented"))
}

func (UnimplementedBoardServiceHandler) CreateGameChallenge(context.Context, *connect_go.Request[v1.CreateGameChallengeRequest]) (*connect_go.Response[v1.CreateGameChallengeResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("tally.v1.BoardService.CreateGameChallenge is not implemented"))
}
