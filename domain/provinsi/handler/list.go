package handler

import (
	"context"
	"net/http"

	"github.com/boombaw/go-wilayah/domain/provinsi"
	"github.com/boombaw/go-wilayah/util"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

// Province struct
type Province struct {
	DBx *sqlx.DB
}

// Handle for province handler
func (p *Province) Handle(c echo.Context) (err error) {
	ctx := c.Request().Context()

	if ctx == nil {
		ctx = context.Background()
	}

	var resp util.Response

	var prov []provinsi.Provinsi

	err = p.DBx.Select(&prov, "SELECT * FROM provinces")
	if err != nil {
		util.LogEntry(ctx).WithField("error", err).Error("Error while reading table")
		return
	}

	resp.Code = http.StatusOK
	resp.Message = http.StatusText(http.StatusOK)
	resp.Data = prov

	util.LogEntry(ctx).WithField("response", resp).Info("Response list provinsi")

	return c.JSON(http.StatusOK, resp)
}

// NewProvince func
func NewProvince(db *sqlx.DB) *Province {
	return &Province{DBx: db}
}
