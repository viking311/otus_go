package internalgrpc

import (
	"context"
	"errors"
	"time"

	pb "github.com/viking311/otus_go/hw12_13_14_15_16_calendar/api"
	"github.com/viking311/otus_go/hw12_13_14_15_16_calendar/internal/app"
	"github.com/viking311/otus_go/hw12_13_14_15_16_calendar/internal/server"
	"github.com/viking311/otus_go/hw12_13_14_15_16_calendar/internal/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Service struct {
	pb.UnimplementedEventServiceServer
	app server.Application
}

func (s *Service) GetAllEvents(_ context.Context, _ *pb.GetAllEventsRequest) (*pb.EventList, error) {
	eventList := s.app.GetEvents()
	response := pb.EventList{
		Events: make([]*pb.Event, 0, len(eventList)),
	}

	for _, event := range eventList {
		response.Events = append(response.Events, s.storageEventToProtoEvent(event))
	}

	return &response, nil
}

func (s *Service) GetEventByID(_ context.Context, request *pb.GetEventByIdRequest) (*pb.Event, error) {
	event := s.app.GetEventByID(request.GetId())
	if event == nil {
		return nil, status.Error(codes.NotFound, "event not found")
	}

	return s.storageEventToProtoEvent(*event), nil
}

func (s *Service) GetEventsByUserID(_ context.Context, request *pb.GetEventsByUserIdRequest) (*pb.EventList, error) {
	eventList := s.app.GetEventsByUserID(request.GetUserId())

	response := pb.EventList{
		Events: make([]*pb.Event, 0, len(eventList)),
	}

	for _, event := range eventList {
		response.Events = append(response.Events, s.storageEventToProtoEvent(event))
	}

	return &response, nil
}

func (s *Service) GetEventsByUserIDAndDates(
	_ context.Context,
	request *pb.GetEventsByUserIdAndDatesRequest,
) (*pb.EventList, error) {
	startTime := time.Unix(request.GetStartTime().GetSeconds(), int64(request.GetStartTime().GetNanos()))
	endTime := time.Unix(request.GetEndTime().GetSeconds(), int64(request.GetEndTime().GetNanos()))

	eventList := s.app.GetEventsByUserIDAndDates(request.GetUserId(), startTime, endTime)

	response := pb.EventList{
		Events: make([]*pb.Event, 0, len(eventList)),
	}

	for _, event := range eventList {
		response.Events = append(response.Events, s.storageEventToProtoEvent(event))
	}

	return &response, nil
}

func (s *Service) SaveEvent(_ context.Context, request *pb.SaveEventRequest) (*pb.Event, error) {
	rawEvent, err := s.protoEventToStorageEvent(request.GetEvent())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	savedEvent, err := s.app.SaveEvent(*rawEvent)
	if err != nil {
		var verr *app.FieldValidationError
		if errors.As(err, &verr) {
			return nil, status.Error(codes.InvalidArgument, verr.Error())
		}
		return nil, status.Error(codes.Internal, "Unexpected error")
	}

	return s.storageEventToProtoEvent(*savedEvent), nil
}

func (s *Service) DeleteEvent(_ context.Context, request *pb.DeleteEventRequest) (*pb.DeleteEventResponse, error) {
	s.app.DeleteEvent(request.GetId())

	return &pb.DeleteEventResponse{
		Status: "ok",
	}, nil
}

func (s *Service) protoEventToStorageEvent(pbEvent *pb.Event) (*storage.Event, error) {
	dr, err := time.ParseDuration(pbEvent.Duration)
	if err != nil {
		return nil, err
	}

	return &storage.Event{
		ID:          pbEvent.Id,
		Title:       pbEvent.Title,
		UserID:      pbEvent.UserId,
		DateTime:    time.Unix(pbEvent.DateTime.Seconds, int64(pbEvent.DateTime.Nanos)),
		Description: pbEvent.Description,
		Duration:    dr,
		RemindTime:  pbEvent.RemindTime,
	}, nil
}

func (s *Service) storageEventToProtoEvent(event storage.Event) *pb.Event {
	return &pb.Event{
		Id:          event.ID,
		Title:       event.Title,
		UserId:      event.UserID,
		DateTime:    timestamppb.New(event.DateTime),
		Description: event.Description,
		Duration:    event.Duration.String(),
		RemindTime:  event.RemindTime,
	}
}

func NewService(app server.Application) *Service {
	return &Service{
		app: app,
	}
}
