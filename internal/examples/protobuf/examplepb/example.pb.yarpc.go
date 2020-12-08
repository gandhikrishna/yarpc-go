// Code generated by protoc-gen-yarpc-go. DO NOT EDIT.
// source: example.proto

package examplepb

import (
	"context"
	"io/ioutil"
	"reflect"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"go.uber.org/fx"
	"go.uber.org/yarpc"
	"go.uber.org/yarpc/api/transport"
	"go.uber.org/yarpc/api/x/restriction"
	"go.uber.org/yarpc/encoding/protobuf"
	"go.uber.org/yarpc/encoding/protobuf/reflection"
)

var _ = ioutil.NopCloser

// KeyValueYARPCClient is the YARPC client-side interface for the KeyValue service.
type KeyValueYARPCClient interface {
	GetValue(context.Context, *GetValueRequest, ...yarpc.CallOption) (*GetValueResponse, error)
	SetValue(context.Context, *SetValueRequest, ...yarpc.CallOption) (*SetValueResponse, error)
}

func newKeyValueYARPCClient(clientConfig transport.ClientConfig, anyResolver jsonpb.AnyResolver, options ...protobuf.ClientOption) KeyValueYARPCClient {
	return &_KeyValueYARPCCaller{protobuf.NewStreamClient(
		protobuf.ClientParams{
			ServiceName:  "uber.yarpc.internal.examples.protobuf.example.KeyValue",
			ClientConfig: clientConfig,
			AnyResolver:  anyResolver,
			Options:      options,
		},
	)}
}

// NewKeyValueYARPCClient builds a new YARPC client for the KeyValue service.
func NewKeyValueYARPCClient(clientConfig transport.ClientConfig, options ...protobuf.ClientOption) KeyValueYARPCClient {
	return newKeyValueYARPCClient(clientConfig, nil, options...)
}

// KeyValueYARPCServer is the YARPC server-side interface for the KeyValue service.
type KeyValueYARPCServer interface {
	GetValue(context.Context, *GetValueRequest) (*GetValueResponse, error)
	SetValue(context.Context, *SetValueRequest) (*SetValueResponse, error)
}

type buildKeyValueYARPCProceduresParams struct {
	Server      KeyValueYARPCServer
	AnyResolver jsonpb.AnyResolver
}

func buildKeyValueYARPCProcedures(params buildKeyValueYARPCProceduresParams) []transport.Procedure {
	handler := &_KeyValueYARPCHandler{params.Server}
	return protobuf.BuildProcedures(
		protobuf.BuildProceduresParams{
			ServiceName: "uber.yarpc.internal.examples.protobuf.example.KeyValue",
			UnaryHandlerParams: []protobuf.BuildProceduresUnaryHandlerParams{
				{
					MethodName: "GetValue",
					Handler: protobuf.NewUnaryHandler(
						protobuf.UnaryHandlerParams{
							Handle:      handler.GetValue,
							NewRequest:  newKeyValueServiceGetValueYARPCRequest,
							AnyResolver: params.AnyResolver,
						},
					),
				},
				{
					MethodName: "SetValue",
					Handler: protobuf.NewUnaryHandler(
						protobuf.UnaryHandlerParams{
							Handle:      handler.SetValue,
							NewRequest:  newKeyValueServiceSetValueYARPCRequest,
							AnyResolver: params.AnyResolver,
						},
					),
				},
			},
			OnewayHandlerParams: []protobuf.BuildProceduresOnewayHandlerParams{},
			StreamHandlerParams: []protobuf.BuildProceduresStreamHandlerParams{},
		},
	)
}

// BuildKeyValueYARPCProcedures prepares an implementation of the KeyValue service for YARPC registration.
func BuildKeyValueYARPCProcedures(server KeyValueYARPCServer) []transport.Procedure {
	return buildKeyValueYARPCProcedures(buildKeyValueYARPCProceduresParams{Server: server})
}

// FxKeyValueYARPCClientParams defines the input
// for NewFxKeyValueYARPCClient. It provides the
// paramaters to get a KeyValueYARPCClient in an
// Fx application.
type FxKeyValueYARPCClientParams struct {
	fx.In

	Provider    yarpc.ClientConfig
	AnyResolver jsonpb.AnyResolver  `name:"yarpcfx" optional:"true"`
	Restriction restriction.Checker `optional:"true"`
}

