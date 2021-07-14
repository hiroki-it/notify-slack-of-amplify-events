package validators

import (
	"fmt"

	"github.com/go-playground/validator"
)

type Validator struct {
}

// RequiredValidation 必須を検証します．
func (v *Validator) RequiredValidation(err validator.FieldError) string {
	return fmt.Sprintf("%s は必須です", err.Field())
}

// StringValidation 必須を検証します．
func (v *Validator) StringValidation(err validator.FieldError) string {
	return fmt.Sprintf("%s は文字列のみ有効です", err.Field())
}
