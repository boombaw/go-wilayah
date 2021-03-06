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

// Villages struct
type Villages struct {
	DBx *sqlx.DB
}

// Handle for Villages handler
func (v *Villages) Handle(c echo.Context) (err error) {
	ctx := c.Request().Context()

	if ctx == nil {
		ctx = context.Background()
	}

	var resp util.Response

	var village []villages.Village

	err = v.DBx.Select(&village, "SELECT * FROM villages")
	switch err {
	case nil:
		util.LogEntry(ctx).Info(fmt.Sprint("Success read villages"))
	case sql.ErrNoRows:
		util.LogEntry(ctx).WithField("error", err).Info(fmt.Sprint("data villages not found"))
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

// NewVillages func
func NewVillages(db *sqlx.DB) *Villages {
	return &Villages{DBx: db}
}
