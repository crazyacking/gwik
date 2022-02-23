package log

import (
	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	logFile = "/var/log/gwik/gwik.log"
)

func newFileLog() *lumberjack.Logger {
	hook := lumberjack.Logger{
		Filename:   logFile,
		MaxSize:    100,  // MB
		MaxBackups: 60,   //files
		MaxAge:     30,   // days
		Compress:   true, // use gzip
	}
	return &hook
}
