package entity

import (
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/domain/detail/value"
)

type Detail struct {
	AppId      *value.AppId      `json:"appId"`
	BranchName *value.BranchName `json:"branchName"`
	JobId      *value.JobId      `json:"jobId"`
	JobStatus  *value.JobStatus  `json:"jobStatus"`
}

// NewDetail コンストラクタ
func NewDetail() *Detail {

	return &Detail{}
}

// GetJobStatus JobStatusを返却します.
func (d *Detail) GetJobStatus() *value.JobStatus {
	return d.JobStatus
}

// IsSucceed ジョブステータスが成功かどうかを判定します
func (d *Detail) IsSucceed() bool {

	if d.JobStatus.Value() == "SUCCEED" {
		return true
	}

	return false
}
