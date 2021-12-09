package api

import (
	"context"

	"git.jetbrains.space/orbi/fcsd/api/internal/config"
	"git.jetbrains.space/orbi/fcsd/api/internal/logger"
	"git.jetbrains.space/orbi/fcsd/api/internal/meta"
	"git.jetbrains.space/orbi/fcsd/api/internal/public/auth"
	"git.jetbrains.space/orbi/fcsd/api/internal/public/billing"
	"git.jetbrains.space/orbi/fcsd/api/internal/public/license"
	"git.jetbrains.space/orbi/fcsd/api/internal/public/location"
	"git.jetbrains.space/orbi/fcsd/api/internal/public/middlewares"
	"git.jetbrains.space/orbi/fcsd/api/internal/public/report"
	"git.jetbrains.space/orbi/fcsd/api/internal/public/storage"
	"git.jetbrains.space/orbi/fcsd/api/internal/public/swagger"
	//"git.jetbrains.space/orbi/fcsd/api/internal/public/timesheet"
	authRep "git.jetbrains.space/orbi/fcsd/api/internal/repository/auth"
	billRep "git.jetbrains.space/orbi/fcsd/api/internal/repository/billing"
	licRep "git.jetbrains.space/orbi/fcsd/api/internal/repository/license"
	locRep "git.jetbrains.space/orbi/fcsd/api/internal/repository/location"
	repoRep "git.jetbrains.space/orbi/fcsd/api/internal/repository/report"
	storeRep "git.jetbrains.space/orbi/fcsd/api/internal/repository/storage"
	timesRep "git.jetbrains.space/orbi/fcsd/api/internal/repository/timesheet"
	kitHttp "git.jetbrains.space/orbi/fcsd/kit/http"
	"git.jetbrains.space/orbi/fcsd/kit/service"
	"golang.org/x/sync/errgroup"
)

type serviceImpl struct {
	cfg              *config.Config
	http             *kitHttp.Server
	authAdapter      authRep.Adapter
	billingAdapter   billRep.Adapter
	licenseAdapter   licRep.Adapter
	reportAdapter    repoRep.Adapter
	storageAdapter   storeRep.Adapter
	timesheetAdapter timesRep.Adapter
	locationAdapter  locRep.Adapter
}

func New() service.Service {
	s := &serviceImpl{}
	s.authAdapter = authRep.NewAdapter()
	s.billingAdapter = billRep.NewAdapter()
	s.licenseAdapter = licRep.NewAdapter()
	s.reportAdapter = repoRep.NewReportAdapter()
	s.storageAdapter = storeRep.NewStorageAdapter()
	s.timesheetAdapter = timesRep.NewAdapter()
	s.locationAdapter = locRep.NewAdapter()

	return s
}

func (s *serviceImpl) GetCode() string {
	return meta.Meta
}

func (s *serviceImpl) initHttp(cfg *config.Config) error {
	mdw := middlewares.NewMiddleware(s.authAdapter)

	s.http = kitHttp.NewHttpServer(cfg.Http, logger.LF())

	s.http.SetAuthMiddleware(mdw.AccessTokenAuthorizationMiddleware)
	s.http.SetNoAuthMiddleware(mdw.NoSessionMiddleware)

	// init controllers
	authController := auth.NewController(s.authAdapter)
	billingController := billing.NewController(s.billingAdapter)
	licenseController := license.NewController(s.licenseAdapter)
	reportController := report.NewController(s.reportAdapter)
	storageController := storage.NewController(s.storageAdapter)
	//timesheetController := timesheet.NewController(s.timesheetAdapter)
	locationController := location.NewController(s.locationAdapter)

	routers := []kitHttp.RouteSetter{
		auth.NewRouter(authController),
		billing.NewRouter(billingController),
		license.NewRouter(licenseController),
		report.NewRouter(reportController),
		storage.NewRouter(storageController),
		//timesheet.NewRouter(timesheetController),
		location.NewRouter(locationController),
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
		return s.billingAdapter.Init(s.cfg.Adapters["billing"])
	})

	grp.Go(func() error {
		return s.licenseAdapter.Init(s.cfg.Adapters["billing"])
	})

	grp.Go(func() error {
		return s.reportAdapter.Init(s.cfg.Adapters["report"])
	})

	grp.Go(func() error {
		return s.storageAdapter.Init(s.cfg.Adapters["storage"])
	})

	//grp.Go(func() error {
	//	return s.timesheetAdapter.Init(s.cfg.Adapters["timesheet"])
	//})

	grp.Go(func() error {
		return s.authAdapter.Init(s.cfg.Adapters["auth"])
	})

	grp.Go(func() error {
		return s.locationAdapter.Init(s.cfg.Adapters["location"])
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
