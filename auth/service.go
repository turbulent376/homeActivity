package auth

import (
	"context"

	"git.jetbrains.space/orbi/fcsd/kit/queue"
	"git.jetbrains.space/orbi/fcsd/kit/queue/stan"
	"git.jetbrains.space/orbi/fcsd/kit/service"

	"git.jetbrains.space/orbi/fcsd/auth/internal/config"
	"git.jetbrains.space/orbi/fcsd/auth/internal/domain"
	"git.jetbrains.space/orbi/fcsd/auth/internal/domain/impl"
	"git.jetbrains.space/orbi/fcsd/auth/internal/grpc"
	"git.jetbrains.space/orbi/fcsd/auth/internal/logger"
	"git.jetbrains.space/orbi/fcsd/auth/internal/meta"
	"git.jetbrains.space/orbi/fcsd/auth/internal/repository/adapters/firebase"
	"git.jetbrains.space/orbi/fcsd/auth/internal/repository/adapters/notification"
	"git.jetbrains.space/orbi/fcsd/auth/internal/repository/storage"
)

// serviceImpl implements a service bootstrapping
// all dependencies between layers must be specified here
type serviceImpl struct {
	service.Cluster
	cfg                 *config.Config
	authService         domain.AuthService
	grpc                *grpc.Server
	storageAdapter      storage.Adapter
	notificationAdapter notification.NotificationAdapter
	firebaseAdapter     firebase.FirebaseAdapter
	queue               queue.Queue
}

// New creates a new instance of the service
func New() service.Service {

	s := &serviceImpl{
		Cluster: service.NewCluster(logger.LF(), meta.Meta),
	}

	s.queue = stan.New(logger.LF())
	s.storageAdapter = storage.NewAdapter()

	s.notificationAdapter = notification.NewAdapter()
	s.firebaseAdapter = firebase.NewAdapter()
	s.authService = impl.NewAuthService(s.storageAdapter.GetUserStorage(),
		s.storageAdapter.GetSessionStorage(), s.notificationAdapter, s.firebaseAdapter)

	s.grpc = grpc.New(s.authService)

	return s
}

func (s *serviceImpl) GetCode() string {
	return meta.Meta.ServiceCode()
}

// Init does all initializations
func (s *serviceImpl) Init(ctx context.Context) error {
	// load config
	var err error
	s.cfg, err = config.Load()
	if err != nil {
		return err
	}

	// set log config
	logger.Logger.Init(s.cfg.Log)

	// init cluster
	if err := s.Cluster.Init(s.cfg.Cluster, s.cfg.Nats.Host, s.cfg.Nats.Port, s.onClusterLeaderChanged(ctx)); err != nil {
		return err
	}

	// init storage
	if err := s.storageAdapter.Init(s.cfg.Storages); err != nil {
		return err
	}
	// init notification
	if err := s.notificationAdapter.Init(s.cfg.Adapters["notification"]); err != nil {
		return err
	}

	// init firebase for auth
	if err := s.firebaseAdapter.Init(); err != nil {
		return err
	}

	// init sing key
	if err := s.authService.SetSignKey(ctx, s.cfg.Auth.KeyRS256); err != nil {
		return err
	}

	// init grpc server
	if err := s.grpc.Init(s.cfg.Grpc); err != nil {
		return err
	}

	// open Queue connection
	if err := s.queue.Open(ctx, meta.Meta.InstanceId(), s.cfg.Nats); err != nil {
		return err
	}

	return nil

}

func (s *serviceImpl) onClusterLeaderChanged(ctx context.Context) service.OnLeaderChangedEvent {

	// if the current node is getting a leader, run daemons
	return func(l bool) {
		if l {
			// do something if the node is turned into a leader
			logger.L().C(ctx).Cmp("cluster").Mth("on-leader-change").Dbg("leader")
		}
	}

}

func (s *serviceImpl) Start(ctx context.Context) error {

	// start cluster
	if err := s.Cluster.Start(); err != nil {
		return err
	}

	// serve gRPC connection
	s.grpc.ListenAsync()

	return nil
}

func (s *serviceImpl) Close(ctx context.Context) {
	s.Cluster.Close()
	_ = s.queue.Close()
	s.storageAdapter.Close()
	s.grpc.Close()
}
