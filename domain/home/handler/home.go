package handler

import (
	"net/http"

	"github.com/boombaw/go-wilayah/util"
	"github.com/labstack/echo/v4"
	"golang.org/x/net/context"
)

// Home struct
type Home struct{}

// Handle handler for home
func (*Home) Handle(c echo.Context) (err error) {
	ctx := c.Request().Context()

	if ctx == nil {
		ctx = context.Background()
	}

	var resp util.Response

	resp.Code = http.StatusOK
	resp.Message = http.StatusText(http.StatusOK)
	resp.Data = map[string]string{
		"provinsi":         c.Request().URL.Host + "/api/v1/wilayah/provinsi",
		"provinsi_detail":  c.Request().URL.Host + "/api/v1/wilayah/provinsi/:id",
		"kabupaten":        c.Request().URL.Host + "/api/v1/wilayah/kabupaten",
		"kabupaten_detail": c.Request().URL.Host + "/api/v1/wilayah/kabupaten/:id",
		"kecamatan":        c.Request().URL.Host + "/api/v1/wilayah/kecamatan",
		"kecamatan_detail": c.Request().URL.Host + "/api/v1/wilayah/kecamatan/:id",
		"kelurahan":        c.Request().URL.Host + "/api/v1/wilayah/kelurahan",
		"kelurahan_detail": c.Request().URL.Host + "/api/v1/wilayah/kelurahan/:id",
	}
	return c.JSON(http.StatusOK, resp)
}

// NewHome func
func NewHome() *Home {
	return &Home{}
}