// FxKeyValueYARPCClientResult defines the output
// of NewFxKeyValueYARPCClient. It provides a
// KeyValueYARPCClient to an Fx application.
type FxKeyValueYARPCClientResult struct {
	fx.Out

	Client KeyValueYARPCClient

	// We are using an fx.Out struct here instead of just returning a client
	// so that we can add more values or add named versions of the client in
	// the future without breaking any existing code.
}

// NewFxKeyValueYARPCClient provides a KeyValueYARPCClient
// to an Fx application using the given name for routing.
//
//  fx.Provide(
//    examplepb.NewFxKeyValueYARPCClient("service-name"),
//    ...
//  )
func NewFxKeyValueYARPCClient(name string, options ...protobuf.ClientOption) interface{} {
	return func(params FxKeyValueYARPCClientParams) FxKeyValueYARPCClientResult {
		cc := params.Provider.ClientConfig(name)

		if params.Restriction != nil {
			if namer, ok := cc.GetUnaryOutbound().(transport.Namer); ok {
				if err := params.Restriction.Check(protobuf.Encoding, namer.TransportName()); err != nil {
					panic(err.Error())
				}
			}
		}

		return FxKeyValueYARPCClientResult{
			Client: newKeyValueYARPCClient(cc, params.AnyResolver, options...),
		}
	}
}

// FxKeyValueYARPCProceduresParams defines the input
// for NewFxKeyValueYARPCProcedures. It provides the
// paramaters to get KeyValueYARPCServer procedures in an
// Fx application.
type FxKeyValueYARPCProceduresParams struct {
	fx.In

	Server      KeyValueYARPCServer
	AnyResolver jsonpb.AnyResolver `name:"yarpcfx" optional:"true"`
}

// FxKeyValueYARPCProceduresResult defines the output
// of NewFxKeyValueYARPCProcedures. It provides
// KeyValueYARPCServer procedures to an Fx application.
//
// The procedures are provided to the "yarpcfx" value group.
// Dig 1.2 or newer must be used for this feature to work.
type FxKeyValueYARPCProceduresResult struct {
	fx.Out

	Procedures     []transport.Procedure `group:"yarpcfx"`
	ReflectionMeta reflection.ServerMeta `group:"yarpcfx"`
}

// NewFxKeyValueYARPCProcedures provides KeyValueYARPCServer procedures to an Fx application.
// It expects a KeyValueYARPCServer to be present in the container.
//
//  fx.Provide(
//    examplepb.NewFxKeyValueYARPCProcedures(),
//    ...
//  )
func NewFxKeyValueYARPCProcedures() interface{} {
	return func(params FxKeyValueYARPCProceduresParams) FxKeyValueYARPCProceduresResult {
		return FxKeyValueYARPCProceduresResult{
			Procedures: buildKeyValueYARPCProcedures(buildKeyValueYARPCProceduresParams{
				Server:      params.Server,
				AnyResolver: params.AnyResolver,
			}),
			ReflectionMeta: reflection.ServerMeta{
				ServiceName:     "uber.yarpc.internal.examples.protobuf.example.KeyValue",
				FileDescriptors: yarpcFileDescriptorClosure15a1dc8d40dadaa6,
			},
		}
	}
}

type _KeyValueYARPCCaller struct {
	streamClient protobuf.StreamClient
}

func (c *_KeyValueYARPCCaller) GetValue(ctx context.Context, request *GetValueRequest, options ...yarpc.CallOption) (*GetValueResponse, error) {
	responseMessage, err := c.streamClient.Call(ctx, "GetValue", request, newKeyValueServiceGetValueYARPCResponse, options...)
	if responseMessage == nil {
		return nil, err
	}
	response, ok := responseMessage.(*GetValueResponse)
	if !ok {
		return nil, protobuf.CastError(emptyKeyValueServiceGetValueYARPCResponse, responseMessage)
	}
	return response, err
}

func (c *_KeyValueYARPCCaller) SetValue(ctx context.Context, request *SetValueRequest, options ...yarpc.CallOption) (*SetValueResponse, error) {
	responseMessage, err := c.streamClient.Call(ctx, "SetValue", request, newKeyValueServiceSetValueYARPCResponse, options...)
	if responseMessage == nil {
		return nil, err
	}
	response, ok := responseMessage.(*SetValueResponse)
	if !ok {
		return nil, protobuf.CastError(emptyKeyValueServiceSetValueYARPCResponse, responseMessage)
	}
	return response, err
}

