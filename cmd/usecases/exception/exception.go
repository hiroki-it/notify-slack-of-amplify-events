package exception

import (
	"golang.org/x/xerrors"
)

func Error(err error) error {
	return xerrors.Errorf("%w", err)
}
