package eventbridge

/**
 * ジョブステータスを構成します．
 */
type JobStatus struct {
	name string
}

/**
 * コンストラクタ
 * JobStatusを作成します．
 */
func NewJobStatus(name string) *JobStatus {
	return &JobStatus{
		name: name,
	}
}

/**
 * ジョブステータスが成功かどうかを判定します．
 */
func (js *JobStatus) IsSucceed() bool {

	if js.name == "SUCCEED" {
		return true
	}

	return false
}
