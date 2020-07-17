package util

import (
	"github.com/go-playground/locales/id"
	ut "github.com/go-playground/universal-translator"
	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"
	idTrans "gopkg.in/go-playground/validator.v9/translations/id"
)

// CustomValidator validation that handle validation
type CustomValidator struct {
	Validator *validator.Validate
}

var translator ut.Translator

func setValidator(e *echo.Echo) {
	id := id.New()
	uni := ut.New(id, id)

	translator, _ = uni.GetTranslator("id")
	validator := validator.New()
	_ = idTrans.RegisterDefaultTranslations(validator, translator)
	e.Validator = &CustomValidator{Validator: validator}
}

// Validate Struct
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}
