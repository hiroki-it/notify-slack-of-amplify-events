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
func ThrowNew(exception string) error {
	return errors.New(exception)
}
