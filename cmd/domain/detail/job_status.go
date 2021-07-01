package detail

type JobStatus struct {
	value string
}

// NewJobStatus コンストラクタ
func NewJobStatus(value string) *JobStatus {

	return &JobStatus{
		value: value,
	}
}

// Value 属性を返却します．
func (js JobStatus) Value() string {
	return js.value
}

// Equals 等価性を検証します．
func (js JobStatus) Equals(target *JobStatus) bool {
	return js.value == target.Value()
}
