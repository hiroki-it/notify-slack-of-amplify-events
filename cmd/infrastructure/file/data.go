package file

import (
	"io/ioutil"

	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/infrastructure/logger"
)

func ReadDataFile(path string) []byte {

	log := logger.NewLogger()

	data, err := ioutil.ReadFile(path)

	if err != nil {
		log.Error(err.Error())
		return nil
	}

	return data
}
