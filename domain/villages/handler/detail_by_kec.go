package handler

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/boombaw/go-wilayah/domain/villages"
	"github.com/boombaw/go-wilayah/util"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

// DetailByDistrict struct
type DetailByDistrict struct {
	DBx *sqlx.DB
}

// Handle Detail Village By District handler
func (d *DetailByDistrict) Handle(c echo.Context) (err error) {
	ctx := c.Request().Context()

	if ctx == nil {
		ctx = context.Background()
	}

	var resp util.Response

	var kecid = c.Param("kecid")

	var village []villages.Village

	err = d.DBx.Select(&village, "SELECT * FROM villages WHERE district_id = $1", kecid)

	switch err {
	case nil:
		util.LogEntry(ctx).Info(fmt.Sprint("Success read villages with kecid ", kecid))
	case sql.ErrNoRows:
		util.LogEntry(ctx).WithField("error", err).Info(fmt.Sprint("data villages with kecid ", kecid, " not found"))
		return errors.New("data not found")
	default:
		log.Printf("error: %s\n", err)
		return
	}

	resp.Code = http.StatusOK
	resp.Message = http.StatusText(http.StatusOK)
	resp.Data = village

	return c.JSON(http.StatusOK, resp)
}

// NewDetailByDistrict func
func NewDetailByDistrict(db *sqlx.DB) *DetailByDistrict {
	return &DetailByDistrict{DBx: db}
}
