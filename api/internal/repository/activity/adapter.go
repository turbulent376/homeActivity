package activity

import (
	"context"
	"github.com/turbulent376/homeactivity/api/internal/config"
	"github.com/turbulent376/homeactivity/api/internal/public"
	kitGrpc "github.com/turbulent376/kit/grpc"
	pb "github.com/turbulent376/proto/activity"
	"time"
)

const ReadyTimeout = time.Second * 3

type Adapter interface {
	public.ActivityRepository
	Init(cfg *config.Adapter) error
	Close()
}

type adapterImpl struct {
	pb.ActivityServiceClient
	client *kitGrpc.Client
}

func NewAdapter() Adapter {
	return &adapterImpl{}
}

func (a *adapterImpl) Init(cfg *config.Adapter) error {
	cl, err := kitGrpc.NewClient(cfg.Grpc)

	if err != nil {
		return err
	}

	a.client = cl

	a.ActivityServiceClient = pb.NewActivityServiceClient(cl.Conn)

	return nil
}

func (a *adapterImpl) CreateActivity(ctx context.Context, rq *pb.CreateActivityRequest) (*pb.Activity, error) {
	res, err := a.ActivityServiceClient.Create(ctx, rq)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (a *adapterImpl) UpdateActivity(ctx context.Context, rq *pb.UpdateActivityRequest) (*pb.Activity, error) {
	res, err := a.ActivityServiceClient.Update(ctx, rq)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (a *adapterImpl) GetActivity(ctx context.Context, rq *pb.ActivityIdRequest) (*pb.Activity, error) {
	res, err := a.ActivityServiceClient.Get(ctx, rq)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (a *adapterImpl) ListActivities(ctx context.Context, rq *pb.ListActivitiesRequest) (*pb.ListActivitiesResponse, error) {
	res, err := a.ActivityServiceClient.ListActivities(ctx, rq)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (a *adapterImpl) ListActivitiesByFamily(ctx context.Context, rq *pb.ListActivitiesByFamilyRequest) (*pb.ListActivitiesResponse, error) {
	res, err := a.ActivityServiceClient.ListActivitiesByFamily(ctx, rq)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (a *adapterImpl) DeleteActivity(ctx context.Context, rq *pb.ActivityIdRequest) error {
	_, err := a.ActivityServiceClient.Delete(ctx, rq)

	if err != nil {
		return err
	}

	return nil
}

func (a *adapterImpl) CreateActivityType(ctx context.Context, rq *pb.CreateActivityTypeRequest) (*pb.ActivityType, error) {
	res, err := a.ActivityServiceClient.CreateActivityType(ctx, rq)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (a *adapterImpl) UpdateActivityType(ctx context.Context, rq *pb.UpdateActivityTypeRequest) (*pb.ActivityType, error) {
	res, err := a.ActivityServiceClient.UpdateActivityType(ctx, rq)

	if err != nil {
		return nil, err
	}

	return res, nil
}
func (a *adapterImpl) GetActivityType(ctx context.Context, rq *pb.ActivityTypeIdRequest) (*pb.ActivityType, error) {
	res, err := a.ActivityServiceClient.GetActivityType(ctx, rq)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (a *adapterImpl) DeleteActivityType(ctx context.Context, rq *pb.ActivityTypeIdRequest) error {
	_, err := a.ActivityServiceClient.DeleteActivityType(ctx, rq)

	if err != nil {
		return err
	}

	return nil
}

func (a *adapterImpl) ListActivityTypes(ctx context.Context, rq *pb.ListActivityTypesRequest) (*pb.ListActivityTypesResponse, error) {
	res, err := a.ActivityServiceClient.ListActivityTypes(ctx, rq)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (a *adapterImpl) Close() {
	_ = a.client.Conn.Close()
}
