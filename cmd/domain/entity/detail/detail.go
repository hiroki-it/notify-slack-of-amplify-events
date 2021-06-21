package detail

/**
 * Detailを構成します．
 */
type Detail struct {
	AppId      string `json:"appId"`
	BranchName string `json:"branchName"`
	JobId      string `json:"jobId"`
	JobStatus  string `json:"jobStatus"`
}

/**
 * コンストラクタ
 * Detailを作成します．
 */
func NewDetail() *Detail {

	return &Detail{}
}

/**
 * JobStatusを返却します．
 */
func (d *Detail) GetJobStatus() string {
	return d.JobStatus
}

/**
 * ジョブステータスが成功かどうかを判定します．
 */
func (d *Detail) IsSucceed() bool {

	if d.JobStatus == "SUCCEED" {
		return true
	}

	return false
}
