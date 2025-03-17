// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: api/EventService.proto

package pb

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
	EventService_GetAllEvents_FullMethodName              = "/event.EventService/GetAllEvents"
	EventService_GetEventById_FullMethodName              = "/event.EventService/GetEventById"
	EventService_GetEventsByUserId_FullMethodName         = "/event.EventService/GetEventsByUserId"
	EventService_GetEventsByUserIdAndDates_FullMethodName = "/event.EventService/GetEventsByUserIdAndDates"
	EventService_SaveEvent_FullMethodName                 = "/event.EventService/SaveEvent"
	EventService_DeleteEvent_FullMethodName               = "/event.EventService/DeleteEvent"
)

// EventServiceClient is the client API for EventService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EventServiceClient interface {
	GetAllEvents(ctx context.Context, in *GetAllEventsRequest, opts ...grpc.CallOption) (*EventList, error)
	GetEventById(ctx context.Context, in *GetEventByIdRequest, opts ...grpc.CallOption) (*Event, error)
	GetEventsByUserId(ctx context.Context, in *GetEventsByUserIdRequest, opts ...grpc.CallOption) (*EventList, error)
	GetEventsByUserIdAndDates(ctx context.Context, in *GetEventsByUserIdAndDatesRequest, opts ...grpc.CallOption) (*EventList, error)
	SaveEvent(ctx context.Context, in *SaveEventRequest, opts ...grpc.CallOption) (*Event, error)
	DeleteEvent(ctx context.Context, in *DeleteEventRequest, opts ...grpc.CallOption) (*DeleteEventResponse, error)
}

type eventServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEventServiceClient(cc grpc.ClientConnInterface) EventServiceClient {
	return &eventServiceClient{cc}
}

func (c *eventServiceClient) GetAllEvents(ctx context.Context, in *GetAllEventsRequest, opts ...grpc.CallOption) (*EventList, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EventList)
	err := c.cc.Invoke(ctx, EventService_GetAllEvents_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventServiceClient) GetEventById(ctx context.Context, in *GetEventByIdRequest, opts ...grpc.CallOption) (*Event, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Event)
	err := c.cc.Invoke(ctx, EventService_GetEventById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventServiceClient) GetEventsByUserId(ctx context.Context, in *GetEventsByUserIdRequest, opts ...grpc.CallOption) (*EventList, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EventList)
	err := c.cc.Invoke(ctx, EventService_GetEventsByUserId_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventServiceClient) GetEventsByUserIdAndDates(ctx context.Context, in *GetEventsByUserIdAndDatesRequest, opts ...grpc.CallOption) (*EventList, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EventList)
	err := c.cc.Invoke(ctx, EventService_GetEventsByUserIdAndDates_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventServiceClient) SaveEvent(ctx context.Context, in *SaveEventRequest, opts ...grpc.CallOption) (*Event, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Event)
	err := c.cc.Invoke(ctx, EventService_SaveEvent_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventServiceClient) DeleteEvent(ctx context.Context, in *DeleteEventRequest, opts ...grpc.CallOption) (*DeleteEventResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteEventResponse)
	err := c.cc.Invoke(ctx, EventService_DeleteEvent_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EventServiceServer is the server API for EventService service.
// All implementations must embed UnimplementedEventServiceServer
// for forward compatibility.
type EventServiceServer interface {
	GetAllEvents(context.Context, *GetAllEventsRequest) (*EventList, error)
	GetEventById(context.Context, *GetEventByIdRequest) (*Event, error)
	GetEventsByUserId(context.Context, *GetEventsByUserIdRequest) (*EventList, error)
	GetEventsByUserIdAndDates(context.Context, *GetEventsByUserIdAndDatesRequest) (*EventList, error)
	SaveEvent(context.Context, *SaveEventRequest) (*Event, error)
	DeleteEvent(context.Context, *DeleteEventRequest) (*DeleteEventResponse, error)
	mustEmbedUnimplementedEventServiceServer()
}

// UnimplementedEventServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedEventServiceServer struct{}

func (UnimplementedEventServiceServer) GetAllEvents(context.Context, *GetAllEventsRequest) (*EventList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllEvents not implemented")
}
func (UnimplementedEventServiceServer) GetEventById(context.Context, *GetEventByIdRequest) (*Event, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEventById not implemented")
}
func (UnimplementedEventServiceServer) GetEventsByUserId(context.Context, *GetEventsByUserIdRequest) (*EventList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEventsByUserId not implemented")
}
func (UnimplementedEventServiceServer) GetEventsByUserIdAndDates(context.Context, *GetEventsByUserIdAndDatesRequest) (*EventList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEventsByUserIdAndDates not implemented")
}
func (UnimplementedEventServiceServer) SaveEvent(context.Context, *SaveEventRequest) (*Event, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveEvent not implemented")
}
func (UnimplementedEventServiceServer) DeleteEvent(context.Context, *DeleteEventRequest) (*DeleteEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEvent not implemented")
}
func (UnimplementedEventServiceServer) mustEmbedUnimplementedEventServiceServer() {}
func (UnimplementedEventServiceServer) testEmbeddedByValue()                      {}

// UnsafeEventServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EventServiceServer will
// result in compilation errors.
type UnsafeEventServiceServer interface {
	mustEmbedUnimplementedEventServiceServer()
}

func RegisterEventServiceServer(s grpc.ServiceRegistrar, srv EventServiceServer) {
	// If the following call pancis, it indicates UnimplementedEventServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&EventService_ServiceDesc, srv)
}

func _EventService_GetAllEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllEventsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).GetAllEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EventService_GetAllEvents_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).GetAllEvents(ctx, req.(*GetAllEventsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventService_GetEventById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEventByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).GetEventById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EventService_GetEventById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).GetEventById(ctx, req.(*GetEventByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventService_GetEventsByUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEventsByUserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).GetEventsByUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EventService_GetEventsByUserId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).GetEventsByUserId(ctx, req.(*GetEventsByUserIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventService_GetEventsByUserIdAndDates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEventsByUserIdAndDatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).GetEventsByUserIdAndDates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EventService_GetEventsByUserIdAndDates_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).GetEventsByUserIdAndDates(ctx, req.(*GetEventsByUserIdAndDatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventService_SaveEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).SaveEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EventService_SaveEvent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).SaveEvent(ctx, req.(*SaveEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventService_DeleteEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).DeleteEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EventService_DeleteEvent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).DeleteEvent(ctx, req.(*DeleteEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// EventService_ServiceDesc is the grpc.ServiceDesc for EventService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EventService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "event.EventService",
	HandlerType: (*EventServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllEvents",
			Handler:    _EventService_GetAllEvents_Handler,
		},
		{
			MethodName: "GetEventById",
			Handler:    _EventService_GetEventById_Handler,
		},
		{
			MethodName: "GetEventsByUserId",
			Handler:    _EventService_GetEventsByUserId_Handler,
		},
		{
			MethodName: "GetEventsByUserIdAndDates",
			Handler:    _EventService_GetEventsByUserIdAndDates_Handler,
		},
		{
			MethodName: "SaveEvent",
			Handler:    _EventService_SaveEvent_Handler,
		},
		{
			MethodName: "DeleteEvent",
			Handler:    _EventService_DeleteEvent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/EventService.proto",
}
