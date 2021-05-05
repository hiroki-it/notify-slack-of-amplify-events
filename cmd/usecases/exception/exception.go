package exception

import (
	"errors"
)

/**
 * 例外をスローします．
 */
func ThrowNew(exception string) error {
	return errors.New(exception)
}
