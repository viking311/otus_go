package internalgrpc

import (
	"context"
	"time"

	"google.golang.org/grpc/codes"

	"github.com/viking311/otus_go/hw12_13_14_15_16_calendar/internal/storage"

	status "google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/viking311/otus_go/hw12_13_14_15_16_calendar/api"
	"github.com/viking311/otus_go/hw12_13_14_15_16_calendar/internal/server"
)

type Service struct {
	pb.UnimplementedEventServiceServer
	app server.Application
}

func (s *Service) GetAllEvents(_ context.Context, request *pb.GetAllEventsRequest) (*pb.EventList, error) {
	eventList := s.app.GetEvents()
	response := pb.EventList{
		Events: make([]*pb.Event, 0, len(eventList)),
	}

	for _, event := range eventList {
		response.Events = append(response.Events, s.storageEventToProtoEvent(event))
	}

	return &response, nil
}

func (s *Service) GetEventById(_ context.Context, request *pb.GetEventByIdRequest) (*pb.Event, error) {
	event := s.app.GetEventById(request.GetId())
	if event == nil {
		return nil, status.Error(codes.NotFound, "event not found")
	}

	return s.storageEventToProtoEvent(*event), nil
}

func (s *Service) GetEventsByUserId(_ context.Context, request *pb.GetEventsByUserIdRequest) (*pb.EventList, error) {
	eventList := s.app.GetEventsByUserId(request.GetUserId())

	response := pb.EventList{
		Events: make([]*pb.Event, 0, len(eventList)),
	}

	for _, event := range eventList {
		response.Events = append(response.Events, s.storageEventToProtoEvent(event))
	}

	return &response, nil
}

func (s *Service) GetEventsByUserIdAndDates(_ context.Context, request *pb.GetEventsByUserIdAndDatesRequest) (*pb.EventList, error) {
	startTime := time.Unix(request.GetStartTime().GetSeconds(), int64(request.GetStartTime().GetNanos()))
	endTime := time.Unix(request.GetEndTime().GetSeconds(), int64(request.GetEndTime().GetNanos()))

	eventList := s.app.GetEventsByUserIdAndDates(request.GetUserId(), startTime, endTime)

	response := pb.EventList{
		Events: make([]*pb.Event, 0, len(eventList)),
	}

	for _, event := range eventList {
		response.Events = append(response.Events, s.storageEventToProtoEvent(event))
	}

	return &response, nil
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
