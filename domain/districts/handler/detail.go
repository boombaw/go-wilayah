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

type DetailDistricts struct {
	DBx *sqlx.DB
}

// Handle Detail District handler
func (d *DetailDistricts) Handle(c echo.Context) (err error) {
	ctx := c.Request().Context()

	if ctx == nil {
		ctx = context.Background()
	}

	var resp util.Response

	var id = c.Param("id")

	var dist districts.District

	err = d.DBx.Get(&dist, "SELECT * FROM districts WHERE id = $1", id)

	switch err {
	case nil:
		util.LogEntry(ctx).WithField("info", resp).Info(fmt.Sprint("Success read districts with id", id))
	case sql.ErrNoRows:
		util.LogEntry(ctx).WithField("error", resp).Info(fmt.Sprint("data districts with id ", id, " not found"))
		return errors.New("data not found")
	default:
		log.Printf("error: %s\n", err)
		return
	}

	resp.Code = http.StatusOK
	resp.Message = http.StatusText(http.StatusOK)
	resp.Data = []districts.District{dist}

	return c.JSON(http.StatusOK, resp)
}

func NewDetailDistricts(db *sqlx.DB) *DetailDistricts {
	return &DetailDistricts{DBx: db}
}
