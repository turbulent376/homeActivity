package api

import (
	"context"

	"github.com/turbulent376/homeactivity/api/internal/config"
	//"github.com/turbulent376/homeactivity/api/internal/logger"
	"github.com/turbulent376/homeactivity/api/internal/meta"
	"github.com/turbulent376/homeactivity/api/internal/public/auth"
	"github.com/turbulent376/homeactivity/api/internal/public/activity"
	"github.com/turbulent376/homeactivity/api/internal/public/swagger"
	authRep "github.com/turbulent376/homeactivity/api/internal/repository/auth"
	activRep "github.com/turbulent376/homeactivity/api/internal/repository/activity"
	kitHttp "github.com/turbulent376/kit/http"
	"github.com/turbulent376/kit/service"
	"golang.org/x/sync/errgroup"
)

type serviceImpl struct {
	cfg              *config.Config
	http             *kitHttp.Server
	authAdapter      authRep.Adapter
	activityAdapter  activRep.Adapter
}

func New() service.Service {
	s := &serviceImpl{}
	s.authAdapter = authRep.NewAdapter()
	s.activityAdapter = activRep.NewAdapter()

	return s
}

func (s *serviceImpl) GetCode() string {
	return meta.Meta
}

func (s *serviceImpl) initHttp(cfg *config.Config) error {

	// init controllers
	authController := auth.NewController(s.authAdapter)
	activityController := activity.NewController(s.activityAdapter)

	routers := []kitHttp.RouteSetter{
		auth.NewRouter(authController),
		activity.NewRouter(activityController),
		swagger.NewRouter(),
	}

	s.http.SetRouters(routers...)

	for _, r := range routers {
		if cr, ok := r.(kitHttp.CustomRouteSetter); ok {
			cr.SetCustom(s.http.RootRouter)
		}
	}

	return nil
}

func (s *serviceImpl) initAdapters() error {
	grp, _ := errgroup.WithContext(context.Background())

	grp.Go(func() error {
		return s.activityAdapter.Init(s.cfg.Adapters["activity"])
	})

	grp.Go(func() error {
		return s.authAdapter.Init(s.cfg.Adapters["auth"])
	})

	return grp.Wait()
}

func (s *serviceImpl) Init(ctx context.Context) error {
	var err error
	s.cfg, err = config.Load()

	if err != nil {
		return err
	}

	if err := s.initAdapters(); err != nil {
		return err
	}

	if err := s.initHttp(s.cfg); err != nil {
		return err
	}

	return nil
}

func (s *serviceImpl) Start(ctx context.Context) error {
	s.http.Listen()

	return nil
}

func (s *serviceImpl) Close(ctx context.Context) {
	s.http.Close()
}
