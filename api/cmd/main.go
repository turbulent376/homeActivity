package main

import (
	"context"
	"github.com/turbulent376/homeactivity/api"
	"github.com/turbulent376/homeactivity/api/internal/logger"
	kitContext "github.com/turbulent376/kit/context"
	"os"
	"os/signal"
	"syscall"
)

// @title Focusednow swagger
// @version 0.0.1
// @contact.name Nikita Zaitsev
// @contact.email nikita.zaitsev@teamlabs.cc
// @BasePath /api
func main() {
	// init context
	ctx := kitContext.NewRequestCtx().Empty().WithNewRequestId().ToContext(context.Background())

	// create a new service
	s := api.New()

	l := logger.L().Mth("main").Inf("created")

	// init service
	if err := s.Init(ctx); err != nil {
		l.E(err).St().Err("initialization")
		os.Exit(1)
	}

	l.Inf("initialized")

	// start listening
	if err := s.Start(ctx); err != nil {
		l.E(err).St().Err("listen")
		os.Exit(1)
	}

	l.Inf("listening")

	// handle app close
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	l.Inf("quit signal")
	s.Close(ctx)
	os.Exit(0)
}
