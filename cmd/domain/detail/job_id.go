package detail

type JobId struct {
	value string
}

// NewJobId コンストラクタ
func NewJobId(value string) *JobId {

	return &JobId{
		value: value,
	}
}

// Value 属性を返却します．
func (ji JobId) Value() string {
	return ji.value
}

// Equals 等価性を検証します．
func (ji JobId) Equals(target *JobId) bool {
	return ji.value == target.Value()
}
