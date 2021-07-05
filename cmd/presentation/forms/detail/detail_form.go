package detail

import (
	"github.com/go-playground/validator"
)

type DetailForm struct {
	AppId         string `json:"appId" validate:"required"`
	BranchName    string `json:"branchName" validate:"required"`
	JobId         string `json:"jobId" validate:"required"`
	JobStatusType string `json:"jobStatusType" validate:"required"`
}

// NewDetailForm コンストラクタ
func NewDetailForm() *DetailForm {

	return &DetailForm{}
}

// Validate バリデーションを実行します．
func (f *DetailForm) Validate() error {
	v := validator.New()
	return v.Struct(f)
}
