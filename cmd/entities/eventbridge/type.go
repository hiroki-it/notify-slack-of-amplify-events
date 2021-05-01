package eventbridge

/**
 * EventのDetailを構成します．
 */
type EventDetail struct {
	AppId      string `json:"appId"`
	BranchName string `json:"branchName"`
	JobId      string `json:"jobId"`
	JobStatus  string `json:"jobStatus"`
}
