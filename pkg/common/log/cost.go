package log

import (
	"go.uber.org/zap"
	"time"
)

func CostLogs(l *zap.SugaredLogger, msg string) func() {
	startAt := time.Now()
	return func() {
		l.Infow(msg, "cast", time.Since(startAt).String())
	}
}