type _KeyValueYARPCHandler struct {
	server KeyValueYARPCServer
}

func (h *_KeyValueYARPCHandler) GetValue(ctx context.Context, requestMessage proto.Message) (proto.Message, error) {
	var request *GetValueRequest
	var ok bool
	if requestMessage != nil {
		request, ok = requestMessage.(*GetValueRequest)
		if !ok {
			return nil, protobuf.CastError(emptyKeyValueServiceGetValueYARPCRequest, requestMessage)
		}
	}
	response, err := h.server.GetValue(ctx, request)
	if response == nil {
		return nil, err
	}
	return response, err
}

func (h *_KeyValueYARPCHandler) SetValue(ctx context.Context, requestMessage proto.Message) (proto.Message, error) {
	var request *SetValueRequest
	var ok bool
	if requestMessage != nil {
		request, ok = requestMessage.(*SetValueRequest)
		if !ok {
			return nil, protobuf.CastError(emptyKeyValueServiceSetValueYARPCRequest, requestMessage)
		}
	}
	response, err := h.server.SetValue(ctx, request)
	if response == nil {
		return nil, err
	}
	return response, err
}

func newKeyValueServiceGetValueYARPCRequest() proto.Message {
	return &GetValueRequest{}
}

func newKeyValueServiceGetValueYARPCResponse() proto.Message {
	return &GetValueResponse{}
}

func newKeyValueServiceSetValueYARPCRequest() proto.Message {
	return &SetValueRequest{}
}

func newKeyValueServiceSetValueYARPCResponse() proto.Message {
	return &SetValueResponse{}
}

var (
	emptyKeyValueServiceGetValueYARPCRequest  = &GetValueRequest{}
	emptyKeyValueServiceGetValueYARPCResponse = &GetValueResponse{}
	emptyKeyValueServiceSetValueYARPCRequest  = &SetValueRequest{}
	emptyKeyValueServiceSetValueYARPCResponse = &SetValueResponse{}
)

// FooYARPCClient is the YARPC client-side interface for the Foo service.
type FooYARPCClient interface {
	EchoOut(context.Context, ...yarpc.CallOption) (FooServiceEchoOutYARPCClient, error)
	EchoIn(context.Context, *EchoInRequest, ...yarpc.CallOption) (FooServiceEchoInYARPCClient, error)
	EchoBoth(context.Context, ...yarpc.CallOption) (FooServiceEchoBothYARPCClient, error)
}

// FooServiceEchoOutYARPCClient sends EchoOutRequests and receives the single EchoOutResponse when sending is done.
type FooServiceEchoOutYARPCClient interface {
	Context() context.Context
	Send(*EchoOutRequest, ...yarpc.StreamOption) error
	CloseAndRecv(...yarpc.StreamOption) (*EchoOutResponse, error)
}

// FooServiceEchoInYARPCClient receives EchoInResponses, returning io.EOF when the stream is complete.
type FooServiceEchoInYARPCClient interface {
	Context() context.Context
	Recv(...yarpc.StreamOption) (*EchoInResponse, error)
	CloseSend(...yarpc.StreamOption) error
}

// FooServiceEchoBothYARPCClient sends EchoBothRequests and receives EchoBothResponses, returning io.EOF when the stream is complete.
type FooServiceEchoBothYARPCClient interface {
	Context() context.Context
	Send(*EchoBothRequest, ...yarpc.StreamOption) error
	Recv(...yarpc.StreamOption) (*EchoBothResponse, error)
	CloseSend(...yarpc.StreamOption) error
}

func newFooYARPCClient(clientConfig transport.ClientConfig, anyResolver jsonpb.AnyResolver, options ...protobuf.ClientOption) FooYARPCClient {
	return &_FooYARPCCaller{protobuf.NewStreamClient(
		protobuf.ClientParams{
			ServiceName:  "uber.yarpc.internal.examples.protobuf.example.Foo",
			ClientConfig: clientConfig,
			AnyResolver:  anyResolver,
			Options:      options,
		},
	)}
}

// NewFooYARPCClient builds a new YARPC client for the Foo service.
func NewFooYARPCClient(clientConfig transport.ClientConfig, options ...protobuf.ClientOption) FooYARPCClient {
	return newFooYARPCClient(clientConfig, nil, options...)
}

