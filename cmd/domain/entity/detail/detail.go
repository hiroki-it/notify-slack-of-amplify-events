package detail

type Detail struct {
	AppId      string `json:"appId"`
	BranchName string `json:"branchName"`
	JobId      string `json:"jobId"`
	JobStatus  string `json:"jobStatus"`
}

// NewDetail コンストラクタ
func NewDetail() *Detail {

	return &Detail{}
}

// GetJobStatus JobStatusを返却します.
func (d *Detail) GetJobStatus() string {
	return d.JobStatus
}

// IsSucceed ジョブステータスが成功かどうかを判定します
func (d *Detail) IsSucceed() bool {

	if d.JobStatus == "SUCCEED" {
		return true
	}

	return false
}
