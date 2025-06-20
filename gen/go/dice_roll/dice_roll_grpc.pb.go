// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.25.3
// source: dice_roll/dice_roll.proto

package proto_dice_roll

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	DiceRollGameAPI_Play_FullMethodName = "/dice_roll.DiceRollGameAPI/Play"
)

// DiceRollGameAPIClient is the client API for DiceRollGameAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DiceRollGameAPIClient interface {
	Play(ctx context.Context, in *PlayRequest, opts ...grpc.CallOption) (*PlayResponse, error)
}

type diceRollGameAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewDiceRollGameAPIClient(cc grpc.ClientConnInterface) DiceRollGameAPIClient {
	return &diceRollGameAPIClient{cc}
}

func (c *diceRollGameAPIClient) Play(ctx context.Context, in *PlayRequest, opts ...grpc.CallOption) (*PlayResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PlayResponse)
	err := c.cc.Invoke(ctx, DiceRollGameAPI_Play_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DiceRollGameAPIServer is the server API for DiceRollGameAPI service.
// All implementations must embed UnimplementedDiceRollGameAPIServer
// for forward compatibility.
type DiceRollGameAPIServer interface {
	Play(context.Context, *PlayRequest) (*PlayResponse, error)
	mustEmbedUnimplementedDiceRollGameAPIServer()
}

// UnimplementedDiceRollGameAPIServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedDiceRollGameAPIServer struct{}

func (UnimplementedDiceRollGameAPIServer) Play(context.Context, *PlayRequest) (*PlayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Play not implemented")
}
func (UnimplementedDiceRollGameAPIServer) mustEmbedUnimplementedDiceRollGameAPIServer() {}
func (UnimplementedDiceRollGameAPIServer) testEmbeddedByValue()                         {}

// UnsafeDiceRollGameAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DiceRollGameAPIServer will
// result in compilation errors.
type UnsafeDiceRollGameAPIServer interface {
	mustEmbedUnimplementedDiceRollGameAPIServer()
}

func RegisterDiceRollGameAPIServer(s grpc.ServiceRegistrar, srv DiceRollGameAPIServer) {
	// If the following call pancis, it indicates UnimplementedDiceRollGameAPIServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&DiceRollGameAPI_ServiceDesc, srv)
}

func _DiceRollGameAPI_Play_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiceRollGameAPIServer).Play(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DiceRollGameAPI_Play_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiceRollGameAPIServer).Play(ctx, req.(*PlayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DiceRollGameAPI_ServiceDesc is the grpc.ServiceDesc for DiceRollGameAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DiceRollGameAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dice_roll.DiceRollGameAPI",
	HandlerType: (*DiceRollGameAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Play",
			Handler:    _DiceRollGameAPI_Play_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dice_roll/dice_roll.proto",
}
