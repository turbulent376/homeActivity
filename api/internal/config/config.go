package config

import (
	"github.com/turbulent376/homeactivity/api/internal/logger"
	"github.com/turbulent376/homeactivity/api/internal/meta"
	kitConfig "github.com/turbulent376/kit/config"
	"github.com/turbulent376/kit/grpc"
	kitHttp "github.com/turbulent376/kit/http"
	"github.com/turbulent376/kit/log"
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
