package value

import (
	"encoding/json"

	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/domain/core"
)

type BranchName struct {
	core.ValueObject
	value string
}

// NewBranchName コンストラクタ
func NewBranchName(value string) *BranchName {

	return &BranchName{
		value: value,
	}
}

// Equals 等価性を検証します．
func (bn *BranchName) Equals(target BranchName) bool {
	return bn.value == target.Value()
}

// Value 属性を返却します．
func (bn *BranchName) Value() string {
	return bn.value
}

// MarshalJSON 構造体をJSONに変換します．
func (bn *BranchName) MarshalJSON() ([]byte, error) {
	return json.Marshal(bn.Value())
}
