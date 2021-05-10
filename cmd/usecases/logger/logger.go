package logger

import (
	"go.uber.org/zap"
)

/**
 * エラーをロギングします．
 */
func NewLogger() *zap.Logger {
	logger, _ := zap.NewDevelopment()
	return logger
}
