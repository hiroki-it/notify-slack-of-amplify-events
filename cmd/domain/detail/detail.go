package detail

import (
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/domain/core"
)

type Detail struct {
	core.Entity
	AppId      *AppId
	BranchName *BranchName
	JobId      *JobId
	JobStatus  *JobStatus
}

// NewDetail コンストラクタ
func NewDetail(appId *AppId, branchName *BranchName, jobId *JobId, jobStatus *JobStatus) *Detail {

	return &Detail{
		AppId:      appId,
		BranchName: branchName,
		JobId:      jobId,
		JobStatus:  jobStatus,
	}
}

// GetJobStatus JobStatusを返却します.
func (d *Detail) GetJobStatus() *JobStatus {
	return d.JobStatus
}

// IsSucceed ジョブステータスが成功かどうかを判定します
func (d *Detail) IsSucceed() bool {

	if d.JobStatus.Status() == "SUCCEED" {
		return true
	}

	return false
}
