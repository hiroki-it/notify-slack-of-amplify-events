package detail

type Detail struct {
	AppId      *AppId      `json:"appId"`
	BranchName *BranchName `json:"branchName"`
	JobId      *JobId      `json:"jobId"`
	JobStatus  *JobStatus  `json:"jobStatus"`
}

// NewDetail コンストラクタ
func NewDetail() *Detail {

	return &Detail{}
}

// GetJobStatus JobStatusを返却します.
func (d *Detail) GetJobStatus() *JobStatus {
	return d.JobStatus
}

// IsSucceed ジョブステータスが成功かどうかを判定します
func (d *Detail) IsSucceed() bool {

	if d.JobStatus.Value() == "SUCCEED" {
		return true
	}

	return false
}
