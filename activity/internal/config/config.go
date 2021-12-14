package config

import (
	"github.com/turbulent376/kit/cache/redis"
	kitConfig "github.com/turbulent376/kit/config"
	"github.com/turbulent376/kit/db"
	"github.com/turbulent376/kit/grpc"
	"github.com/turbulent376/kit/log"
	"github.com/turbulent376/kit/queue"
	"github.com/turbulent376/kit/service"
	"github.com/turbulent376/homeactivity/activity/internal/logger"
	"github.com/turbulent376/homeactivity/activity/internal/meta"
	"os"
	"path/filepath"
)

// Here are defined all types for your configuration
// You can remove not needed types or add your own

type Storages struct {
	Redis    *redis.Config
	Database *db.DbClusterConfig
}

type Config struct {
	Grpc     *grpc.ServerConfig
	Storages *Storages
	Nats     *queue.Config
	Log      *log.Config
	Cluster  *service.Config
}

func Load() (*Config, error) {

	// get root folder from env
	rootPath := os.Getenv("FOCROOT")
	if rootPath == "" {
		return nil, kitConfig.ErrConfigPaErrConfigPathIsEmpty()
	}

	// config path
	configPath := filepath.Join(rootPath, meta.Meta.ServiceCode(), "config.yml")

	// load config
	config := &Config{}
	err := kitConfig.NewConfigLoader(logger.LF()).
		WithConfigPath(configPath).
		Load(config)

	if err != nil {
		return nil, err
	}
	return config, nil
}
