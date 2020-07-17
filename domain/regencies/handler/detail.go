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

// DetailRegencies struct
type DetailRegencies struct {
	DBx *sqlx.DB
}

// Handle Detail Regencies handler
func (d *DetailRegencies) Handle(c echo.Context) (err error) {
	ctx := c.Request().Context()

	if ctx == nil {
		ctx = context.Background()
	}

	var resp util.Response

	var id = c.Param("id")

	var reg regencies.Regency

	err = d.DBx.Get(&reg, "SELECT * FROM regencies WHERE id = $1", id)

	switch err {
	case nil:
		util.LogEntry(ctx).WithField("info", resp).Info(fmt.Sprint("Success read regency with id", id))
	case sql.ErrNoRows:
		util.LogEntry(ctx).WithField("error", resp).Info(fmt.Sprint("data regency with id ", id, " not found"))
		return errors.New("data not found")
	default:
		log.Printf("error: %s\n", err)
		return
	}

	resp.Code = http.StatusOK
	resp.Message = http.StatusText(http.StatusOK)
	resp.Data = []regencies.Regency{reg}

	return c.JSON(http.StatusOK, resp)
}

// NewDetailRegencies func
func NewDetailRegencies(db *sqlx.DB) *DetailRegencies {
	return &DetailRegencies{DBx: db}
}
