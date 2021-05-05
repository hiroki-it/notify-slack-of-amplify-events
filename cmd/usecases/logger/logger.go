package logger

import (
	"golang.org/x/xerrors"
	"log"
)

/**
 * エラーをロギングします．
 */
func ErrorLog(err error) {
	trace := xerrors.Errorf("%w", err)
	log.Printf("ERROR: %+v\n", trace)
}
