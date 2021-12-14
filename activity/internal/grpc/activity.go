package grpc

import (
	"context"
	pb "github.com/turbulent376/proto/activity"
)

func (s *Server) Create(ctx context.Context, rq *pb.CreateActivityRequest) (*pb.Activity, error) {
	activity, err := s.activityService.Create(ctx, s.toCreateActivityDomain(rq))
	if err != nil {
		return nil, err
	}
	return s.toActivityPb(activity), nil
}

func (s *Server) Update(ctx context.Context, rq *pb.UpdateActivityRequest) (*pb.Activity, error) {
	activity, err := s.activityService.Update(ctx, s.toUpdateActivityDomain(rq))
	if err != nil {
		return nil, err
	}
	return s.toActivityPb(activity), nil
}

func (s *Server) Get(ctx context.Context, rq *pb.ActivityIdRequest) (*pb.Activity, error) {
	found, activity, err := s.activityService.Get(ctx, rq.Id)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, nil
	}
	return s.toActivityPb(activity), nil
}


func (s *Server) Delete(ctx context.Context, rq *pb.ActivityIdRequest) (*pb.EmptyResponse, error) {
	err := s.activityService.Delete(ctx, rq.Id)
	if err != nil {
		return nil, err
	}
	return &pb.EmptyResponse{}, nil
}

func (s *Server) CreateActivityType(ctx context.Context, rq *pb.CreateActivityTypeRequest) (*pb.ActivityType, error) {
	res, err := s.activityService.CreateActivityType(ctx, s.toCreateActivityTypeDomain(rq))

	if err != nil {
		return nil, err
	}

	return s.toActivityTypePb(res), nil
}

func (s *Server) UpdateActivityType(ctx context.Context, rq *pb.UpdateActivityTypeRequest) (*pb.ActivityType, error) {
	res, err := s.activityService.UpdateActivityType(ctx, s.toUpdateActivityTypeDomain(rq))

	if err != nil {
		return nil, err
	}

	return s.toActivityTypePb(res), nil
}

func (s *Server) GetActivityType(ctx context.Context, rq *pb.ActivityTypeIdRequest) (*pb.ActivityType, error) {
	found, res, err := s.activityService.GetActivityType(ctx, rq.Id)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, nil
	}
	return s.toActivityTypePb(res), nil
}

func (s *Server) DeleteActivityType(ctx context.Context, rq *pb.ActivityTypeIdRequest) (*pb.EmptyResponse, error) {
	err := s.activityService.DeleteActivityType(ctx, rq.Id)

	if err != nil {
		return nil, err
	}

	return &pb.EmptyResponse{}, nil
}
