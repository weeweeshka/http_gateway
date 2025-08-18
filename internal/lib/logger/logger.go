package logger

import "go.uber.org/zap"

func SetupLogger() *zap.Logger {
	logr, _ := zap.NewDevelopment()
	return logr
}
