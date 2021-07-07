package file

import (
	"io/ioutil"

	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/infrastructure/logger"
)

// ReadFile ファイルを読み込みます．
func ReadFile(path string) []byte {

	log := logger.NewLogger()

	d, err := ioutil.ReadFile(path)

	if err != nil {
		log.Error(err.Error())
		return nil
	}

	return d
}
