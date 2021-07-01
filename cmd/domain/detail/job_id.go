package detail

import (
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/domain/core"
)

type JobId struct {
	core.ID
	id string
}

// NewJobId コンストラクタ
func NewJobId(id string) *JobId {

	return &JobId{
		id: id,
	}
}

// Id 属性を返却します．
func (ji *JobId) Id() string {
	return ji.id
}

// Equals 等価性を検証します．
func (ji *JobId) Equals(target *JobId) bool {
	return ji.id == target.Id()
}
