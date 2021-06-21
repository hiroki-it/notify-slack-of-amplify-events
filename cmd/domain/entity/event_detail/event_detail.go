package event

/**
 * EventDetailを構成します．
 */
type EventDetail struct {
	AppId      string `json:"appId"`
	BranchName string `json:"branchName"`
	JobId      string `json:"jobId"`
	JobStatus  string `json:"jobStatus"`
}

/**
 * コンストラクタ
 * EventDetailを作成します．
 */
func NewEventDetail() *EventDetail {

	return &EventDetail{}
}

/**
 * JobStatusを返却します．
 */
func (eventDetail *EventDetail) GetJobStatus() string {
	return eventDetail.JobStatus
}

/**
 * ジョブステータスが成功かどうかを判定します．
 */
func (eventDetail *EventDetail) IsSucceed() bool {

	if eventDetail.JobStatus == "SUCCEED" {
		return true
	}

	return false
}
