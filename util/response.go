package util

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

// Response struct
type Response struct {
	Code    int                    `json:"code"`
	Message string                 `json:"message,omitempty"`
	Data    interface{}            `json:"data,omitempty"`
	Errors  []string               `json:"errors,omitempty"`
	Header  map[string]interface{} `json:"-"`
}

// JSON render response as JSON
func (r *Response) JSON(c echo.Context) error {
	for k, v := range r.Header {
		c.Response().Header().Set(k, fmt.Sprintf("%v, %v", c.Response().Header().Get(k), v))
	}
	return c.JSON(r.Code, r)
}