// FooYARPCServer is the YARPC server-side interface for the Foo service.
type FooYARPCServer interface {
	EchoOut(FooServiceEchoOutYARPCServer) (*EchoOutResponse, error)
	EchoIn(*EchoInRequest, FooServiceEchoInYARPCServer) error
	EchoBoth(FooServiceEchoBothYARPCServer) error
}

// FooServiceEchoOutYARPCServer receives EchoOutRequests.
type FooServiceEchoOutYARPCServer interface {
	Context() context.Context
	Recv(...yarpc.StreamOption) (*EchoOutRequest, error)
}

// FooServiceEchoInYARPCServer sends EchoInResponses.
type FooServiceEchoInYARPCServer interface {
	Context() context.Context
	Send(*EchoInResponse, ...yarpc.StreamOption) error
}

// FooServiceEchoBothYARPCServer receives EchoBothRequests and sends EchoBothResponse.
type FooServiceEchoBothYARPCServer interface {
	Context() context.Context
	Recv(...yarpc.StreamOption) (*EchoBothRequest, error)
	Send(*EchoBothResponse, ...yarpc.StreamOption) error
}

type buildFooYARPCProceduresParams struct {
	Server      FooYARPCServer
	AnyResolver jsonpb.AnyResolver
}

func buildFooYARPCProcedures(params buildFooYARPCProceduresParams) []transport.Procedure {
	handler := &_FooYARPCHandler{params.Server}
	return protobuf.BuildProcedures(
		protobuf.BuildProceduresParams{
			ServiceName:         "uber.yarpc.internal.examples.protobuf.example.Foo",
			UnaryHandlerParams:  []protobuf.BuildProceduresUnaryHandlerParams{},
			OnewayHandlerParams: []protobuf.BuildProceduresOnewayHandlerParams{},
			StreamHandlerParams: []protobuf.BuildProceduresStreamHandlerParams{
				{
					MethodName: "EchoBoth",
					Handler: protobuf.NewStreamHandler(
						protobuf.StreamHandlerParams{
							Handle: handler.EchoBoth,
						},
					),
				},

				{
					MethodName: "EchoIn",
					Handler: protobuf.NewStreamHandler(
						protobuf.StreamHandlerParams{
							Handle: handler.EchoIn,
						},
					),
				},

				{
					MethodName: "EchoOut",
					Handler: protobuf.NewStreamHandler(
						protobuf.StreamHandlerParams{
							Handle: handler.EchoOut,
						},
					),
				},
			},
		},
	)
}

// BuildFooYARPCProcedures prepares an implementation of the Foo service for YARPC registration.
func BuildFooYARPCProcedures(server FooYARPCServer) []transport.Procedure {
	return buildFooYARPCProcedures(buildFooYARPCProceduresParams{Server: server})
}

// FxFooYARPCClientParams defines the input
// for NewFxFooYARPCClient. It provides the
// paramaters to get a FooYARPCClient in an
// Fx application.
type FxFooYARPCClientParams struct {
	fx.In

	Provider    yarpc.ClientConfig
	AnyResolver jsonpb.AnyResolver  `name:"yarpcfx" optional:"true"`
	Restriction restriction.Checker `optional:"true"`
}

// FxFooYARPCClientResult defines the output
// of NewFxFooYARPCClient. It provides a
// FooYARPCClient to an Fx application.
type FxFooYARPCClientResult struct {
	fx.Out

	Client FooYARPCClient

	// We are using an fx.Out struct here instead of just returning a client
	// so that we can add more values or add named versions of the client in
	// the future without breaking any existing code.
}

// NewFxFooYARPCClient provides a FooYARPCClient
// to an Fx application using the given name for routing.
//
//  fx.Provide(
//    examplepb.NewFxFooYARPCClient("service-name"),
//    ...
//  )
func NewFxFooYARPCClient(name string, options ...protobuf.ClientOption) interface{} {
	return func(params FxFooYARPCClientParams) FxFooYARPCClientResult {
		cc := params.Provider.ClientConfig(name)

		if params.Restriction != nil {
			if namer, ok := cc.GetUnaryOutbound().(transport.Namer); ok {
				if err := params.Restriction.Check(protobuf.Encoding, namer.TransportName()); err != nil {
					panic(err.Error())
				}
			}
		}

		return FxFooYARPCClientResult{
			Client: newFooYARPCClient(cc, params.AnyResolver, options...),
		}
	}
}

