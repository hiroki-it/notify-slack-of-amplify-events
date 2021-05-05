package logger

import (
	"log"
)

/**
 * エラーをロギングします．
 */
func ErrorLog(err error) {
	log.Printf("ERROR: %+v\n", err)
}
