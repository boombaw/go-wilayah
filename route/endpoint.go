package route

import (
	"github.com/boombaw/go-wilayah/database"
	districts "github.com/boombaw/go-wilayah/domain/districts/handler"
	home "github.com/boombaw/go-wilayah/domain/home/handler"
	provinsi "github.com/boombaw/go-wilayah/domain/provinsi/handler"
	regencies "github.com/boombaw/go-wilayah/domain/regencies/handler"
	villages "github.com/boombaw/go-wilayah/domain/villages/handler"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

// Handler endpoint to use it later
type Handler interface {
	Handle(c echo.Context) (err error)
}

var db *sqlx.DB = database.Init()

var endpoint = map[string]Handler{
	// Home
	"home": home.NewHome(),
	// Provinsi
	"provinsi":        provinsi.NewProvince(db),
	"provinsi_detail": provinsi.NewDetailProvince(db),
	// Kabupaten
	"kabupaten":             regencies.NewRegencies(db),
	"kabupaten_detail":      regencies.NewDetailRegencies(db),
	"kabupaten_by_provinsi": regencies.NewDetailRegenciesByProv(db),
	// Kecamatan
	"kecamatan":        districts.NewDistricts(db),
	"kecamatan_detail": districts.NewDetailDistricts(db),
	"kecamatan_by_kab": districts.NewDetailByRegency(db),
	// Kelurahan
	"kelurahan":        villages.NewVillages(db),
	"kelurahan_detail": villages.NewDetailVillage(db),
	"kelurahan_by_kec": villages.NewDetailByDistrict(db),
}
