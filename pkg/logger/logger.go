package logger

import "go.uber.org/zap"

var log *zap.SugaredLogger

func Init(service string) *zap.SugaredLogger {
	logger, _ := zap.NewProduction()
	log = logger.Sugar().With("service", service)
	return log
}

func Info(args ...interface{})  { log.Info(args...) }
func Error(args ...interface{}) { log.Error(args...) }
func Sync()                     { _ = log.Sync() }
