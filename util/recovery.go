package util

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"
)

// CustomErrorHandler for handling Echo Recovery
func CustomErrorHandler(c *echo.Echo) {
	setValidator(c)
	c.HTTPErrorHandler = func(err error, c echo.Context) {

		// Validation Error
		if errs, ok := err.(validator.ValidationErrors); ok {
			var message []string

			translated := errs.Translate(translator)
			for _, v := range translated {
				message = append(message, v)
			}
			resp := Response{
				Code:    http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
				Errors:  message,
			}
			_ = resp.JSON(c)
		}
		resp := Response{
			Code:    http.StatusInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
			Errors:  []string{err.Error()},
		}
		_ = resp.JSON(c)
		return
	}
}
