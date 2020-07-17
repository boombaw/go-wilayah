package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/boombaw/go-wilayah/domain/regencies"
	"github.com/boombaw/go-wilayah/util"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

// Regencies struct
type Regencies struct {
	DBx *sqlx.DB
}

// Handle for Regencies handler
func (r *Regencies) Handle(c echo.Context) (err error) {
	ctx := c.Request().Context()

	if ctx == nil {
		ctx = context.Background()
	}

	var resp util.Response

	var reg []regencies.Regency
	var regen regencies.Regency
	fmt.Printf("%#v\n", regen)
	err = r.DBx.Select(&reg, "SELECT id,name,province_id FROM regencies")
	if err != nil {
		util.LogEntry(ctx).WithField("error", err).Error("Error while reading table")
		return
	}

	resp.Code = http.StatusOK
	resp.Message = http.StatusText(http.StatusOK)
	resp.Data = reg

	util.LogEntry(ctx).WithField("response", resp).Info("Response list regencies")

	return c.JSON(http.StatusOK, resp)
}

// NewRegencies func
func NewRegencies(db *sqlx.DB) *Regencies {
	return &Regencies{DBx: db}
}
