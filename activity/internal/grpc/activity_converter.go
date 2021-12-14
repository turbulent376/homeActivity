package grpc

import (
	pb "github.com/turbulent376/proto/activity"
	"github.com/turbulent376/homeactivity/activity/internal/domain"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *Server) toActivityPb(rq *domain.Activity) *pb.Timesheet {
	return &pb.Timesheet{
		Id:       rq.Id,
		Owner:    rq.Owner,
		DateFrom: timestamppb.New(rq.DateFrom),
		DateTo:   timestamppb.New(rq.DateTo),
	}
}

func (s *Server) toCreateActivityDomain(rq *pb.CreateTimesheetRequest) *domain.Activity {
	return &domain.Activity{
		Owner:    rq.Owner,
		DateFrom: rq.DateFrom.AsTime(),
		DateTo:   rq.DateTo.AsTime(),
	}
}

func (s *Server) toUpdateTimesheetDomain(rq *pb.UpdateTimesheetRequest) *domain.Activity {
	return &domain.Activity{
		Id:       rq.Id,
		Owner:    rq.Owner,
		DateFrom: rq.DateFrom.AsTime(),
		DateTo:   rq.DateTo.AsTime(),
	}
}


func (s *Server) toCreateActivityTypeDomain(rq *pb.CreateEventRequest) *domain.ActivityType {
	return &domain.ActivityType{

	}
}

func (s *Server) toUpdateActivityTypeDomain(rq *pb.UpdateEventRequest) *domain.ActivityType {
	return &domain.ActivityType{
		Id:          rq.Id,
	}
}

func (s *Server) toActivityTypePb(rq *domain.ActivityType) *pb.ActycityType {
	return &pb.ActycityType{
		Id:          rq.Id,
	}
}
