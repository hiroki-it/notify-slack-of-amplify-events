package logger

import (
	"go.uber.org/zap"
)

// NewLogger コンストラクタ
func NewLogger() *zap.Logger {
	log, _ := zap.NewDevelopment()
	return log
}
