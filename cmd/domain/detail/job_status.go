package detail

type JobStatus string

// NewJobStatus コンストラクタ
func NewJobStatus(value string) JobStatus {

	return JobStatus(value)
}

// Value 属性を返却します．
func (js JobStatus) Value() string {
	return string(js)
}

// Equals 等価性を検証します．
func (js JobStatus) Equals(target *JobStatus) bool {
	return js.Value() == target.Value()
}
