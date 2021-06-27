package detail

type BranchName string

// NewBranchName コンストラクタ
func NewBranchName(value string) BranchName {

	return BranchName(value)
}

// Value 属性を返却します．
func (bn BranchName) Value() string {
	return string(bn)
}

// Equals 等価性を検証します．
func (bn BranchName) Equals(target *BranchName) bool {
	return bn.Value() == target.Value()
}
