// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: proto/gossip/v1/gossip.proto

package gossipv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/getghost/ghost/apps/agent/gen/proto/gossip/v1"
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
	// GossipServiceName is the fully-qualified name of the GossipService service.
	GossipServiceName = "gossip.v1.GossipService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// GossipServicePingProcedure is the fully-qualified name of the GossipService's Ping RPC.
	GossipServicePingProcedure = "/gossip.v1.GossipService/Ping"
	// GossipServiceIndirectPingProcedure is the fully-qualified name of the GossipService's
	// IndirectPing RPC.
	GossipServiceIndirectPingProcedure = "/gossip.v1.GossipService/IndirectPing"
	// GossipServiceSyncMembersProcedure is the fully-qualified name of the GossipService's SyncMembers
	// RPC.
	GossipServiceSyncMembersProcedure = "/gossip.v1.GossipService/SyncMembers"
	// GossipServiceJoinProcedure is the fully-qualified name of the GossipService's Join RPC.
	GossipServiceJoinProcedure = "/gossip.v1.GossipService/Join"
	// GossipServiceLeaveProcedure is the fully-qualified name of the GossipService's Leave RPC.
	GossipServiceLeaveProcedure = "/gossip.v1.GossipService/Leave"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	gossipServiceServiceDescriptor            = v1.File_proto_gossip_v1_gossip_proto.Services().ByName("GossipService")
	gossipServicePingMethodDescriptor         = gossipServiceServiceDescriptor.Methods().ByName("Ping")
	gossipServiceIndirectPingMethodDescriptor = gossipServiceServiceDescriptor.Methods().ByName("IndirectPing")
	gossipServiceSyncMembersMethodDescriptor  = gossipServiceServiceDescriptor.Methods().ByName("SyncMembers")
	gossipServiceJoinMethodDescriptor         = gossipServiceServiceDescriptor.Methods().ByName("Join")
	gossipServiceLeaveMethodDescriptor        = gossipServiceServiceDescriptor.Methods().ByName("Leave")
)

// GossipServiceClient is a client for the gossip.v1.GossipService service.
type GossipServiceClient interface {
	// Ping asks for the state of a peer
	// If the peer is healthy, it should respond with its state
	Ping(context.Context, *connect.Request[v1.PingRequest]) (*connect.Response[v1.PingResponse], error)
	// IndirectPing asks a peer to ping another node because we can not reach it outselves
	// the peer should respond with the state of the node
	IndirectPing(context.Context, *connect.Request[v1.IndirectPingRequest]) (*connect.Response[v1.IndirectPingResponse], error)
	// Periodially we do a full sync of the members
	// Both nodes tell each other about every member they know and then reconcile by taking the union
	// of the two sets.
	// Afterwards, both nodes should have the same view of the cluster and regular gossip will get rid
	// of any dead nodes
	//
	// If they disagree on the state of a node, the most favourable state should be chosen
	// ie: if one node thinks a peer is dead and the other thinks it is alive, the node should be
	// marked as alive to prevent a split brain or unnecessary false positives
	SyncMembers(context.Context, *connect.Request[v1.SyncMembersRequest]) (*connect.Response[v1.SyncMembersResponse], error)
	// Join allows a node to advertise itself to the cluster
	// The node sends their own information, so the cluster may add them to the list of known members
	// The cluster responds with the list of known members to bootstrap the new node
	//
	// It's sufficient to call join on one node, the rest of the cluster will be updated through
	// gossip, however it is recommended to call join on multiple nodes to ensure the information is
	// propagated quickly and to minimize the chance of a single node failing before propagating the
	// information.
	Join(context.Context, *connect.Request[v1.JoinRequest]) (*connect.Response[v1.JoinResponse], error)
	// Leave should be broadcasted to all nodes in the cluster when a node is leaving for any reason.
	Leave(context.Context, *connect.Request[v1.LeaveRequest]) (*connect.Response[v1.LeaveResponse], error)
}