// FxFooYARPCProceduresParams defines the input
// for NewFxFooYARPCProcedures. It provides the
// paramaters to get FooYARPCServer procedures in an
// Fx application.
type FxFooYARPCProceduresParams struct {
	fx.In

	Server      FooYARPCServer
	AnyResolver jsonpb.AnyResolver `name:"yarpcfx" optional:"true"`
}

// FxFooYARPCProceduresResult defines the output
// of NewFxFooYARPCProcedures. It provides
// FooYARPCServer procedures to an Fx application.
//
// The procedures are provided to the "yarpcfx" value group.
// Dig 1.2 or newer must be used for this feature to work.
type FxFooYARPCProceduresResult struct {
	fx.Out

	Procedures     []transport.Procedure `group:"yarpcfx"`
	ReflectionMeta reflection.ServerMeta `group:"yarpcfx"`
}

// NewFxFooYARPCProcedures provides FooYARPCServer procedures to an Fx application.
// It expects a FooYARPCServer to be present in the container.
//
//  fx.Provide(
//    examplepb.NewFxFooYARPCProcedures(),
//    ...
//  )
func NewFxFooYARPCProcedures() interface{} {
	return func(params FxFooYARPCProceduresParams) FxFooYARPCProceduresResult {
		return FxFooYARPCProceduresResult{
			Procedures: buildFooYARPCProcedures(buildFooYARPCProceduresParams{
				Server:      params.Server,
				AnyResolver: params.AnyResolver,
			}),
			ReflectionMeta: reflection.ServerMeta{
				ServiceName:     "uber.yarpc.internal.examples.protobuf.example.Foo",
				FileDescriptors: yarpcFileDescriptorClosure15a1dc8d40dadaa6,
			},
		}
	}
}

type _FooYARPCCaller struct {
	streamClient protobuf.StreamClient
}

func (c *_FooYARPCCaller) EchoOut(ctx context.Context, options ...yarpc.CallOption) (FooServiceEchoOutYARPCClient, error) {
	stream, err := c.streamClient.CallStream(ctx, "EchoOut", options...)
	if err != nil {
		return nil, err
	}
	return &_FooServiceEchoOutYARPCClient{stream: stream}, nil
}

