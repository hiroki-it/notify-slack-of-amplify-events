package detail

import (
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/domain/core"
)

type AppId struct {
	core.ID
	id string
}

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
func (ai *AppId) Equals(target *AppId) bool {
	return ai.id == target.Id()
}
