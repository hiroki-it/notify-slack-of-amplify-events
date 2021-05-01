package exception

import (
	"log"
)

func Error(err error) {
	log.Fatalf("ERROR: %#v\n", err)
}
