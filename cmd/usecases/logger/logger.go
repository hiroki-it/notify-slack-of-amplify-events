package logger

import (
	"log"

	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/entities/exception"
)

/**
 * エラーをロギングします．
 */
func ErrorLog(exception *exception.Exception) {
	log.Printf(
		"ERROR: %s %+v\n",
		exception.GetMessage(),
		exception.GetStackTrace(),
	)
}
