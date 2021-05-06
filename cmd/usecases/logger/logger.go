package logger

import (
	"go.uber.org/zap"
)

/**
 * エラーをロギングします．
 */
func ErrorLog(err error) {
	logger, _ := zap.NewDevelopment()
	logger.Error(err.Error())
}
