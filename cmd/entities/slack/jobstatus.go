package slack

/**
 * コンストラクタ
 * JobStatusを作成します．
 */
func NewJobStatus(status string) *JobStatus {
	return &JobStatus{
		Status: status,
	}
}

/**
 * ジョブ状態を表現するメッセージを返却します．
 */
func (jobStatus *JobStatus) PrintJobStatus() (string, string) {

	if jobStatus.Status == "SUCCEED" {
		return "成功", "#00FF00"
	}

	return "失敗", "#ff0000"
}
