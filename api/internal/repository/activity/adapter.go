package timesheet

import (
	"context"
	"git.jetbrains.space/orbi/fcsd/api/internal/config"
	"git.jetbrains.space/orbi/fcsd/api/internal/public"
	kitGrpc "git.jetbrains.space/orbi/fcsd/kit/grpc"
	pb "git.jetbrains.space/orbi/fcsd/proto/timesheet"
	"time"
)

const ReadyTimeout = time.Second * 3

type Adapter interface {
	public.TimesheetRepository
	Init(cfg *config.Adapter) error
	Close()
}

type adapterImpl struct {
	pb.TimesheetServiceClient
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

	if !a.client.AwaitReadiness(ReadyTimeout) {
		return kitGrpc.ErrGrpcSrvNotReady("billing")
	}

	a.TimesheetServiceClient = pb.NewTimesheetServiceClient(cl.Conn)

	return nil
}

func (a *adapterImpl) CreateTimesheet(ctx context.Context, rq *pb.CreateTimesheetRequest) (*pb.Timesheet, error) {
	res, err := a.TimesheetServiceClient.Create(ctx, rq)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (a *adapterImpl) UpdateTimesheet(ctx context.Context, rq *pb.UpdateTimesheetRequest) (*pb.Timesheet, error) {
	res, err := a.TimesheetServiceClient.Update(ctx, rq)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (a *adapterImpl) GetTimesheet(ctx context.Context, rq *pb.TimesheetIdRequest) (*pb.Timesheet, error) {
	res, err := a.TimesheetServiceClient.Get(ctx, rq)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (a *adapterImpl) SearchTimesheet(ctx context.Context, rq *pb.SearchTimesheetRequest) (*pb.Timesheet, error) {
	res, err := a.TimesheetServiceClient.Search(ctx, rq)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (a *adapterImpl) DeleteTimesheet(ctx context.Context, rq *pb.TimesheetIdRequest) error {
	_, err := a.TimesheetServiceClient.Delete(ctx, rq)

	if err != nil {
		return err
	}

	return nil
}

func (a *adapterImpl) CreateEvent(ctx context.Context, rq *pb.CreateEventRequest) (*pb.Event, error) {
	res, err := a.TimesheetServiceClient.CreateEvent(ctx, rq)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (a *adapterImpl) UpdateEvent(ctx context.Context, rq *pb.UpdateEventRequest) (*pb.Event, error) {
	res, err := a.TimesheetServiceClient.UpdateEvent(ctx, rq)

	if err != nil {
		return nil, err
	}

	return res, nil
}
func (a *adapterImpl) GetEvent(ctx context.Context, rq *pb.EventIdRequest) (*pb.Event, error) {
	res, err := a.TimesheetServiceClient.GetEvent(ctx, rq)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (a *adapterImpl) DeleteEvent(ctx context.Context, rq *pb.EventIdRequest) error {
	_, err := a.TimesheetServiceClient.DeleteEvent(ctx, rq)

	if err != nil {
		return err
	}

	return nil
}

func (a *adapterImpl) SearchEvents(ctx context.Context, rq *pb.EventIdRequest) (*pb.SearchResponse, error) {
	res, err := a.TimesheetServiceClient.SearchEvents(ctx, rq)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (a *adapterImpl) Close() {
	_ = a.client.Conn.Close()
}
