package detail

import (
	"encoding/json"
	"errors"

	"github.com/go-playground/validator"
	"github.com/hiroki-it/notify-slack-of-amplify-events/cmd/presentation/validators"
)

type DetailValidator struct {
	*validators.Validator
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
func (v *DetailValidator) Validate() error {

	err := validator.New().Struct(v)

	var errorMessages = make(map[string]string)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Field() {
			case "appId":
				errorMessages["appId"] = v.StringValidation(err)
				errorMessages["appId"] = v.RequiredValidation(err)
			case "branchName":
				errorMessages["branchName"] = v.StringValidation(err)
				errorMessages["branchName"] = v.RequiredValidation(err)
			case "jobId":
				errorMessages["jobId"] = v.StringValidation(err)
				errorMessages["jobId"] = v.RequiredValidation(err)
			case "jobStatus":
				errorMessages["jobStatus"] = v.StringValidation(err)
				errorMessages["jobStatus"] = v.RequiredValidation(err)
			}
		}
	}

	byteJson, _ := json.Marshal(errorMessages)

	// Lambdaではエラーをerrorインターフェースで扱う必要があるため．独自エラーとして定義します．
	return errors.New(string(byteJson))
}
