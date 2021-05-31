package slack

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
func (sc *JobStatusColor) PrintStatus() string {

	if sc.status == "SUCCEED" {
		return "成功"
	}

	return "失敗"
}

/**
 * ジョブ状態を表現する色を返却します．
 */
func (sc *JobStatusColor) PrintColor() string {

	if sc.status == "SUCCEED" {
		return "#00FF00"
	}

	return "#ff0000"
}
