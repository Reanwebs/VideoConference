// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: conference.proto

package pb

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

const (
	Conference_HealthCheck_FullMethodName             = "/pb.Conference/HealthCheck"
	Conference_StartConference_FullMethodName         = "/pb.Conference/StartConference"
	Conference_JoinConference_FullMethodName          = "/pb.Conference/JoinConference"
	Conference_AcceptJoining_FullMethodName           = "/pb.Conference/AcceptJoining"
	Conference_DeclineJoining_FullMethodName          = "/pb.Conference/DeclineJoining"
	Conference_RemoveParticipant_FullMethodName       = "/pb.Conference/RemoveParticipant"
	Conference_ToggleCamera_FullMethodName            = "/pb.Conference/ToggleCamera"
	Conference_ToggleMic_FullMethodName               = "/pb.Conference/ToggleMic"
	Conference_ToggleParticipantCamera_FullMethodName = "/pb.Conference/ToggleParticipantCamera"
	Conference_ToggleParticipantMic_FullMethodName    = "/pb.Conference/ToggleParticipantMic"
	Conference_EndConference_FullMethodName           = "/pb.Conference/EndConference"
)

// ConferenceClient is the client API for Conference service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConferenceClient interface {
	HealthCheck(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	StartConference(ctx context.Context, in *StartConferenceRequest, opts ...grpc.CallOption) (*StartConferenceResponse, error)
	JoinConference(ctx context.Context, in *JoinConferenceRequest, opts ...grpc.CallOption) (*JoinConferenceResponse, error)
	AcceptJoining(ctx context.Context, in *AcceptJoiningRequest, opts ...grpc.CallOption) (*AcceptJoiningResponse, error)
	DeclineJoining(ctx context.Context, in *DeclineJoiningRequest, opts ...grpc.CallOption) (*DeclineJoiningResponse, error)
	RemoveParticipant(ctx context.Context, in *RemoveParticipantRequest, opts ...grpc.CallOption) (*RemoveParticipantResponse, error)
	ToggleCamera(ctx context.Context, in *ToggleCameraRequest, opts ...grpc.CallOption) (*ToggleCameraResponse, error)
	ToggleMic(ctx context.Context, in *ToggleMicRequest, opts ...grpc.CallOption) (*ToggleMicResponse, error)
	ToggleParticipantCamera(ctx context.Context, in *ToggleParticipantCameraRequest, opts ...grpc.CallOption) (*ToggleParticipantCameraResponse, error)
	ToggleParticipantMic(ctx context.Context, in *ToggleParticipantMicRequest, opts ...grpc.CallOption) (*ToggleParticipantMicResponse, error)
	EndConference(ctx context.Context, in *EndConferenceRequest, opts ...grpc.CallOption) (*EndConferenceResponse, error)
}

type conferenceClient struct {
	cc grpc.ClientConnInterface
}

func NewConferenceClient(cc grpc.ClientConnInterface) ConferenceClient {
	return &conferenceClient{cc}
}

