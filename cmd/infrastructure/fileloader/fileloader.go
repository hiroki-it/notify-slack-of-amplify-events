package fileloader

import (
	"io/ioutil"

	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/domain/file"
)

type FileLoader struct {
	file *file.File
}

// NewFileLoader コンストラクタ
func NewFileLoader(file *file.File) *FileLoader {

	return &FileLoader{
		file: file,
	}
}

// ByteLoad 文字列型で返却します．．
func (l *FileLoader) ByteLoad() []byte {
	return l.loadFile()
}

// StringLoad 文字列型で返却します．
func (l *FileLoader) StringLoad() string {
	return string(l.loadFile())
}

// loadFile ファイルを読み込みます．
func (l *FileLoader) loadFile() []byte {

	d, err := ioutil.ReadFile(l.file.FilePath().Path())

	if err != nil {
		return nil
	}

	return d
}
