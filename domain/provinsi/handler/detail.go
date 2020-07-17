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

type DetailProvince struct {
	DBx *sqlx.DB
}

// Handle Detail Provinsi handler
func (d *DetailProvince) Handle(c echo.Context) (err error) {
	ctx := c.Request().Context()

	if ctx == nil {
		ctx = context.Background()
	}

	var resp util.Response

	var id = c.Param("id")

	var p provinsi.Provinsi

	err = d.DBx.Get(&p, "SELECT * FROM provinces WHERE id = $1", id)

	switch err {
	case nil:
		util.LogEntry(ctx).WithField("info", resp).Info(fmt.Sprint("Success read provinces with id", id))
	case sql.ErrNoRows:
		util.LogEntry(ctx).WithField("error", resp).Info(fmt.Sprint("data provinces with id ", id, " not found"))
		return errors.New("data not found")
	default:
		log.Printf("error: %s\n", err)
		return
	}

	resp.Code = http.StatusOK
	resp.Message = http.StatusText(http.StatusOK)
	resp.Data = []provinsi.Provinsi{p}

	return c.JSON(http.StatusOK, resp)
}

func NewDetailProvince(db *sqlx.DB) *DetailProvince {
	return &DetailProvince{DBx: db}
}
