package values

import (
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/domain"
)

type Path struct {
	path string
}

// NewPath コンストラクタ
func NewPath(path string) *Path {

	return &Path{
		path: path,
	}
}

// Path 属性を返却します．
func (p *Path) Path() string {
	return p.path
}

// Equals 等価性を検証します．
func (p *Path) Equals(target domain.ValueObject) bool {
	return p.path == target.(*Path).Path()
}
