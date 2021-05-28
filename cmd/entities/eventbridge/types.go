package eventbridge

/**
 * EventDetailインターフェースを構成します．
 */
type EventDetailInterface interface {
}

/**
 * EventDetailインターフェースの実装を構成します．
 */
type EventDetail struct {
	EventDetailInterface
	AppId      string `json:"appId"`
	BranchName string `json:"branchName"`
	JobId      string `json:"jobId"`
	JobStatus  string `json:"jobStatus"`
}

/**
 * ジョブステータスを構成します．
 */
type JobStatusColor struct {
	Status string
}
