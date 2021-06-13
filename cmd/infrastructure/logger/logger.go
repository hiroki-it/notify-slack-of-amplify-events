package logger

import (
	"go.uber.org/zap"
)

/**
 * コンストラクタ
 * zap.Loggerを作成します．
 */
func NewLogger() *zap.Logger {
	logger, _ := zap.NewDevelopment()
	return logger
}
