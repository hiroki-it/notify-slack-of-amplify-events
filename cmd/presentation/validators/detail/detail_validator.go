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

	if err != nil {
		var errorMessages = map[string]string{}
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Field() {
			case "AppId":
				// NOTE: 視認性の観点から，バリデーションメッセージは一つだけ出力するようにします．
				errorMessages["appId"] = v.StringValidation(err)
			case "BranchName":
				errorMessages["branchName"] = v.StringValidation(err)
			case "JobId":
				errorMessages["jobId"] = v.StringValidation(err)
			case "JobStatusType":
				errorMessages["jobStatusType"] = v.StringValidation(err)
			}
		}
		byteJson, _ := json.Marshal(errorMessages)
		// Lambdaではエラーをerrorインターフェースで扱う必要があるため．独自エラーとして定義します．
		return errors.New(string(byteJson))
	}

	return err
}
