package midlewares

import "github.com/go-playground/validator/v10"

type CustomValidator struct {
	Validator *validator.Validate
}

// Validate : Validate Data
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}
