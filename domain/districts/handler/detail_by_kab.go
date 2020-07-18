package handler

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/boombaw/go-wilayah/domain/districts"
	"github.com/boombaw/go-wilayah/util"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type DetailByRegency struct {
	DBx *sqlx.DB
}

// Handle Detail District By Regency handler
func (d *DetailByRegency) Handle(c echo.Context) (err error) {
	ctx := c.Request().Context()

	if ctx == nil {
		ctx = context.Background()
	}

	var resp util.Response

	var kabid = c.Param("kabid")

	var dist []districts.District

	err = d.DBx.Select(&dist, "SELECT * FROM districts WHERE regency_id = $1", kabid)

	switch err {
	case nil:
		util.LogEntry(ctx).Info(fmt.Sprint("Success read districts with kabid ", kabid))
	case sql.ErrNoRows:
		util.LogEntry(ctx).WithField("error", err).Info(fmt.Sprint("data districts with kabid ", kabid, " not found"))
		return errors.New("data not found")
	default:
		log.Printf("error: %s\n", err)
		return
	}

	resp.Code = http.StatusOK
	resp.Message = http.StatusText(http.StatusOK)
	resp.Data = dist

	return c.JSON(http.StatusOK, resp)
}

func NewDetailByRegency(db *sqlx.DB) *DetailByRegency {
	return &DetailByRegency{DBx: db}
}
