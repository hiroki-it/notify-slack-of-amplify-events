package slack

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
 * ジョブステータスを表現する文言を返却します．
 */
func (js *JobStatus) PrintStatusWord() string {

	if js.name == "SUCCEED" {
		return "成功"
	}

	return "失敗"
}

/**
 * ジョブステータスを表現する色を返却します．
 */
func (js *JobStatus) PrintStatusColorCode() string {

	if js.name == "SUCCEED" {
		return "#00FF00"
	}

	return "#ff0000"
}
