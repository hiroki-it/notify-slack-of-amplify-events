package detail

import (
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/domain/core"
)

type AppId struct {
	*core.ID
	id string
}

// NewAppId コンストラクタ
func NewAppId(id string) *AppId {

	return &AppId{
		ID: &core.ID{},
		id: id,
	}
}

// Equals 等価性を検証します．
func (ai *AppId) Equals(target *AppId) bool {
	return ai.id == target.Id()
}
