package file

import (
	"io/ioutil"

	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/usecases/exception"
)

func ReadTestData(path string) []byte {

	data, err := ioutil.ReadFile(path)

	if err != nil {
		exception.Error(err)
	}

	return data
}
