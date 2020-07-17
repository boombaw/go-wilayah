package route

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/boombaw/go-wilayah/util"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.elastic.co/apm/module/apmechov4"
)

// Route for mapping from json file
type Route struct {
	Path       string   `json:"path"`
	Method     string   `json:"method"`
	Module     string   `json:"module"`
	Endpoint   string   `json:"endpoint_filter"`
	Middleware []string `json:"middleware"`
}

// Init gateway router
func Init() *echo.Echo {
	routes := loadRoutes("routes.json")

	e := echo.New()

	// Set Bundle Middleware
	e.Use(middleware.RequestID())
	e.Use(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	e.Use(middleware.Gzip())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:  []string{"*"},
		AllowHeaders:  []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAcceptEncoding, echo.HeaderAuthorization, echo.HeaderContentLength, echo.HeaderAccessControlAllowOrigin, echo.HeaderContentDisposition, echo.HeaderAccessControlAllowHeaders},
		ExposeHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAcceptEncoding, echo.HeaderAuthorization, echo.HeaderContentLength, echo.HeaderAccessControlAllowOrigin, echo.HeaderContentDisposition, echo.HeaderAccessControlAllowHeaders},
	}))
	e.Use(apmechov4.Middleware())
	prefix := "/api/v1/wilayah"
	// TODO: add middleware to routes.json and mapping
	for _, route := range routes {
		e.Add(route.Method, prefix+route.Path, endpoint[route.Endpoint].Handle, chainMiddleware(route.Middleware)...)
	}
	util.CustomErrorHandler(e)
	return e
}

func loadRoutes(filePath string) []Route {
	var routes []Route
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Failed to load file: %v", err)
	}

	if err := json.Unmarshal(file, &routes); err != nil {
		log.Fatalf("Failed to marshal file: %v", err)
	}

	return routes
}

func chainMiddleware(tags []string) []echo.MiddlewareFunc {
	mwHandlers := []echo.MiddlewareFunc{}

	for _, v := range tags {
		mwHandlers = append(mwHandlers, middlewareHandler[v])
	}

	return mwHandlers
}
