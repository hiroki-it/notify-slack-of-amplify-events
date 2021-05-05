package file

import (
	"io/ioutil"

	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/usecases/logger"
)

func ReadTestDataFile(path string) []byte {

	data, err := ioutil.ReadFile(path)

	if err != nil {
		logger.Error(err)
	}

	return data
}
