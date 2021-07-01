package detail

type BranchName struct {
	value string
}

// NewBranchName コンストラクタ
func NewBranchName(value string) *BranchName {

	return &BranchName{
		value: value,
	}
}

// Value 属性を返却します．
func (bn BranchName) Value() string {
	return bn.value
}

// Equals 等価性を検証します．
func (bn BranchName) Equals(target *BranchName) bool {
	return bn.value == target.Value()
}
