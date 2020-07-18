package handler

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/boombaw/go-wilayah/domain/regencies"
	"github.com/boombaw/go-wilayah/util"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

// DetailRegenciesByProv struct
type DetailRegenciesByProv struct {
	DBx *sqlx.DB
}

// Handle Detail Regencies By Provinces handler
func (d *DetailRegenciesByProv) Handle(c echo.Context) (err error) {
	ctx := c.Request().Context()

	if ctx == nil {
		ctx = context.Background()
	}

	var resp util.Response

	var provid = c.Param("provid")

	var reg []regencies.Regency

	err = d.DBx.Select(&reg, "SELECT * FROM regencies WHERE province_id = $1", provid)

	switch err {
	case nil:
		util.LogEntry(ctx).Info(fmt.Sprint("Success read regency with provid ", provid))
	case sql.ErrNoRows:
		util.LogEntry(ctx).WithField("error", err).Info(fmt.Sprint("data regency with provid ", provid, " not found"))
		return errors.New("data not found")
	default:
		log.Printf("error: %s\n", err)
		return
	}

	resp.Code = http.StatusOK
	resp.Message = http.StatusText(http.StatusOK)
	resp.Data = reg

	return c.JSON(http.StatusOK, resp)
}

// NewDetailRegenciesByProv func
func NewDetailRegenciesByProv(db *sqlx.DB) *DetailRegenciesByProv {
	return &DetailRegenciesByProv{DBx: db}
}
