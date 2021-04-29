package exception

import (
	"fmt"
)

func Error(err error) string {
	return fmt.Sprintf("ERROR: %#v\n", err)
}
