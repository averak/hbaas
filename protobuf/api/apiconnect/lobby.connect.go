// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: api/lobby.proto

package apiconnect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	api "github.com/averak/hbaas/protobuf/api"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// LobbyServiceName is the fully-qualified name of the LobbyService service.
	LobbyServiceName = "api.LobbyService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// LobbyServiceSearchRoomsV1Procedure is the fully-qualified name of the LobbyService's
	// SearchRoomsV1 RPC.
	LobbyServiceSearchRoomsV1Procedure = "/api.LobbyService/SearchRoomsV1"
	// LobbyServiceCreateRoomV1Procedure is the fully-qualified name of the LobbyService's CreateRoomV1
	// RPC.
	LobbyServiceCreateRoomV1Procedure = "/api.LobbyService/CreateRoomV1"
	// LobbyServiceDeleteRoomV1Procedure is the fully-qualified name of the LobbyService's DeleteRoomV1
	// RPC.
	LobbyServiceDeleteRoomV1Procedure = "/api.LobbyService/DeleteRoomV1"
	// LobbyServiceJoinRoomV1Procedure is the fully-qualified name of the LobbyService's JoinRoomV1 RPC.
	LobbyServiceJoinRoomV1Procedure = "/api.LobbyService/JoinRoomV1"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	lobbyServiceServiceDescriptor             = api.File_api_lobby_proto.Services().ByName("LobbyService")
	lobbyServiceSearchRoomsV1MethodDescriptor = lobbyServiceServiceDescriptor.Methods().ByName("SearchRoomsV1")
	lobbyServiceCreateRoomV1MethodDescriptor  = lobbyServiceServiceDescriptor.Methods().ByName("CreateRoomV1")
	lobbyServiceDeleteRoomV1MethodDescriptor  = lobbyServiceServiceDescriptor.Methods().ByName("DeleteRoomV1")
	lobbyServiceJoinRoomV1MethodDescriptor    = lobbyServiceServiceDescriptor.Methods().ByName("JoinRoomV1")
)

// LobbyServiceClient is a client for the api.LobbyService service.
type LobbyServiceClient interface {
	SearchRoomsV1(context.Context, *connect.Request[api.LobbyServiceSearchRoomsV1Request]) (*connect.Response[api.LobbyServiceSearchRoomsV1Response], error)
	CreateRoomV1(context.Context, *connect.Request[api.LobbyServiceCreateRoomV1Request]) (*connect.Response[api.LobbyServiceCreateRoomV1Response], error)
	DeleteRoomV1(context.Context, *connect.Request[api.LobbyServiceDeleteRoomV1Request]) (*connect.Response[api.LobbyServiceDeleteRoomV1Response], error)
	JoinRoomV1(context.Context, *connect.Request[api.LobbyServiceJoinRoomV1Request]) (*connect.Response[api.LobbyServiceJoinRoomV1Response], error)
}

