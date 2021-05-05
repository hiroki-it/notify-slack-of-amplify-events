package exception

import (
	"errors"
	"golang.org/x/xerrors"
)

/**
 * コンストラクタ
 * Exceptionを作成します．
 */
func NewException(err error) *Exception {
	return &Exception{
		trace: xerrors.Errorf("%w", err),
	}
}

/**
 * 例外をスローします．
 */
func (exception *Exception) Throw(message string) error {
	return errors.New(message)
}

/**
 * スタックトレースを取得します．
 */
func (exception *Exception) GetTrace() error {
	return exception.trace
}
