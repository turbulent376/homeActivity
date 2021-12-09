package grpc

import (
	"context"
	pb "git.jetbrains.space/orbi/fcsd/proto/timesheet"
)

func (s *Server) Create(ctx context.Context, rq *pb.CreateTimesheetRequest) (*pb.Timesheet, error) {
	timesheet, err := s.timesheetService.Create(ctx, s.toCreateTimesheetDomain(rq))
	if err != nil {
		return nil, err
	}
	return s.toTimesheetPb(timesheet), nil
}

func (s *Server) Update(ctx context.Context, rq *pb.UpdateTimesheetRequest) (*pb.Timesheet, error) {
	timesheet, err := s.timesheetService.Update(ctx, s.toUpdateTimesheetDomain(rq))
	if err != nil {
		return nil, err
	}
	return s.toTimesheetPb(timesheet), nil
}

func (s *Server) Get(ctx context.Context, rq *pb.TimesheetIdRequest) (*pb.Timesheet, error) {
	found, timesheet, err := s.timesheetService.Get(ctx, rq.Id)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, nil
	}
	return s.toTimesheetPb(timesheet), nil
}

func (s *Server) Search(ctx context.Context, rq *pb.SearchTimesheetRequest) (*pb.Timesheet, error) {
	found, timesheet, err := s.timesheetService.Search(ctx, s.toSearchTimesheetDomain(rq))

	if err != nil {
		return nil, err
	}
	if !found {
		return nil, nil
	}
	return s.toTimesheetPb(timesheet), nil
}

func (s *Server) Delete(ctx context.Context, rq *pb.TimesheetIdRequest) (*pb.EmptyResponse, error) {
	err := s.timesheetService.Delete(ctx, rq.Id)
	if err != nil {
		return nil, err
	}
	return &pb.EmptyResponse{}, nil
}

func (s *Server) CreateEvent(ctx context.Context, rq *pb.CreateEventRequest) (*pb.Event, error) {
	res, err := s.timesheetService.CreateEvent(ctx, s.toCreateEventDomain(rq))

	if err != nil {
		return nil, err
	}

	return s.toEventPb(res), nil
}

func (s *Server) UpdateEvent(ctx context.Context, rq *pb.UpdateEventRequest) (*pb.Event, error) {
	res, err := s.timesheetService.UpdateEvent(ctx, s.toUpdateEventDomain(rq))

	if err != nil {
		return nil, err
	}

	return s.toEventPb(res), nil
}

func (s *Server) GetEvent(ctx context.Context, rq *pb.EventIdRequest) (*pb.Event, error) {
	found, res, err := s.timesheetService.GetEvent(ctx, rq.Id)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, nil
	}
	return s.toEventPb(res), nil
}

func (s *Server) DeleteEvent(ctx context.Context, rq *pb.EventIdRequest) (*pb.EmptyResponse, error) {
	err := s.timesheetService.DeleteEvent(ctx, rq.Id)

	if err != nil {
		return nil, err
	}

	return &pb.EmptyResponse{}, nil
}
