package detail

import (
	"encoding/json"

	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/domain/core"
)

type JobId struct {
	core.ValueObject
	value string
}

// NewJobId コンストラクタ
func NewJobId(value string) *JobId {

	return &JobId{
		value: value,
	}
}

// Equals 等価性を検証します．
func (ji *JobId) Equals(target JobId) bool {
	return ji.value == target.Value()
}

// Value 属性を返却します．
func (ji *JobId) Value() string {
	return ji.value
}

// MarshalJSON 構造体をJSONに変換します．
func (ji *JobId) MarshalJSON() ([]byte, error) {
	return json.Marshal(ji.Value())
}
