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

// DetailVillage struct
type DetailVillage struct {
	DBx *sqlx.DB
}

// Handle Detail Village handler
func (d *DetailVillage) Handle(c echo.Context) (err error) {
	ctx := c.Request().Context()

	if ctx == nil {
		ctx = context.Background()
	}

	var resp util.Response

	var id = c.Param("id")

	var village villages.Village

	err = d.DBx.Get(&village, "SELECT * FROM villages WHERE id = $1", id)

	switch err {
	case nil:
		util.LogEntry(ctx).WithField("info", resp).Info(fmt.Sprint("Success read villages with id", id))
	case sql.ErrNoRows:
		util.LogEntry(ctx).WithField("error", resp).Info(fmt.Sprint("data villages with id ", id, " not found"))
		return errors.New("data not found")
	default:
		log.Printf("error: %s\n", err)
		return
	}

	resp.Code = http.StatusOK
	resp.Message = http.StatusText(http.StatusOK)
	resp.Data = []villages.Village{village}

	return c.JSON(http.StatusOK, resp)
}

// NewDetailVillage func
func NewDetailVillage(db *sqlx.DB) *DetailVillage {
	return &DetailVillage{DBx: db}
}
