package detail

type JobStatusType string

const (
	SUCCEED JobStatusType = "SUCCEED"
	FAILED  JobStatusType = "FAILED"
)

// NewJobStatusType コンストラクタ
func NewJobStatusType(jobStatusType string) JobStatusType {
	return JobStatusType(jobStatusType)
}

// String 区分値を返却します．
func (js JobStatusType) String() string {

	switch js {
	case SUCCEED:
		return "成功"
	case FAILED:
		return "失敗"
	default:
		return "不明のステータス"
	}
}