func (c *conferenceClient) HealthCheck(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, Conference_HealthCheck_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *conferenceClient) StartConference(ctx context.Context, in *StartConferenceRequest, opts ...grpc.CallOption) (*StartConferenceResponse, error) {
	out := new(StartConferenceResponse)
	err := c.cc.Invoke(ctx, Conference_StartConference_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *conferenceClient) JoinConference(ctx context.Context, in *JoinConferenceRequest, opts ...grpc.CallOption) (*JoinConferenceResponse, error) {
	out := new(JoinConferenceResponse)
	err := c.cc.Invoke(ctx, Conference_JoinConference_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *conferenceClient) AcceptJoining(ctx context.Context, in *AcceptJoiningRequest, opts ...grpc.CallOption) (*AcceptJoiningResponse, error) {
	out := new(AcceptJoiningResponse)
	err := c.cc.Invoke(ctx, Conference_AcceptJoining_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *conferenceClient) DeclineJoining(ctx context.Context, in *DeclineJoiningRequest, opts ...grpc.CallOption) (*DeclineJoiningResponse, error) {
	out := new(DeclineJoiningResponse)
	err := c.cc.Invoke(ctx, Conference_DeclineJoining_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *conferenceClient) RemoveParticipant(ctx context.Context, in *RemoveParticipantRequest, opts ...grpc.CallOption) (*RemoveParticipantResponse, error) {
	out := new(RemoveParticipantResponse)
	err := c.cc.Invoke(ctx, Conference_RemoveParticipant_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *conferenceClient) ToggleCamera(ctx context.Context, in *ToggleCameraRequest, opts ...grpc.CallOption) (*ToggleCameraResponse, error) {
	out := new(ToggleCameraResponse)
	err := c.cc.Invoke(ctx, Conference_ToggleCamera_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *conferenceClient) ToggleMic(ctx context.Context, in *ToggleMicRequest, opts ...grpc.CallOption) (*ToggleMicResponse, error) {
	out := new(ToggleMicResponse)
	err := c.cc.Invoke(ctx, Conference_ToggleMic_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *conferenceClient) ToggleParticipantCamera(ctx context.Context, in *ToggleParticipantCameraRequest, opts ...grpc.CallOption) (*ToggleParticipantCameraResponse, error) {
	out := new(ToggleParticipantCameraResponse)
	err := c.cc.Invoke(ctx, Conference_ToggleParticipantCamera_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *conferenceClient) ToggleParticipantMic(ctx context.Context, in *ToggleParticipantMicRequest, opts ...grpc.CallOption) (*ToggleParticipantMicResponse, error) {
	out := new(ToggleParticipantMicResponse)
	err := c.cc.Invoke(ctx, Conference_ToggleParticipantMic_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *conferenceClient) EndConference(ctx context.Context, in *EndConferenceRequest, opts ...grpc.CallOption) (*EndConferenceResponse, error) {
	out := new(EndConferenceResponse)
	err := c.cc.Invoke(ctx, Conference_EndConference_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConferenceServer is the server API for Conference service.
// All implementations must embed UnimplementedConferenceServer
// for forward compatibility
type ConferenceServer interface {
	HealthCheck(context.Context, *Request) (*Response, error)
	StartConference(context.Context, *StartConferenceRequest) (*StartConferenceResponse, error)
	JoinConference(context.Context, *JoinConferenceRequest) (*JoinConferenceResponse, error)
	AcceptJoining(context.Context, *AcceptJoiningRequest) (*AcceptJoiningResponse, error)
	DeclineJoining(context.Context, *DeclineJoiningRequest) (*DeclineJoiningResponse, error)
	RemoveParticipant(context.Context, *RemoveParticipantRequest) (*RemoveParticipantResponse, error)
	ToggleCamera(context.Context, *ToggleCameraRequest) (*ToggleCameraResponse, error)
	ToggleMic(context.Context, *ToggleMicRequest) (*ToggleMicResponse, error)
	ToggleParticipantCamera(context.Context, *ToggleParticipantCameraRequest) (*ToggleParticipantCameraResponse, error)
	ToggleParticipantMic(context.Context, *ToggleParticipantMicRequest) (*ToggleParticipantMicResponse, error)
	EndConference(context.Context, *EndConferenceRequest) (*EndConferenceResponse, error)
	mustEmbedUnimplementedConferenceServer()
}

// UnimplementedConferenceServer must be embedded to have forward compatible implementations.
type UnimplementedConferenceServer struct {
}

func (UnimplementedConferenceServer) HealthCheck(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HealthCheck not implemented")
}
func (UnimplementedConferenceServer) StartConference(context.Context, *StartConferenceRequest) (*StartConferenceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartConference not implemented")
}
func (UnimplementedConferenceServer) JoinConference(context.Context, *JoinConferenceRequest) (*JoinConferenceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinConference not implemented")
}
func (UnimplementedConferenceServer) AcceptJoining(context.Context, *AcceptJoiningRequest) (*AcceptJoiningResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AcceptJoining not implemented")
}
func (UnimplementedConferenceServer) DeclineJoining(context.Context, *DeclineJoiningRequest) (*DeclineJoiningResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeclineJoining not implemented")
}
func (UnimplementedConferenceServer) RemoveParticipant(context.Context, *RemoveParticipantRequest) (*RemoveParticipantResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveParticipant not implemented")
}
func (UnimplementedConferenceServer) ToggleCamera(context.Context, *ToggleCameraRequest) (*ToggleCameraResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ToggleCamera not implemented")
}
func (UnimplementedConferenceServer) ToggleMic(context.Context, *ToggleMicRequest) (*ToggleMicResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ToggleMic not implemented")
}
func (UnimplementedConferenceServer) ToggleParticipantCamera(context.Context, *ToggleParticipantCameraRequest) (*ToggleParticipantCameraResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ToggleParticipantCamera not implemented")
}
func (UnimplementedConferenceServer) ToggleParticipantMic(context.Context, *ToggleParticipantMicRequest) (*ToggleParticipantMicResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ToggleParticipantMic not implemented")
}
func (UnimplementedConferenceServer) EndConference(context.Context, *EndConferenceRequest) (*EndConferenceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EndConference not implemented")
}
func (UnimplementedConferenceServer) mustEmbedUnimplementedConferenceServer() {}

// UnsafeConferenceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConferenceServer will
// result in compilation errors.
type UnsafeConferenceServer interface {
	mustEmbedUnimplementedConferenceServer()
}

func RegisterConferenceServer(s grpc.ServiceRegistrar, srv ConferenceServer) {
	s.RegisterService(&Conference_ServiceDesc, srv)
}

func _Conference_HealthCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConferenceServer).HealthCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Conference_HealthCheck_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConferenceServer).HealthCheck(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Conference_StartConference_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartConferenceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConferenceServer).StartConference(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Conference_StartConference_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConferenceServer).StartConference(ctx, req.(*StartConferenceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Conference_JoinConference_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinConferenceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConferenceServer).JoinConference(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Conference_JoinConference_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConferenceServer).JoinConference(ctx, req.(*JoinConferenceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Conference_AcceptJoining_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AcceptJoiningRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConferenceServer).AcceptJoining(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Conference_AcceptJoining_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConferenceServer).AcceptJoining(ctx, req.(*AcceptJoiningRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Conference_DeclineJoining_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeclineJoiningRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConferenceServer).DeclineJoining(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Conference_DeclineJoining_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConferenceServer).DeclineJoining(ctx, req.(*DeclineJoiningRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Conference_RemoveParticipant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveParticipantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConferenceServer).RemoveParticipant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Conference_RemoveParticipant_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConferenceServer).RemoveParticipant(ctx, req.(*RemoveParticipantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Conference_ToggleCamera_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ToggleCameraRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConferenceServer).ToggleCamera(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Conference_ToggleCamera_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConferenceServer).ToggleCamera(ctx, req.(*ToggleCameraRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Conference_ToggleMic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ToggleMicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConferenceServer).ToggleMic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Conference_ToggleMic_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConferenceServer).ToggleMic(ctx, req.(*ToggleMicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Conference_ToggleParticipantCamera_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ToggleParticipantCameraRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConferenceServer).ToggleParticipantCamera(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Conference_ToggleParticipantCamera_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConferenceServer).ToggleParticipantCamera(ctx, req.(*ToggleParticipantCameraRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Conference_ToggleParticipantMic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ToggleParticipantMicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConferenceServer).ToggleParticipantMic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Conference_ToggleParticipantMic_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConferenceServer).ToggleParticipantMic(ctx, req.(*ToggleParticipantMicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Conference_EndConference_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EndConferenceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConferenceServer).EndConference(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Conference_EndConference_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConferenceServer).EndConference(ctx, req.(*EndConferenceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Conference_ServiceDesc is the grpc.ServiceDesc for Conference service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Conference_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Conference",
	HandlerType: (*ConferenceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HealthCheck",
			Handler:    _Conference_HealthCheck_Handler,
		},
		{
			MethodName: "StartConference",
			Handler:    _Conference_StartConference_Handler,
		},
		{
			MethodName: "JoinConference",
			Handler:    _Conference_JoinConference_Handler,
		},
		{
			MethodName: "AcceptJoining",
			Handler:    _Conference_AcceptJoining_Handler,
		},
		{
			MethodName: "DeclineJoining",
			Handler:    _Conference_DeclineJoining_Handler,
		},
		{
			MethodName: "RemoveParticipant",
			Handler:    _Conference_RemoveParticipant_Handler,
		},
		{
			MethodName: "ToggleCamera",
			Handler:    _Conference_ToggleCamera_Handler,
		},
		{
			MethodName: "ToggleMic",
			Handler:    _Conference_ToggleMic_Handler,
		},
		{
			MethodName: "ToggleParticipantCamera",
			Handler:    _Conference_ToggleParticipantCamera_Handler,
		},
		{
			MethodName: "ToggleParticipantMic",
			Handler:    _Conference_ToggleParticipantMic_Handler,
		},
		{
			MethodName: "EndConference",
			Handler:    _Conference_EndConference_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "conference.proto",
}
