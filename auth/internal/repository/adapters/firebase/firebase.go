package firebase

import (
	"context"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"git.jetbrains.space/orbi/fcsd/auth/internal/logger"
	"git.jetbrains.space/orbi/fcsd/kit/log"
)

type FirebaseAdapter interface {
	Init() error
	Close() error
	VerifyIDToken(ctx context.Context, idToken string) (*auth.Token, error)
}

type adapterImpl struct {
	client *auth.Client
}

func (a *adapterImpl) l() log.CLogger {
	return logger.L().Cmp("firebase-auth-adapter")
}

func (a *adapterImpl) Init() error {
	l := a.l().Mth("init-firebase-auth")

	l.Inf("Init")

	var err error
	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		return err
	}
	a.client, err = app.Auth(context.Background())
	if err != nil {
		return err
	}

	return nil
}

func (a *adapterImpl) VerifyIDToken(ctx context.Context, token string) (*auth.Token, error) {
	return a.client.VerifyIDToken(ctx, token)
}

func (a *adapterImpl) Close() error {
	return nil
}

func NewAdapter() FirebaseAdapter {
	return &adapterImpl{}
}
