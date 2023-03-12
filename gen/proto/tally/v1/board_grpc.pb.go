// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: proto/tally/v1/board.proto

package tallyv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// BoardServiceClient is the client API for BoardService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BoardServiceClient interface {
	NewGame(ctx context.Context, in *NewGameRequest, opts ...grpc.CallOption) (*NewGameResponse, error)
	NewGameFromTemplate(ctx context.Context, in *NewGameFromTemplateRequest, opts ...grpc.CallOption) (*NewGameResponse, error)
	GetHint(ctx context.Context, in *GetHintRequest, opts ...grpc.CallOption) (*GetHintResponse, error)
	RestartGame(ctx context.Context, in *RestartGameRequest, opts ...grpc.CallOption) (*RestartGameResponse, error)
	GetSession(ctx context.Context, in *GetSessionRequest, opts ...grpc.CallOption) (*GetSessionResponse, error)
	SwipeBoard(ctx context.Context, in *SwipeBoardRequest, opts ...grpc.CallOption) (*SwipeBoardResponse, error)
	CombineCells(ctx context.Context, in *CombineCellsRequest, opts ...grpc.CallOption) (*CombineCellsResponse, error)
	GenerateGame(ctx context.Context, in *GenerateGameRequest, opts ...grpc.CallOption) (*GenerateGameResponse, error)
	VoteBoard(ctx context.Context, in *VoteBoardRequest, opts ...grpc.CallOption) (*VoteBoardResponse, error)
	GetGameChallenges(ctx context.Context, in *GetGameChallengesRequest, opts ...grpc.CallOption) (*GetGameChallengesResponse, error)
	CreateGameChallenge(ctx context.Context, in *CreateGameChallengeRequest, opts ...grpc.CallOption) (*CreateGameChallengeResponse, error)
}

type boardServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBoardServiceClient(cc grpc.ClientConnInterface) BoardServiceClient {
	return &boardServiceClient{cc}
}

