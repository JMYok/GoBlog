package utils

import "go.uber.org/zap"

var Logger *zap.Logger

func init() {
	Logger, _ = zap.NewDevelopment()
	Logger.Info("Logger has initialized!")
}
