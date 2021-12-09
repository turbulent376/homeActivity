package timesheet

import (
	pb "git.jetbrains.space/orbi/fcsd/proto/timesheet"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (c *ctrlImpl) toCreateTimesheetPb(rq *CreateTimesheetRequest) *pb.CreateTimesheetRequest {
	return &pb.CreateTimesheetRequest{
		Owner:    rq.Owner,
		DateFrom: timestamppb.New(rq.DateFrom),
		DateTo:   timestamppb.New(rq.DateTo),
	}
}

func (c *ctrlImpl) toTimesheetApi(rq *pb.Timesheet) *Timesheet {
	return &Timesheet{
		Id:       rq.Id,
		Owner:    rq.Owner,
		DateFrom: rq.DateFrom.AsTime(),
		DateTo:   rq.DateTo.AsTime(),
	}
}

func (c *ctrlImpl) toUpdateTimesheetPb(rq *UpdateTimesheetRequest) *pb.UpdateTimesheetRequest {
	return &pb.UpdateTimesheetRequest{
		Id:      rq.Id,
		DateFrom: timestamppb.New(rq.DateFrom),
		DateTo: timestamppb.New(rq.DateTo),
	}
}

func (c *ctrlImpl) toCreateEventPb(rq *CreateEventRequest) *pb.CreateEventRequest {
	return &pb.CreateEventRequest{
		TimesheetId:  rq.TimesheetId,
		Subject:      rq.Subject,
		WeekDay:      rq.WeekDay,
		TimeStart:    timestamppb.New(rq.TimeStart),
		TimeEnd:      timestamppb.New(rq.TimeEnd),
	}
}

func (c *ctrlImpl) toEventApi(rq *pb.Event) *Event {
	return &Event{
		Id:           rq.Id,
		TimesheetId:  rq.TimesheetId,
		Subject:      rq.Subject,
		TimeStart:    rq.TimeStart.AsTime(),
		TimeEnd:      rq.TimeEnd.AsTime(),
	}
}

func (c *ctrlImpl) toUpdateEventPb(rq *UpdateEventRequest) *pb.UpdateEventRequest {
	return &pb.UpdateEventRequest{
		Id:          rq.Id,
		Subject:      rq.Subject,
		WeekDay:      rq.WeekDay,
		TimeStart:    timestamppb.New(rq.TimeStart),
		TimeEnd:      timestamppb.New(rq.TimeEnd),
	}
}

func (c *ctrlImpl) toEventsApi(rq *pb.SearchResponse) EventSearchResponse {
	var res []*Event
	for _, result := range rq.GetEvents() {
		res = append(res, c.toEventApi(result))
	}
	return EventSearchResponse{Result: res}
}