func (c *boardServiceClient) NewGame(ctx context.Context, in *NewGameRequest, opts ...grpc.CallOption) (*NewGameResponse, error) {
	out := new(NewGameResponse)
	err := c.cc.Invoke(ctx, "/tally.v1.BoardService/NewGame", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *boardServiceClient) NewGameFromTemplate(ctx context.Context, in *NewGameFromTemplateRequest, opts ...grpc.CallOption) (*NewGameResponse, error) {
	out := new(NewGameResponse)
	err := c.cc.Invoke(ctx, "/tally.v1.BoardService/NewGameFromTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *boardServiceClient) GetHint(ctx context.Context, in *GetHintRequest, opts ...grpc.CallOption) (*GetHintResponse, error) {
	out := new(GetHintResponse)
	err := c.cc.Invoke(ctx, "/tally.v1.BoardService/GetHint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *boardServiceClient) RestartGame(ctx context.Context, in *RestartGameRequest, opts ...grpc.CallOption) (*RestartGameResponse, error) {
	out := new(RestartGameResponse)
	err := c.cc.Invoke(ctx, "/tally.v1.BoardService/RestartGame", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *boardServiceClient) GetSession(ctx context.Context, in *GetSessionRequest, opts ...grpc.CallOption) (*GetSessionResponse, error) {
	out := new(GetSessionResponse)
	err := c.cc.Invoke(ctx, "/tally.v1.BoardService/GetSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *boardServiceClient) SwipeBoard(ctx context.Context, in *SwipeBoardRequest, opts ...grpc.CallOption) (*SwipeBoardResponse, error) {
	out := new(SwipeBoardResponse)
	err := c.cc.Invoke(ctx, "/tally.v1.BoardService/SwipeBoard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *boardServiceClient) CombineCells(ctx context.Context, in *CombineCellsRequest, opts ...grpc.CallOption) (*CombineCellsResponse, error) {
	out := new(CombineCellsResponse)
	err := c.cc.Invoke(ctx, "/tally.v1.BoardService/CombineCells", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *boardServiceClient) GenerateGame(ctx context.Context, in *GenerateGameRequest, opts ...grpc.CallOption) (*GenerateGameResponse, error) {
	out := new(GenerateGameResponse)
	err := c.cc.Invoke(ctx, "/tally.v1.BoardService/GenerateGame", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *boardServiceClient) VoteBoard(ctx context.Context, in *VoteBoardRequest, opts ...grpc.CallOption) (*VoteBoardResponse, error) {
	out := new(VoteBoardResponse)
	err := c.cc.Invoke(ctx, "/tally.v1.BoardService/VoteBoard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *boardServiceClient) GetGameChallenges(ctx context.Context, in *GetGameChallengesRequest, opts ...grpc.CallOption) (*GetGameChallengesResponse, error) {
	out := new(GetGameChallengesResponse)
	err := c.cc.Invoke(ctx, "/tally.v1.BoardService/GetGameChallenges", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *boardServiceClient) CreateGameChallenge(ctx context.Context, in *CreateGameChallengeRequest, opts ...grpc.CallOption) (*CreateGameChallengeResponse, error) {
	out := new(CreateGameChallengeResponse)
	err := c.cc.Invoke(ctx, "/tally.v1.BoardService/CreateGameChallenge", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BoardServiceServer is the server API for BoardService service.
// All implementations should embed UnimplementedBoardServiceServer
// for forward compatibility
type BoardServiceServer interface {
	NewGame(context.Context, *NewGameRequest) (*NewGameResponse, error)
	NewGameFromTemplate(context.Context, *NewGameFromTemplateRequest) (*NewGameResponse, error)
	GetHint(context.Context, *GetHintRequest) (*GetHintResponse, error)
	RestartGame(context.Context, *RestartGameRequest) (*RestartGameResponse, error)
	GetSession(context.Context, *GetSessionRequest) (*GetSessionResponse, error)
	SwipeBoard(context.Context, *SwipeBoardRequest) (*SwipeBoardResponse, error)
	CombineCells(context.Context, *CombineCellsRequest) (*CombineCellsResponse, error)
	GenerateGame(context.Context, *GenerateGameRequest) (*GenerateGameResponse, error)
	VoteBoard(context.Context, *VoteBoardRequest) (*VoteBoardResponse, error)
	GetGameChallenges(context.Context, *GetGameChallengesRequest) (*GetGameChallengesResponse, error)
	CreateGameChallenge(context.Context, *CreateGameChallengeRequest) (*CreateGameChallengeResponse, error)
}

// UnimplementedBoardServiceServer should be embedded to have forward compatible implementations.
type UnimplementedBoardServiceServer struct {
}

func (UnimplementedBoardServiceServer) NewGame(context.Context, *NewGameRequest) (*NewGameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewGame not implemented")
}
func (UnimplementedBoardServiceServer) NewGameFromTemplate(context.Context, *NewGameFromTemplateRequest) (*NewGameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewGameFromTemplate not implemented")
}
func (UnimplementedBoardServiceServer) GetHint(context.Context, *GetHintRequest) (*GetHintResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHint not implemented")
}
func (UnimplementedBoardServiceServer) RestartGame(context.Context, *RestartGameRequest) (*RestartGameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RestartGame not implemented")
}
func (UnimplementedBoardServiceServer) GetSession(context.Context, *GetSessionRequest) (*GetSessionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSession not implemented")
}
func (UnimplementedBoardServiceServer) SwipeBoard(context.Context, *SwipeBoardRequest) (*SwipeBoardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SwipeBoard not implemented")
}
func (UnimplementedBoardServiceServer) CombineCells(context.Context, *CombineCellsRequest) (*CombineCellsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CombineCells not implemented")
}
func (UnimplementedBoardServiceServer) GenerateGame(context.Context, *GenerateGameRequest) (*GenerateGameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateGame not implemented")
}
func (UnimplementedBoardServiceServer) VoteBoard(context.Context, *VoteBoardRequest) (*VoteBoardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VoteBoard not implemented")
}
func (UnimplementedBoardServiceServer) GetGameChallenges(context.Context, *GetGameChallengesRequest) (*GetGameChallengesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGameChallenges not implemented")
}
func (UnimplementedBoardServiceServer) CreateGameChallenge(context.Context, *CreateGameChallengeRequest) (*CreateGameChallengeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGameChallenge not implemented")
}

// UnsafeBoardServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BoardServiceServer will
// result in compilation errors.
type UnsafeBoardServiceServer interface {
	mustEmbedUnimplementedBoardServiceServer()
}

func RegisterBoardServiceServer(s grpc.ServiceRegistrar, srv BoardServiceServer) {
	s.RegisterService(&BoardService_ServiceDesc, srv)
}

func _BoardService_NewGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewGameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoardServiceServer).NewGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tally.v1.BoardService/NewGame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoardServiceServer).NewGame(ctx, req.(*NewGameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BoardService_NewGameFromTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewGameFromTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoardServiceServer).NewGameFromTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tally.v1.BoardService/NewGameFromTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoardServiceServer).NewGameFromTemplate(ctx, req.(*NewGameFromTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BoardService_GetHint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHintRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoardServiceServer).GetHint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tally.v1.BoardService/GetHint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoardServiceServer).GetHint(ctx, req.(*GetHintRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BoardService_RestartGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RestartGameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoardServiceServer).RestartGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tally.v1.BoardService/RestartGame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoardServiceServer).RestartGame(ctx, req.(*RestartGameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BoardService_GetSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoardServiceServer).GetSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tally.v1.BoardService/GetSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoardServiceServer).GetSession(ctx, req.(*GetSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BoardService_SwipeBoard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SwipeBoardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoardServiceServer).SwipeBoard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tally.v1.BoardService/SwipeBoard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoardServiceServer).SwipeBoard(ctx, req.(*SwipeBoardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BoardService_CombineCells_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CombineCellsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoardServiceServer).CombineCells(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tally.v1.BoardService/CombineCells",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoardServiceServer).CombineCells(ctx, req.(*CombineCellsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BoardService_GenerateGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateGameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoardServiceServer).GenerateGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tally.v1.BoardService/GenerateGame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoardServiceServer).GenerateGame(ctx, req.(*GenerateGameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BoardService_VoteBoard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VoteBoardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoardServiceServer).VoteBoard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tally.v1.BoardService/VoteBoard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoardServiceServer).VoteBoard(ctx, req.(*VoteBoardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BoardService_GetGameChallenges_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGameChallengesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoardServiceServer).GetGameChallenges(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tally.v1.BoardService/GetGameChallenges",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoardServiceServer).GetGameChallenges(ctx, req.(*GetGameChallengesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BoardService_CreateGameChallenge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGameChallengeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoardServiceServer).CreateGameChallenge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tally.v1.BoardService/CreateGameChallenge",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoardServiceServer).CreateGameChallenge(ctx, req.(*CreateGameChallengeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BoardService_ServiceDesc is the grpc.ServiceDesc for BoardService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BoardService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tally.v1.BoardService",
	HandlerType: (*BoardServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewGame",
			Handler:    _BoardService_NewGame_Handler,
		},
		{
			MethodName: "NewGameFromTemplate",
			Handler:    _BoardService_NewGameFromTemplate_Handler,
		},
		{
			MethodName: "GetHint",
			Handler:    _BoardService_GetHint_Handler,
		},
		{
			MethodName: "RestartGame",
			Handler:    _BoardService_RestartGame_Handler,
		},
		{
			MethodName: "GetSession",
			Handler:    _BoardService_GetSession_Handler,
		},
		{
			MethodName: "SwipeBoard",
			Handler:    _BoardService_SwipeBoard_Handler,
		},
		{
			MethodName: "CombineCells",
			Handler:    _BoardService_CombineCells_Handler,
		},
		{
			MethodName: "GenerateGame",
			Handler:    _BoardService_GenerateGame_Handler,
		},
		{
			MethodName: "VoteBoard",
			Handler:    _BoardService_VoteBoard_Handler,
		},
		{
			MethodName: "GetGameChallenges",
			Handler:    _BoardService_GetGameChallenges_Handler,
		},
		{
			MethodName: "CreateGameChallenge",
			Handler:    _BoardService_CreateGameChallenge_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/tally/v1/board.proto",
}
