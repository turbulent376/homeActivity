package notification

import (
	"context"
	"time"

	"git.jetbrains.space/orbi/fcsd/auth/internal/config"
	"git.jetbrains.space/orbi/fcsd/auth/internal/logger"
	kitGrpc "git.jetbrains.space/orbi/fcsd/kit/grpc"
	"git.jetbrains.space/orbi/fcsd/kit/log"

	pb "git.jetbrains.space/orbi/fcsd/proto/notification"
)

const ReadyTimeout = time.Second * 3

type NotificationAdapter interface {
	Init(cfg *config.Adapter) error
	Close()
	SendNotify(ctx context.Context, notifToken string) (*pb.MessageResponse, error)
}

// adapterImpl implements notification adapter
type adapterImpl struct {
	pb.NotificationClient
	client *kitGrpc.Client
}

// NewAdapter creates a new instance of the adapter
func NewAdapter() NotificationAdapter {
	return &adapterImpl{}

}

func (a *adapterImpl) l() log.CLogger {
	return logger.L().Srv("auth").Cmp("adapter-notification")
}

func (a *adapterImpl) Init(cfg *config.Adapter) error {
	cl, err := kitGrpc.NewClient(cfg.Grpc)

	if err != nil {
		return err
	}

	a.client = cl

	if !a.client.AwaitReadiness(ReadyTimeout) {
		return kitGrpc.ErrGrpcSrvNotReady("notification")
	}

	a.NotificationClient = pb.NewNotificationClient(cl.Conn)

	return nil
}

func (a *adapterImpl) SendNotify(ctx context.Context, notifToken string) (*pb.MessageResponse, error) {
	a.l().Mth("SendNotify").C(ctx)
	rq := &pb.MessageRequest{
		Address:     notifToken,
		MessageType: "firebase",
		TemplateId: "test_template",
	}

	res, err := a.NotificationClient.NewMessage(ctx, rq)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (a *adapterImpl) Close() {
	_ = a.client.Conn.Close()
}
