package validators

import (
	"fmt"

	"github.com/go-playground/validator"
)

type Validator struct {
}

// StringValidation 文字列型を検証します．
func (v *Validator) StringValidation(err validator.FieldError) string {

	switch err.Tag() {
	// 必須かどうかを検証します．
	case "required":
		return fmt.Sprintf("%s は必須です", err.Field())
	}

	return fmt.Sprintf("%s は文字列のみ有効です", err.Field())
}
