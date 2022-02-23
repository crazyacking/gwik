package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var (
	rootLogger *zap.SugaredLogger
	atom       zap.AtomicLevel
)

func init() {
	initConfigMode()

	atom = zap.NewAtomicLevel()
	var core zapcore.Core
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	fileHook := newFileLog()
	writeSyncer := zapcore.AddSync(fileHook)

	switch logMode {
	case productionMode:
		core = zapcore.NewTee(
			zapcore.NewCore(
				zapcore.NewJSONEncoder(encoderCfg),
				writeSyncer,
				atom,
			),
			zapcore.NewCore(
				zapcore.NewJSONEncoder(encoderCfg),
				zapcore.AddSync(os.Stdout),
				atom,
			),
		)
	case developmentMode:
		zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderCfg),
			zapcore.AddSync(os.Stdout),
			zapcore.DebugLevel,
		)
	case testMode:
		zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderCfg),
			zapcore.AddSync(os.Stdout),
			zapcore.ErrorLevel,
		)
	}

	logger := zap.New(core)
	rootLogger = logger.Sugar()

}

func Sync() {
	_ = rootLogger.Sync()
}

func DebugLog(enable bool) {
	if enable {
		atom.SetLevel(zapcore.DebugLevel)
		return
	}
	atom.SetLevel(zapcore.InfoLevel)
}

func NewLogger(name string) *zap.SugaredLogger {
	return rootLogger.Named(name)
}
