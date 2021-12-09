package config

import (
	"git.jetbrains.space/orbi/fcsd/api/internal/logger"
	"git.jetbrains.space/orbi/fcsd/api/internal/meta"
	kitConfig "git.jetbrains.space/orbi/fcsd/kit/config"
	"git.jetbrains.space/orbi/fcsd/kit/grpc"
	kitHttp "git.jetbrains.space/orbi/fcsd/kit/http"
	"git.jetbrains.space/orbi/fcsd/kit/log"
	"os"
	"path/filepath"
)

type Adapter struct {
	Grpc *grpc.ClientConfig
}

type Config struct {
	Log        *log.Config
	Adapters   map[string]*Adapter
	Http       *kitHttp.Config
}

func Load() (*Config, error) {

	// get root folder from env
	rootPath := os.Getenv("FOCROOT")
	if rootPath == "" {
		return nil, kitConfig.ErrConfigPaErrConfigPathIsEmpty()
	}

	// config path
	configPath := filepath.Join(rootPath, meta.Meta, "config.yml")

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
