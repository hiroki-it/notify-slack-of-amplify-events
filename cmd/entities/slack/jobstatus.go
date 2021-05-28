package slack

/**
 * コンストラクタ
 * JobStatusColorを作成します．
 */
func NewJobStatusColor(status string) *JobStatusColor {
	return &JobStatusColor{
		Status: status,
	}
}

/**
 * ジョブ状態を表現するメッセージを返却します．
 */
func (jobStatusColor *JobStatusColor) PrintJobStatusColor() (string, string) {

	if jobStatusColor.Status == "SUCCEED" {
		return "成功", "#00FF00"
	}

	return "失敗", "#ff0000"
}
