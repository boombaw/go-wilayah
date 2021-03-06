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

	err = r.DBx.Select(&reg, "SELECT * FROM regencies")
	switch err {
	case nil:
		util.LogEntry(ctx).Info(fmt.Sprint("Success read regencies "))
	case sql.ErrNoRows:
		util.LogEntry(ctx).WithField("error", err).Info(fmt.Sprint("data regency not found"))
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

// NewRegencies func
func NewRegencies(db *sqlx.DB) *Regencies {
	return &Regencies{DBx: db}
}
