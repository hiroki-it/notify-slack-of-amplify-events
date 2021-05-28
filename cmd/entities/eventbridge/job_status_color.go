package eventbridge

/**
 * コンストラクタ
 * JobStatusColorを作成します．
 */
func NewJobStatusColor(status string) *JobStatusColor {
	return &JobStatusColor{
		status: status,
	}
}

/**
 * ジョブ状態を表現するステータスを返却します．
 */
func (jobStatusColor *JobStatusColor) PrintStatus() string {

	if jobStatusColor.status == "SUCCEED" {
		return "成功"
	}

	return "失敗"
}

/**
 * ジョブ状態を表現する色を返却します．
 */
func (jobStatusColor *JobStatusColor) PrintColor() string {

	if jobStatusColor.status == "SUCCEED" {
		return "#00FF00"
	}

	return "#ff0000"
}
