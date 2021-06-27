package detail

type JobId string

// NewJobId コンストラクタ
func NewJobId(value string) JobId {

	return JobId(value)
}

// Value 属性を返却します．
func (ji JobId) Value() string {
	return string(ji)
}

// Equals 等価性を検証します．
func (ji JobId) Equals(target *JobId) bool {
	return ji.Value() == target.Value()
}