// NewLobbyServiceClient constructs a client for the api.LobbyService service. By default, it uses
// the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewLobbyServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) LobbyServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &lobbyServiceClient{
		searchRoomsV1: connect.NewClient[api.LobbyServiceSearchRoomsV1Request, api.LobbyServiceSearchRoomsV1Response](
			httpClient,
			baseURL+LobbyServiceSearchRoomsV1Procedure,
			connect.WithSchema(lobbyServiceSearchRoomsV1MethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		createRoomV1: connect.NewClient[api.LobbyServiceCreateRoomV1Request, api.LobbyServiceCreateRoomV1Response](
			httpClient,
			baseURL+LobbyServiceCreateRoomV1Procedure,
			connect.WithSchema(lobbyServiceCreateRoomV1MethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		deleteRoomV1: connect.NewClient[api.LobbyServiceDeleteRoomV1Request, api.LobbyServiceDeleteRoomV1Response](
			httpClient,
			baseURL+LobbyServiceDeleteRoomV1Procedure,
			connect.WithSchema(lobbyServiceDeleteRoomV1MethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		joinRoomV1: connect.NewClient[api.LobbyServiceJoinRoomV1Request, api.LobbyServiceJoinRoomV1Response](
			httpClient,
			baseURL+LobbyServiceJoinRoomV1Procedure,
			connect.WithSchema(lobbyServiceJoinRoomV1MethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// lobbyServiceClient implements LobbyServiceClient.
type lobbyServiceClient struct {
	searchRoomsV1 *connect.Client[api.LobbyServiceSearchRoomsV1Request, api.LobbyServiceSearchRoomsV1Response]
	createRoomV1  *connect.Client[api.LobbyServiceCreateRoomV1Request, api.LobbyServiceCreateRoomV1Response]
	deleteRoomV1  *connect.Client[api.LobbyServiceDeleteRoomV1Request, api.LobbyServiceDeleteRoomV1Response]
	joinRoomV1    *connect.Client[api.LobbyServiceJoinRoomV1Request, api.LobbyServiceJoinRoomV1Response]
}

// SearchRoomsV1 calls api.LobbyService.SearchRoomsV1.
func (c *lobbyServiceClient) SearchRoomsV1(ctx context.Context, req *connect.Request[api.LobbyServiceSearchRoomsV1Request]) (*connect.Response[api.LobbyServiceSearchRoomsV1Response], error) {
	return c.searchRoomsV1.CallUnary(ctx, req)
}

// CreateRoomV1 calls api.LobbyService.CreateRoomV1.
func (c *lobbyServiceClient) CreateRoomV1(ctx context.Context, req *connect.Request[api.LobbyServiceCreateRoomV1Request]) (*connect.Response[api.LobbyServiceCreateRoomV1Response], error) {
	return c.createRoomV1.CallUnary(ctx, req)
}

// DeleteRoomV1 calls api.LobbyService.DeleteRoomV1.
func (c *lobbyServiceClient) DeleteRoomV1(ctx context.Context, req *connect.Request[api.LobbyServiceDeleteRoomV1Request]) (*connect.Response[api.LobbyServiceDeleteRoomV1Response], error) {
	return c.deleteRoomV1.CallUnary(ctx, req)
}

// JoinRoomV1 calls api.LobbyService.JoinRoomV1.
func (c *lobbyServiceClient) JoinRoomV1(ctx context.Context, req *connect.Request[api.LobbyServiceJoinRoomV1Request]) (*connect.Response[api.LobbyServiceJoinRoomV1Response], error) {
	return c.joinRoomV1.CallUnary(ctx, req)
}

// LobbyServiceHandler is an implementation of the api.LobbyService service.
type LobbyServiceHandler interface {
	SearchRoomsV1(context.Context, *connect.Request[api.LobbyServiceSearchRoomsV1Request]) (*connect.Response[api.LobbyServiceSearchRoomsV1Response], error)
	CreateRoomV1(context.Context, *connect.Request[api.LobbyServiceCreateRoomV1Request]) (*connect.Response[api.LobbyServiceCreateRoomV1Response], error)
	DeleteRoomV1(context.Context, *connect.Request[api.LobbyServiceDeleteRoomV1Request]) (*connect.Response[api.LobbyServiceDeleteRoomV1Response], error)
	JoinRoomV1(context.Context, *connect.Request[api.LobbyServiceJoinRoomV1Request]) (*connect.Response[api.LobbyServiceJoinRoomV1Response], error)
}

// NewLobbyServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewLobbyServiceHandler(svc LobbyServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	lobbyServiceSearchRoomsV1Handler := connect.NewUnaryHandler(
		LobbyServiceSearchRoomsV1Procedure,
		svc.SearchRoomsV1,
		connect.WithSchema(lobbyServiceSearchRoomsV1MethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	lobbyServiceCreateRoomV1Handler := connect.NewUnaryHandler(
		LobbyServiceCreateRoomV1Procedure,
		svc.CreateRoomV1,
		connect.WithSchema(lobbyServiceCreateRoomV1MethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	lobbyServiceDeleteRoomV1Handler := connect.NewUnaryHandler(
		LobbyServiceDeleteRoomV1Procedure,
		svc.DeleteRoomV1,
		connect.WithSchema(lobbyServiceDeleteRoomV1MethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	lobbyServiceJoinRoomV1Handler := connect.NewUnaryHandler(
		LobbyServiceJoinRoomV1Procedure,
		svc.JoinRoomV1,
		connect.WithSchema(lobbyServiceJoinRoomV1MethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/api.LobbyService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case LobbyServiceSearchRoomsV1Procedure:
			lobbyServiceSearchRoomsV1Handler.ServeHTTP(w, r)
		case LobbyServiceCreateRoomV1Procedure:
			lobbyServiceCreateRoomV1Handler.ServeHTTP(w, r)
		case LobbyServiceDeleteRoomV1Procedure:
			lobbyServiceDeleteRoomV1Handler.ServeHTTP(w, r)
		case LobbyServiceJoinRoomV1Procedure:
			lobbyServiceJoinRoomV1Handler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedLobbyServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedLobbyServiceHandler struct{}

func (UnimplementedLobbyServiceHandler) SearchRoomsV1(context.Context, *connect.Request[api.LobbyServiceSearchRoomsV1Request]) (*connect.Response[api.LobbyServiceSearchRoomsV1Response], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.LobbyService.SearchRoomsV1 is not implemented"))
}

func (UnimplementedLobbyServiceHandler) CreateRoomV1(context.Context, *connect.Request[api.LobbyServiceCreateRoomV1Request]) (*connect.Response[api.LobbyServiceCreateRoomV1Response], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.LobbyService.CreateRoomV1 is not implemented"))
}

func (UnimplementedLobbyServiceHandler) DeleteRoomV1(context.Context, *connect.Request[api.LobbyServiceDeleteRoomV1Request]) (*connect.Response[api.LobbyServiceDeleteRoomV1Response], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.LobbyService.DeleteRoomV1 is not implemented"))
}

func (UnimplementedLobbyServiceHandler) JoinRoomV1(context.Context, *connect.Request[api.LobbyServiceJoinRoomV1Request]) (*connect.Response[api.LobbyServiceJoinRoomV1Response], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.LobbyService.JoinRoomV1 is not implemented"))
}