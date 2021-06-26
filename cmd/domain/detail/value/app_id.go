package detail

import (
	"encoding/json"

	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/domain/core"
)

type AppId struct {
	core.ValueObject
	value string
}

// NewAppId コンストラクタ
func NewAppId(value string) *AppId {

	return &AppId{
		value: value,
	}
}

// Equals 等価性を検証します．
func (ai AppId) Equals(target AppId) bool {
	return ai.value == target.value
}

// Value 属性を返却します．
func (ai AppId) Value() string {
	return ai.value
}

// MarshalJSON 構造体をJSONに変換します．
func (ai AppId) MarshalJSON() ([]byte, error) {
	return json.Marshal(ai.value)
}
