package value

import (
	"encoding/json"

	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/domain/core"
)

type JobStatus struct {
	core.ValueObject
	value string
}

// NewJobStatus コンストラクタ
func NewJobStatus(value string) *JobStatus {

	return &JobStatus{
		value: value,
	}
}

// Equals 等価性を検証します．
func (js *JobStatus) Equals(target JobStatus) bool {
	return js.value == target.Value()
}

// Value 属性を返却します．
func (js *JobStatus) Value() string {
	return js.value
}

// MarshalJSON 構造体をJSONに変換します．
func (js JobStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(js.Value())
}
