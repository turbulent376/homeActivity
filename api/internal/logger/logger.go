package logger

import (
	"github.com/turbulent376/homeactivity/api/internal/meta"
	"github.com/turbulent376/kit/log"
)

var Logger = log.Init(&log.Config{Level: log.TraceLevel})

func LF() log.CLoggerFunc {
	return func() log.CLogger {
		return log.L(Logger).Srv(meta.Meta)
	}
}

func L() log.CLogger {
	return LF()()
}
