package exception

import (
	"errors"
	"golang.org/x/xerrors"
)

/**
 * コンストラクタ
 * Exceptionを作成します．
 */
func NewException(err error, message string) *Exception {
	return &Exception{
		message:    errors.New(message),
		stackTrace: xerrors.Errorf("%w", err),
	}
}

/**
 * 例外メッセージを取得します．
 */
func (exception *Exception) GetMessage() error {
	return exception.message
}

/**
 * スタックトレースを取得します．
 */
func (exception *Exception) GetStackTrace() error {
	return exception.stackTrace
}
