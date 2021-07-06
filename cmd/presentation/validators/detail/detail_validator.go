package detail

import (
	"github.com/go-playground/validator"
)

type DetailValidator struct {
	AppId         string `json:"appId" validate:"required"`
	BranchName    string `json:"branchName" validate:"required"`
	JobId         string `json:"jobId" validate:"required"`
	JobStatusType string `json:"jobStatus" validate:"required"`
}

// NewDetailValidator コンストラクタ
func NewDetailValidator() *DetailValidator {

	return &DetailValidator{}
}

// Validate バリデーションを実行します．
func (f *DetailValidator) Validate() error {
	v := validator.New()
	return v.Struct(f)
}
