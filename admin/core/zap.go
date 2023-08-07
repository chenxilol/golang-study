package core

import (
	"go.uber.org/zap"
	"net/http"
)

var logger *zap.Logger

func InitLoggers() *zap.Logger {
	logger, _ = zap.NewProduction()
	return logger
}
func SimpleHttpGet(url string) {
	resp, err := http.Get(url)
	if err != nil {
		logger.Error(
			"Error fetching url..",
			zap.String("url", url),
			zap.Error(err))
	} else {
		logger.Info("Success..",
			zap.String("statusCode", resp.Status),
			zap.String("url", url))
		resp.Body.Close()
	}
}
