package detail

import (
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/domain/core"
)

type BranchName struct {
	core.ValueObject
	name string
}

// NewBranchName コンストラクタ
func NewBranchName(name string) *BranchName {

	return &BranchName{
		name: name,
	}
}

// Name 属性を返却します．
func (bn *BranchName) Name() string {
	return bn.name
}

// Equals 等価性を検証します．
func (bn *BranchName) Equals(target *BranchName) bool {
	return bn.name == target.Name()
}
