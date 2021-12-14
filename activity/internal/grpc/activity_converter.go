package grpc

import (
	"github.com/turbulent376/homeactivity/activity/internal/domain"
	pb "github.com/turbulent376/proto/activity"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *Server) toActivityPb(rq *domain.Activity) *pb.Activity {
	return &pb.Activity{
		Id:       rq.Id,
		Owner:    rq.Owner,
		Family:   rq.Family,
		Type:     rq.Type,
		DateFrom: timestamppb.New(rq.DateFrom),
		DateTo:   timestamppb.New(rq.DateTo),
	}
}

func (s *Server) toCreateActivityDomain(rq *pb.CreateActivityRequest) *domain.Activity {
	return &domain.Activity{
		Owner:    rq.Owner,
		Family:   rq.Family,
		Type:     rq.Family,
		DateFrom: rq.DateFrom.AsTime(),
		DateTo:   rq.DateTo.AsTime(),
	}
}

func (s *Server) toUpdateActivityDomain(rq *pb.UpdateActivityRequest) *domain.Activity {
	return &domain.Activity{
		Id:       rq.Id,
		Owner:    rq.Owner,
		Family:   rq.Family,
		Type:     rq.Type,
		DateFrom: rq.DateFrom.AsTime(),
		DateTo:   rq.DateTo.AsTime(),
	}
}


func (s *Server) toCreateActivityTypeDomain(rq *pb.CreateActivityTypeRequest) *domain.ActivityType {
	return &domain.ActivityType{
		Family:      rq.Family,
		Name:        rq.Name,
		Description: rq.Description,
	}
}

func (s *Server) toUpdateActivityTypeDomain(rq *pb.UpdateActivityTypeRequest) *domain.ActivityType {
	return &domain.ActivityType{
		Id:          rq.Id,
		Family:      rq.Family,
		Name:        rq.Name,
		Description: rq.Description,
	}
}

func (s *Server) toActivityTypePb(rq *domain.ActivityType) *pb.ActivityType {
	return &pb.ActivityType{
		Id:          rq.Id,
		Family:      rq.Family,
		Name:        rq.Name,
		Description: rq.Description,
	}
}
