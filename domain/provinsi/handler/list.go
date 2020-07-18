package handler

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
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

	switch err {
	case nil:
		util.LogEntry(ctx).Info(fmt.Sprint("Success read provinces "))
	case sql.ErrNoRows:
		util.LogEntry(ctx).WithField("error", err).Info(fmt.Sprint("data provinces not found"))
		return errors.New("data not found")
	default:
		log.Printf("error: %s\n", err)
		return
	}

	resp.Code = http.StatusOK
	resp.Message = http.StatusText(http.StatusOK)
	resp.Data = prov

	return c.JSON(http.StatusOK, resp)
}

// NewProvince func
func NewProvince(db *sqlx.DB) *Province {
	return &Province{DBx: db}
}
