package detail

type JobStatusType int

const (
	SUCCEED JobStatusType = iota + 1
	FAILED
)

// NewJobStatusType コンストラクタ
func NewJobStatusType(jobStatusType int) JobStatusType {
	return JobStatusType(jobStatusType)
}

// Status 属性を返却します．
func (js JobStatusType) String() string {

	switch js {
	case SUCCEED:
		return "SUCCEED"
	case FAILED:
		return "FAILED"
	default:
		return "Unknown"
	}
}