// NewGossipServiceClient constructs a client for the gossip.v1.GossipService service. By default,
// it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and
// sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC()
// or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewGossipServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) GossipServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &gossipServiceClient{
		ping: connect.NewClient[v1.PingRequest, v1.PingResponse](
			httpClient,
			baseURL+GossipServicePingProcedure,
			connect.WithSchema(gossipServicePingMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		indirectPing: connect.NewClient[v1.IndirectPingRequest, v1.IndirectPingResponse](
			httpClient,
			baseURL+GossipServiceIndirectPingProcedure,
			connect.WithSchema(gossipServiceIndirectPingMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		syncMembers: connect.NewClient[v1.SyncMembersRequest, v1.SyncMembersResponse](
			httpClient,
			baseURL+GossipServiceSyncMembersProcedure,
			connect.WithSchema(gossipServiceSyncMembersMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		join: connect.NewClient[v1.JoinRequest, v1.JoinResponse](
			httpClient,
			baseURL+GossipServiceJoinProcedure,
			connect.WithSchema(gossipServiceJoinMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		leave: connect.NewClient[v1.LeaveRequest, v1.LeaveResponse](
			httpClient,
			baseURL+GossipServiceLeaveProcedure,
			connect.WithSchema(gossipServiceLeaveMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// gossipServiceClient implements GossipServiceClient.
type gossipServiceClient struct {
	ping         *connect.Client[v1.PingRequest, v1.PingResponse]
	indirectPing *connect.Client[v1.IndirectPingRequest, v1.IndirectPingResponse]
	syncMembers  *connect.Client[v1.SyncMembersRequest, v1.SyncMembersResponse]
	join         *connect.Client[v1.JoinRequest, v1.JoinResponse]
	leave        *connect.Client[v1.LeaveRequest, v1.LeaveResponse]
}

// Ping calls gossip.v1.GossipService.Ping.
func (c *gossipServiceClient) Ping(ctx context.Context, req *connect.Request[v1.PingRequest]) (*connect.Response[v1.PingResponse], error) {
	return c.ping.CallUnary(ctx, req)
}

// IndirectPing calls gossip.v1.GossipService.IndirectPing.
func (c *gossipServiceClient) IndirectPing(ctx context.Context, req *connect.Request[v1.IndirectPingRequest]) (*connect.Response[v1.IndirectPingResponse], error) {
	return c.indirectPing.CallUnary(ctx, req)
}

// SyncMembers calls gossip.v1.GossipService.SyncMembers.
func (c *gossipServiceClient) SyncMembers(ctx context.Context, req *connect.Request[v1.SyncMembersRequest]) (*connect.Response[v1.SyncMembersResponse], error) {
	return c.syncMembers.CallUnary(ctx, req)
}

// Join calls gossip.v1.GossipService.Join.
func (c *gossipServiceClient) Join(ctx context.Context, req *connect.Request[v1.JoinRequest]) (*connect.Response[v1.JoinResponse], error) {
	return c.join.CallUnary(ctx, req)
}

// Leave calls gossip.v1.GossipService.Leave.
func (c *gossipServiceClient) Leave(ctx context.Context, req *connect.Request[v1.LeaveRequest]) (*connect.Response[v1.LeaveResponse], error) {
	return c.leave.CallUnary(ctx, req)
}

// GossipServiceHandler is an implementation of the gossip.v1.GossipService service.
type GossipServiceHandler interface {
	// Ping asks for the state of a peer
	// If the peer is healthy, it should respond with its state
	Ping(context.Context, *connect.Request[v1.PingRequest]) (*connect.Response[v1.PingResponse], error)
	// IndirectPing asks a peer to ping another node because we can not reach it outselves
	// the peer should respond with the state of the node
	IndirectPing(context.Context, *connect.Request[v1.IndirectPingRequest]) (*connect.Response[v1.IndirectPingResponse], error)
	// Periodially we do a full sync of the members
	// Both nodes tell each other about every member they know and then reconcile by taking the union
	// of the two sets.
	// Afterwards, both nodes should have the same view of the cluster and regular gossip will get rid
	// of any dead nodes
	//
	// If they disagree on the state of a node, the most favourable state should be chosen
	// ie: if one node thinks a peer is dead and the other thinks it is alive, the node should be
	// marked as alive to prevent a split brain or unnecessary false positives
	SyncMembers(context.Context, *connect.Request[v1.SyncMembersRequest]) (*connect.Response[v1.SyncMembersResponse], error)
	// Join allows a node to advertise itself to the cluster
	// The node sends their own information, so the cluster may add them to the list of known members
	// The cluster responds with the list of known members to bootstrap the new node
	//
	// It's sufficient to call join on one node, the rest of the cluster will be updated through
	// gossip, however it is recommended to call join on multiple nodes to ensure the information is
	// propagated quickly and to minimize the chance of a single node failing before propagating the
	// information.
	Join(context.Context, *connect.Request[v1.JoinRequest]) (*connect.Response[v1.JoinResponse], error)
	// Leave should be broadcasted to all nodes in the cluster when a node is leaving for any reason.
	Leave(context.Context, *connect.Request[v1.LeaveRequest]) (*connect.Response[v1.LeaveResponse], error)
}

// NewGossipServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewGossipServiceHandler(svc GossipServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	gossipServicePingHandler := connect.NewUnaryHandler(
		GossipServicePingProcedure,
		svc.Ping,
		connect.WithSchema(gossipServicePingMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	gossipServiceIndirectPingHandler := connect.NewUnaryHandler(
		GossipServiceIndirectPingProcedure,
		svc.IndirectPing,
		connect.WithSchema(gossipServiceIndirectPingMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	gossipServiceSyncMembersHandler := connect.NewUnaryHandler(
		GossipServiceSyncMembersProcedure,
		svc.SyncMembers,
		connect.WithSchema(gossipServiceSyncMembersMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	gossipServiceJoinHandler := connect.NewUnaryHandler(
		GossipServiceJoinProcedure,
		svc.Join,
		connect.WithSchema(gossipServiceJoinMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	gossipServiceLeaveHandler := connect.NewUnaryHandler(
		GossipServiceLeaveProcedure,
		svc.Leave,
		connect.WithSchema(gossipServiceLeaveMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/gossip.v1.GossipService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case GossipServicePingProcedure:
			gossipServicePingHandler.ServeHTTP(w, r)
		case GossipServiceIndirectPingProcedure:
			gossipServiceIndirectPingHandler.ServeHTTP(w, r)
		case GossipServiceSyncMembersProcedure:
			gossipServiceSyncMembersHandler.ServeHTTP(w, r)
		case GossipServiceJoinProcedure:
			gossipServiceJoinHandler.ServeHTTP(w, r)
		case GossipServiceLeaveProcedure:
			gossipServiceLeaveHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedGossipServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedGossipServiceHandler struct{}

func (UnimplementedGossipServiceHandler) Ping(context.Context, *connect.Request[v1.PingRequest]) (*connect.Response[v1.PingResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("gossip.v1.GossipService.Ping is not implemented"))
}

func (UnimplementedGossipServiceHandler) IndirectPing(context.Context, *connect.Request[v1.IndirectPingRequest]) (*connect.Response[v1.IndirectPingResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("gossip.v1.GossipService.IndirectPing is not implemented"))
}

func (UnimplementedGossipServiceHandler) SyncMembers(context.Context, *connect.Request[v1.SyncMembersRequest]) (*connect.Response[v1.SyncMembersResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("gossip.v1.GossipService.SyncMembers is not implemented"))
}

func (UnimplementedGossipServiceHandler) Join(context.Context, *connect.Request[v1.JoinRequest]) (*connect.Response[v1.JoinResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("gossip.v1.GossipService.Join is not implemented"))
}

func (UnimplementedGossipServiceHandler) Leave(context.Context, *connect.Request[v1.LeaveRequest]) (*connect.Response[v1.LeaveResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("gossip.v1.GossipService.Leave is not implemented"))
}
