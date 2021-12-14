package logger

import (
	"github.com/turbulent376/kit/log"
	"github.com/turbulent376/homeactivity/activity/internal/meta"
)

var Logger = log.Init(&log.Config{Level: log.TraceLevel})

func LF() log.CLoggerFunc {
	return func() log.CLogger {
		return log.L(Logger).Srv(meta.Meta.InstanceId())
	}
}

func L() log.CLogger {
	return LF()()
}
