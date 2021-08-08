package entities

import (
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/domain/file/values"
)

type File struct {
	path *values.Path
}

// NewFile コンストラクタ
func NewFile(path *values.Path) *File {

	return &File{
		path: path,
	}
}

// FilePath Pathを返却します．
func (f *File) FilePath() *values.Path {
	return f.path
}
