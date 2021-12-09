package config

import (
	"../../../kit/cache/redis"
	kitConfig "git.jetbrains.space/orbi/fcsd/kit/config"
	"git.jetbrains.space/orbi/fcsd/kit/db"
	"git.jetbrains.space/orbi/fcsd/kit/grpc"
	"git.jetbrains.space/orbi/fcsd/kit/log"
	"git.jetbrains.space/orbi/fcsd/kit/queue"
	"git.jetbrains.space/orbi/fcsd/kit/service"
	"git.jetbrains.space/orbi/fcsd/timesheet/internal/logger"
	"git.jetbrains.space/orbi/fcsd/timesheet/internal/meta"
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
