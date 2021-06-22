package logger

import (
	"go.uber.org/zap"
)

// NewLogger コンストラクタ
func NewLogger() *zap.Logger {
	logger, _ := zap.NewDevelopment()
	return logger
}
