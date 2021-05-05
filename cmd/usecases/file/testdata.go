package file

import (
	"io/ioutil"

	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/usecases/exception"
)

func ReadTestDataFile(path string) ([]byte, *exception.Exception) {

	data, err := ioutil.ReadFile(path)

	if err != nil {
		return nil, exception.NewException(err, "Failed to read test data file.")
	}

	return data, nil
}
