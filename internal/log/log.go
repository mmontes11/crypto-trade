package log

import "go.uber.org/zap"

// LoggerI defines the different log level operations to be implemented by the logger
type LoggerI interface {
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Warn(args ...interface{})
	Warnf(format string, args ...interface{})
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
}

// Init creates a new logger instance
func Init(env string) LoggerI {
	var zapConfig zap.Config
	if env == "development" {
		zapConfig = zap.NewDevelopmentConfig()
	} else {
		zapConfig = zap.NewProductionConfig()
	}
	zapConfig.DisableCaller = true

	var logger *zap.Logger
	logger, _ = zapConfig.Build()
	defer logger.Sync()

	return logger.Sugar()
}
