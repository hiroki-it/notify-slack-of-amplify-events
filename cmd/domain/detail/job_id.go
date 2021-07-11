package detail

import (
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/domain/core"
)

type JobId struct {
	*core.ID
	id string
}

// NewJobId コンストラクタ
func NewJobId(id string) *JobId {

	return &JobId{
		ID: &core.ID{},
		id: id,
	}
}

// Equals 等価性を検証します．
func (ji *JobId) Equals(target *JobId) bool {
	return ji.id == target.Id()
}
