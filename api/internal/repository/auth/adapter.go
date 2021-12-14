package auth

import (
	"context"
	"time"

	"github.com/turbulent376/homeactivity/api/internal/config"
	"github.com/turbulent376/homeactivity/api/internal/public"
	kitGrpc "github.com/turbulent376/kit/grpc"
	pb "github.com/turbulent376/proto/auth"
)

const ReadyTimeout = time.Second * 3

type Adapter interface {
	public.AuthRepository
	Init(cfg *config.Adapter) error
	Close()
}

type adapterImpl struct {
	pb.AuthServiceClient
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
		return kitGrpc.ErrGrpcSrvNotReady("auth")
	}

	a.AuthServiceClient = pb.NewAuthServiceClient(cl.Conn)

	return nil
}

func (a *adapterImpl) AuthUserByEmail(ctx context.Context, rq *pb.AuthRequest) (*pb.AuthResponse, error) {
	res, err := a.AuthServiceClient.AuthUserByEmail(ctx, rq)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (a *adapterImpl) AuthUserByFirebase(ctx context.Context, rq *pb.OAuthRequest) (*pb.AuthResponse, error) {
	res, err := a.AuthServiceClient.AuthUserByFirebase(ctx, rq)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (a *adapterImpl) RefreshToken(ctx context.Context, rq *pb.RefreshTokenRequest) (*pb.TokenPairResponse, error) {
	res, err := a.AuthServiceClient.RefreshToken(ctx, rq)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (a *adapterImpl) CreateUser(ctx context.Context, rq *pb.CreateUserRequest) (*pb.User, error) {
	res, err := a.AuthServiceClient.CreateUser(ctx, rq)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (a *adapterImpl) GetUserById(ctx context.Context, rq *pb.UserIdRequest) (*pb.User, error) {
	res, err := a.AuthServiceClient.GetUserById(ctx, rq)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (a *adapterImpl) GetSessionByToken(ctx context.Context, rq *pb.TokenRequest) (*pb.Session, error) {
	res, err := a.AuthServiceClient.GetSessionByToken(ctx, rq)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (a *adapterImpl) UpdateUser(ctx context.Context, rq *pb.UpdateUserRequest) (*pb.User, error) {
	res, err := a.AuthServiceClient.UpdateUser(ctx, rq)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (a *adapterImpl) DeleteUser(ctx context.Context, rq *pb.UserIdRequest) error {
	_, err := a.AuthServiceClient.DeleteUser(ctx, rq)

	if err != nil {
		return err
	}

	return nil
}

func (a *adapterImpl) CloseSession(ctx context.Context, rq *pb.CloseSessionRequest) error {
	_, err := a.AuthServiceClient.CloseSession(ctx, rq)

	if err != nil {
		return err
	}

	return nil
}

func (a *adapterImpl) SaveUserFCMToken(ctx context.Context, rq *pb.FCMTokenRequest) error {
	_, err := a.AuthServiceClient.SaveUserFCMToken(ctx, rq)

	if err != nil {
		return err
	}

	return nil
}

func (a *adapterImpl) Close() {
	_ = a.client.Conn.Close()
}
