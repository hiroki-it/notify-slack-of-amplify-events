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
 * ジョブステータスを表現する文言を返却します．
 */
func (sc *JobStatusColor) PrintStatusWord() string {

	if sc.status == "SUCCEED" {
		return "成功"
	}

	return "失敗"
}

/**
 * ジョブステータスを表現する色を返却します．
 */
func (sc *JobStatusColor) PrintStatusColorCode() string {

	if sc.status == "SUCCEED" {
		return "#00FF00"
	}

	return "#ff0000"
}
