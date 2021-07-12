package detail

import (
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/domain"
)

type AppId struct {
	id string
}

var _ domain.ID = &AppId{}

// NewAppId コンストラクタ
func NewAppId(id string) *AppId {

	return &AppId{
		id: id,
	}
}

// Id 属性を返却します．
func (ai *AppId) Id() string {
	return ai.id
}

// Equals 等価性を検証します．
func (ai *AppId) Equals(target domain.ID) bool {
	return ai.id == target.(*AppId).Id()
}
