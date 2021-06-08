package file

import (
	"io/ioutil"
)

func ReadDataFile(path string) ([]byte, error) {

	data, err := ioutil.ReadFile(path)

	if err != nil {
		return nil, err
	}

	return data, nil
}
