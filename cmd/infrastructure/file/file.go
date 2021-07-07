package file

import (
	"io/ioutil"

	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/infrastructure/logger"
)

type File struct {
	content []byte
}

// NewFile コンストラクタ
func NewFile() *File {
	return &File{}
}

// ReadFile ファイルを読み込みます．
func (f *File) ReadFile(path string) *File {

	log := logger.NewLogger()

	d, err := ioutil.ReadFile(path)

	if err != nil {
		log.Error(err.Error())
		return nil
	}

	f.content = d

	return f
}
