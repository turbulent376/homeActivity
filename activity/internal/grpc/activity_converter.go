package grpc

import (
	pb "git.jetbrains.space/orbi/fcsd/proto/timesheet"
	"git.jetbrains.space/orbi/fcsd/timesheet/internal/domain"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *Server) toTimesheetPb(rq *domain.Timesheet) *pb.Timesheet {
	return &pb.Timesheet{
		Id:       rq.Id,
		Owner:    rq.Owner,
		DateFrom: timestamppb.New(rq.DateFrom),
		DateTo:   timestamppb.New(rq.DateTo),
	}
}

func (s *Server) toCreateTimesheetDomain(rq *pb.CreateTimesheetRequest) *domain.Timesheet {
	return &domain.Timesheet{
		Owner:    rq.Owner,
		DateFrom: rq.DateFrom.AsTime(),
		DateTo:   rq.DateTo.AsTime(),
	}
}

func (s *Server) toUpdateTimesheetDomain(rq *pb.UpdateTimesheetRequest) *domain.Timesheet {
	return &domain.Timesheet{
		Id:       rq.Id,
		Owner:    rq.Owner,
		DateFrom: rq.DateFrom.AsTime(),
		DateTo:   rq.DateTo.AsTime(),
	}
}

func (s *Server) toSearchTimesheetDomain(rq *pb.SearchTimesheetRequest) *domain.SearchTimesheetRequest {
	dateFromSearch := rq.DateFromSearch.AsTime()
	dateToSearch := rq.DateToSearch.AsTime()

	return &domain.SearchTimesheetRequest{
		Owner:          rq.Owner,
		DateFromSearch: dateFromSearch,
		DateToSearch:   dateToSearch,
	}
}

func (s *Server) toCreateEventDomain(rq *pb.CreateEventRequest) *domain.Event {
	return &domain.Event{
		TimesheetId: rq.TimesheetId,
		Subject:     rq.Subject,
		WeekDay:     rq.WeekDay,
		TimeStart:   rq.TimeStart.AsTime(),
		TimeEnd:     rq.TimeEnd.AsTime(),
	}
}

func (s *Server) toUpdateEventDomain(rq *pb.UpdateEventRequest) *domain.Event {
	return &domain.Event{
		Id:          rq.Id,
		TimesheetId: rq.TimesheetId,
		Subject:     rq.Subject,
		WeekDay:     rq.WeekDay,
		TimeStart:   rq.TimeStart.AsTime(),
		TimeEnd:     rq.TimeEnd.AsTime(),
	}
}

func (s *Server) toEventPb(rq *domain.Event) *pb.Event {
	return &pb.Event{
		Id:          rq.Id,
		TimesheetId: rq.TimesheetId,
		Subject:     rq.Subject,
		WeekDay:     rq.WeekDay,
		TimeStart:   timestamppb.New(rq.TimeStart),
		TimeEnd:     timestamppb.New(rq.TimeEnd),
	}
}
