package detail

import (
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/domain/core"
)

type AppId struct {
	*core.ID
}

// NewAppId コンストラクタ
func NewAppId(id string) *AppId {

	return &AppId{
		ID: &core.ID{},
	}
}

// Equals 等価性を検証します．
func (ai *AppId) Equals(target *AppId) bool {
	return ai.Id() == target.Id()
}