func (c *_FooYARPCCaller) EchoIn(ctx context.Context, request *EchoInRequest, options ...yarpc.CallOption) (FooServiceEchoInYARPCClient, error) {
	stream, err := c.streamClient.CallStream(ctx, "EchoIn", options...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(request); err != nil {
		return nil, err
	}
	return &_FooServiceEchoInYARPCClient{stream: stream}, nil
}

func (c *_FooYARPCCaller) EchoBoth(ctx context.Context, options ...yarpc.CallOption) (FooServiceEchoBothYARPCClient, error) {
	stream, err := c.streamClient.CallStream(ctx, "EchoBoth", options...)
	if err != nil {
		return nil, err
	}
	return &_FooServiceEchoBothYARPCClient{stream: stream}, nil
}

type _FooYARPCHandler struct {
	server FooYARPCServer
}

func (h *_FooYARPCHandler) EchoOut(serverStream *protobuf.ServerStream) error {
	response, err := h.server.EchoOut(&_FooServiceEchoOutYARPCServer{serverStream: serverStream})
	if err != nil {
		return err
	}
	return serverStream.Send(response)
}

func (h *_FooYARPCHandler) EchoIn(serverStream *protobuf.ServerStream) error {
	requestMessage, err := serverStream.Receive(newFooServiceEchoInYARPCRequest)
	if requestMessage == nil {
		return err
	}

	request, ok := requestMessage.(*EchoInRequest)
	if !ok {
		return protobuf.CastError(emptyFooServiceEchoInYARPCRequest, requestMessage)
	}
	return h.server.EchoIn(request, &_FooServiceEchoInYARPCServer{serverStream: serverStream})
}

func (h *_FooYARPCHandler) EchoBoth(serverStream *protobuf.ServerStream) error {
	return h.server.EchoBoth(&_FooServiceEchoBothYARPCServer{serverStream: serverStream})
}

type _FooServiceEchoOutYARPCClient struct {
	stream *protobuf.ClientStream
}

func (c *_FooServiceEchoOutYARPCClient) Context() context.Context {
	return c.stream.Context()
}

func (c *_FooServiceEchoOutYARPCClient) Send(request *EchoOutRequest, options ...yarpc.StreamOption) error {
	return c.stream.Send(request, options...)
}

func (c *_FooServiceEchoOutYARPCClient) CloseAndRecv(options ...yarpc.StreamOption) (*EchoOutResponse, error) {
	if err := c.stream.Close(options...); err != nil {
		return nil, err
	}
	responseMessage, err := c.stream.Receive(newFooServiceEchoOutYARPCResponse, options...)
	if responseMessage == nil {
		return nil, err
	}
	response, ok := responseMessage.(*EchoOutResponse)
	if !ok {
		return nil, protobuf.CastError(emptyFooServiceEchoOutYARPCResponse, responseMessage)
	}
	return response, err
}

type _FooServiceEchoInYARPCClient struct {
	stream *protobuf.ClientStream
}

func (c *_FooServiceEchoInYARPCClient) Context() context.Context {
	return c.stream.Context()
}

func (c *_FooServiceEchoInYARPCClient) Recv(options ...yarpc.StreamOption) (*EchoInResponse, error) {
	responseMessage, err := c.stream.Receive(newFooServiceEchoInYARPCResponse, options...)
	if responseMessage == nil {
		return nil, err
	}
	response, ok := responseMessage.(*EchoInResponse)
	if !ok {
		return nil, protobuf.CastError(emptyFooServiceEchoInYARPCResponse, responseMessage)
	}
	return response, err
}

func (c *_FooServiceEchoInYARPCClient) CloseSend(options ...yarpc.StreamOption) error {
	return c.stream.Close(options...)
}

type _FooServiceEchoBothYARPCClient struct {
	stream *protobuf.ClientStream
}

func (c *_FooServiceEchoBothYARPCClient) Context() context.Context {
	return c.stream.Context()
}

func (c *_FooServiceEchoBothYARPCClient) Send(request *EchoBothRequest, options ...yarpc.StreamOption) error {
	return c.stream.Send(request, options...)
}

func (c *_FooServiceEchoBothYARPCClient) Recv(options ...yarpc.StreamOption) (*EchoBothResponse, error) {
	responseMessage, err := c.stream.Receive(newFooServiceEchoBothYARPCResponse, options...)
	if responseMessage == nil {
		return nil, err
	}
	response, ok := responseMessage.(*EchoBothResponse)
	if !ok {
		return nil, protobuf.CastError(emptyFooServiceEchoBothYARPCResponse, responseMessage)
	}
	return response, err
}

func (c *_FooServiceEchoBothYARPCClient) CloseSend(options ...yarpc.StreamOption) error {
	return c.stream.Close(options...)
}

type _FooServiceEchoOutYARPCServer struct {
	serverStream *protobuf.ServerStream
}

func (s *_FooServiceEchoOutYARPCServer) Context() context.Context {
	return s.serverStream.Context()
}

func (s *_FooServiceEchoOutYARPCServer) Recv(options ...yarpc.StreamOption) (*EchoOutRequest, error) {
	requestMessage, err := s.serverStream.Receive(newFooServiceEchoOutYARPCRequest, options...)
	if requestMessage == nil {
		return nil, err
	}
	request, ok := requestMessage.(*EchoOutRequest)
	if !ok {
		return nil, protobuf.CastError(emptyFooServiceEchoOutYARPCRequest, requestMessage)
	}
	return request, err
}

type _FooServiceEchoInYARPCServer struct {
	serverStream *protobuf.ServerStream
}

func (s *_FooServiceEchoInYARPCServer) Context() context.Context {
	return s.serverStream.Context()
}

func (s *_FooServiceEchoInYARPCServer) Send(response *EchoInResponse, options ...yarpc.StreamOption) error {
	return s.serverStream.Send(response, options...)
}

type _FooServiceEchoBothYARPCServer struct {
	serverStream *protobuf.ServerStream
}

func (s *_FooServiceEchoBothYARPCServer) Context() context.Context {
	return s.serverStream.Context()
}

func (s *_FooServiceEchoBothYARPCServer) Recv(options ...yarpc.StreamOption) (*EchoBothRequest, error) {
	requestMessage, err := s.serverStream.Receive(newFooServiceEchoBothYARPCRequest, options...)
	if requestMessage == nil {
		return nil, err
	}
	request, ok := requestMessage.(*EchoBothRequest)
	if !ok {
		return nil, protobuf.CastError(emptyFooServiceEchoBothYARPCRequest, requestMessage)
	}
	return request, err
}

func (s *_FooServiceEchoBothYARPCServer) Send(response *EchoBothResponse, options ...yarpc.StreamOption) error {
	return s.serverStream.Send(response, options...)
}

func newFooServiceEchoOutYARPCRequest() proto.Message {
	return &EchoOutRequest{}
}

func newFooServiceEchoOutYARPCResponse() proto.Message {
	return &EchoOutResponse{}
}

func newFooServiceEchoInYARPCRequest() proto.Message {
	return &EchoInRequest{}
}

func newFooServiceEchoInYARPCResponse() proto.Message {
	return &EchoInResponse{}
}

func newFooServiceEchoBothYARPCRequest() proto.Message {
	return &EchoBothRequest{}
}

func newFooServiceEchoBothYARPCResponse() proto.Message {
	return &EchoBothResponse{}
}

var (
	emptyFooServiceEchoOutYARPCRequest   = &EchoOutRequest{}
	emptyFooServiceEchoOutYARPCResponse  = &EchoOutResponse{}
	emptyFooServiceEchoInYARPCRequest    = &EchoInRequest{}
	emptyFooServiceEchoInYARPCResponse   = &EchoInResponse{}
	emptyFooServiceEchoBothYARPCRequest  = &EchoBothRequest{}
	emptyFooServiceEchoBothYARPCResponse = &EchoBothResponse{}
)

var yarpcFileDescriptorClosure15a1dc8d40dadaa6 = [][]byte{
	// example.proto
	[]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0xcf, 0x4a, 0xeb, 0x40,
		0x18, 0xc5, 0x3b, 0x0d, 0xb7, 0x7f, 0xbe, 0xb6, 0xb7, 0x61, 0xb8, 0x8b, 0xd0, 0x55, 0x6f, 0xba,
		0x09, 0xa2, 0xa1, 0x54, 0x37, 0x82, 0xad, 0x50, 0x50, 0x11, 0xf1, 0x0f, 0x09, 0xb8, 0x70, 0x53,
		0x92, 0x32, 0x5a, 0x31, 0xc9, 0xc4, 0x4c, 0x46, 0xec, 0x23, 0x08, 0xba, 0xf0, 0x35, 0x5c, 0xfb,
		0x80, 0x92, 0x64, 0xa6, 0xd5, 0x4a, 0x29, 0x8d, 0xee, 0x3a, 0xc3, 0x39, 0xe7, 0xfb, 0x75, 0xbe,
		0x13, 0x68, 0x90, 0x47, 0xc7, 0x0f, 0x3d, 0x62, 0x86, 0x11, 0x8d, 0x29, 0xde, 0xe2, 0x2e, 0x89,
		0xcc, 0xa9, 0x13, 0x85, 0x63, 0xf3, 0x36, 0x88, 0x49, 0x14, 0x38, 0x9e, 0x29, 0x24, 0x2c, 0xd3,
		0xb8, 0xfc, 0x5a, 0xde, 0xe8, 0x1d, 0x68, 0x1e, 0x91, 0xf8, 0xd2, 0xf1, 0x38, 0xb1, 0xc8, 0x3d,
		0x27, 0x2c, 0xc6, 0x2a, 0x28, 0x77, 0x64, 0xaa, 0xa1, 0x36, 0x32, 0xaa, 0x56, 0xf2, 0x53, 0x37,
		0x40, 0x9d, 0x8b, 0x58, 0x48, 0x03, 0x46, 0xf0, 0x3f, 0xf8, 0xf3, 0x90, 0x5c, 0x68, 0xc5, 0x54,
		0x97, 0x1d, 0xf4, 0x5d, 0x68, 0xda, 0xab, 0xe2, 0x96, 0x58, 0x31, 0xa8, 0xf6, 0xc2, 0x10, 0x7d,
		0x03, 0xfe, 0x1e, 0x8c, 0x27, 0xf4, 0x9c, 0xc7, 0x32, 0x4d, 0x83, 0xb2, 0x4f, 0x18, 0x73, 0x6e,
		0x88, 0x48, 0x94, 0x47, 0x7d, 0x07, 0x9a, 0x33, 0xad, 0x60, 0xfc, 0x0f, 0x75, 0xc7, 0xf3, 0x46,
		0x42, 0xc1, 0xb4, 0x62, 0x5b, 0x31, 0xaa, 0x56, 0xcd, 0xf1, 0xbc, 0x53, 0x71, 0xa5, 0x9f, 0x41,
		0x23, 0x71, 0x1d, 0x07, 0x2b, 0x07, 0xe0, 0x0e, 0x34, 0x02, 0xee, 0x8f, 0x22, 0x91, 0xce, 0x52,
		0x7c, 0xc5, 0xaa, 0x07, 0xdc, 0x97, 0x13, 0x99, 0x24, 0x4e, 0xf2, 0x04, 0xc4, 0x72, 0xe2, 0x8b,
		0x8c, 0x78, 0x48, 0xe3, 0xc9, 0x2f, 0x4d, 0xdf, 0x04, 0x75, 0x9e, 0xb8, 0x6a, 0x7e, 0xef, 0xad,
		0x08, 0x95, 0x13, 0x32, 0x4d, 0x9f, 0x1c, 0xbf, 0x20, 0xa8, 0xc8, 0x25, 0xe3, 0x81, 0xb9, 0x56,
		0x8b, 0xcc, 0x85, 0x0a, 0xb5, 0xf6, 0x73, 0xfb, 0xc5, 0xe2, 0x0b, 0x29, 0x8f, 0x9d, 0x97, 0xc7,
		0xfe, 0x21, 0xcf, 0xb7, 0x22, 0x16, 0x7a, 0xef, 0x0a, 0x28, 0x87, 0x94, 0xe2, 0x67, 0x04, 0x65,
		0xd1, 0x33, 0xdc, 0x5f, 0x33, 0xf6, 0x6b, 0x97, 0x5b, 0x83, 0xbc, 0x76, 0x09, 0x65, 0x20, 0xfc,
		0x84, 0xa0, 0x94, 0x15, 0x0e, 0xef, 0xe5, 0x88, 0x9b, 0xf5, 0xbe, 0xd5, 0xcf, 0xe9, 0x96, 0x2c,
		0x5d, 0x84, 0x5f, 0x11, 0x54, 0x64, 0xfd, 0x70, 0x9e, 0x3f, 0xf7, 0xe9, 0x4b, 0x58, 0x7b, 0x65,
		0x8b, 0xbd, 0x4f, 0x5e, 0xa7, 0x8b, 0x86, 0xb5, 0xab, 0xaa, 0x50, 0x84, 0xae, 0x5b, 0x4a, 0x5d,
		0xdb, 0x1f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xbb, 0x66, 0x8c, 0x91, 0x33, 0x05, 0x00, 0x00,
	},
}

