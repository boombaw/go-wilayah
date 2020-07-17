package handler

import (
	"context"
	"net/http"

	"github.com/boombaw/go-wilayah/domain/villages"
	"github.com/boombaw/go-wilayah/util"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

// Villages struct
type Villages struct {
	DBx *sqlx.DB
}

// Handle for Villages handler
func (v *Villages) Handle(c echo.Context) (err error) {
	ctx := c.Request().Context()

	if ctx == nil {
		ctx = context.Background()
	}

	var resp util.Response

	var village []villages.Village

	err = v.DBx.Select(&village, "SELECT * FROM villages")
	if err != nil {
		util.LogEntry(ctx).WithField("error", err).Error("Error while reading table")
		return
	}

	resp.Code = http.StatusOK
	resp.Message = http.StatusText(http.StatusOK)
	resp.Data = village

	util.LogEntry(ctx).WithField("response", resp).Info("Response list villages")

	return c.JSON(http.StatusOK, resp)
}

// NewVillages func
func NewVillages(db *sqlx.DB) *Villages {
	return &Villages{DBx: db}
}
