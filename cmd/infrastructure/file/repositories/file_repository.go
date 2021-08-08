package repositories

import (
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/domain/file/entities"
	"io/ioutil"
)

type FileRepository struct {
	file *entities.File
}

// NewFileRepository コンストラクタ
func NewFileRepository(file *entities.File) *FileRepository {

	return &FileRepository{
		file: file,
	}
}

// ByteLoad 文字列型で返却します．．
func (l *FileRepository) ByteLoad() []byte {
	return l.loadFile()
}

// StringLoad 文字列型で返却します．
func (l *FileRepository) StringLoad() string {
	return string(l.loadFile())
}

// loadFile ファイルを読み込みます．
func (l *FileRepository) loadFile() []byte {

	d, err := ioutil.ReadFile(l.file.FilePath().Path())

	if err != nil {
		return nil
	}

	return d
}
