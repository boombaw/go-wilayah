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

// Districts struct
type Districts struct {
	DBx *sqlx.DB
}

// Handle for Districts handler
func (p *Districts) Handle(c echo.Context) (err error) {
	ctx := c.Request().Context()

	if ctx == nil {
		ctx = context.Background()
	}

	var resp util.Response

	var dist []districts.District

	err = p.DBx.Select(&dist, "SELECT * FROM districts")

	switch err {
	case nil:
		util.LogEntry(ctx).Info(fmt.Sprint("Success read districts"))
	case sql.ErrNoRows:
		util.LogEntry(ctx).WithField("error", err).Info(fmt.Sprint("data districts not found"))
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

// NewDistricts func
func NewDistricts(db *sqlx.DB) *Districts {
	return &Districts{DBx: db}
}
