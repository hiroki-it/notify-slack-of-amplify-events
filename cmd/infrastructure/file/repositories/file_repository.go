package repositories

import (
	"io/ioutil"

	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/domain/file/entities"
)

type FileRepository struct {
}

// NewFileRepository コンストラクタ
func NewFileRepository() *FileRepository {

	return &FileRepository{}
}

// GetFile ファイルを読み出します．
func (fr *FileRepository) GetFile(file *entities.File) []byte {
	d, err := ioutil.ReadFile(file.FilePath().Path())

	if err != nil {
		return nil
	}

	return d
}