func init() {
	yarpc.RegisterClientBuilder(
		func(clientConfig transport.ClientConfig, structField reflect.StructField) KeyValueYARPCClient {
			return NewKeyValueYARPCClient(clientConfig, protobuf.ClientBuilderOptions(clientConfig, structField)...)
		},
	)
	yarpc.RegisterClientBuilder(
		func(clientConfig transport.ClientConfig, structField reflect.StructField) FooYARPCClient {
			return NewFooYARPCClient(clientConfig, protobuf.ClientBuilderOptions(clientConfig, structField)...)
		},
	)
}

func init() {
	proto.RegisterType((*GetValueRequest)(nil), "uber.yarpc.internal.examples.protobuf.example.GetValueRequest")
	proto.RegisterType((*GetValueResponse)(nil), "uber.yarpc.internal.examples.protobuf.example.GetValueResponse")
	proto.RegisterType((*SetValueRequest)(nil), "uber.yarpc.internal.examples.protobuf.example.SetValueRequest")
	proto.RegisterType((*SetValueResponse)(nil), "uber.yarpc.internal.examples.protobuf.example.SetValueResponse")
	proto.RegisterType((*EchoOutRequest)(nil), "uber.yarpc.internal.examples.protobuf.example.EchoOutRequest")
	proto.RegisterType((*EchoOutResponse)(nil), "uber.yarpc.internal.examples.protobuf.example.EchoOutResponse")
	proto.RegisterType((*EchoInRequest)(nil), "uber.yarpc.internal.examples.protobuf.example.EchoInRequest")
	proto.RegisterType((*EchoInResponse)(nil), "uber.yarpc.internal.examples.protobuf.example.EchoInResponse")
	proto.RegisterType((*EchoBothRequest)(nil), "uber.yarpc.internal.examples.protobuf.example.EchoBothRequest")
	proto.RegisterType((*EchoBothResponse)(nil), "uber.yarpc.internal.examples.protobuf.example.EchoBothResponse")
}
