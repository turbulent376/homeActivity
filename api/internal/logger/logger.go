package logger

import (
	"git.jetbrains.space/orbi/fcsd/api/internal/meta"
	"git.jetbrains.space/orbi/fcsd/kit/log"
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
