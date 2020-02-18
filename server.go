package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	// logger
	e.Use(middleware.Logger())
	// Stream recovery
	e.Use(middleware.Recover())

	//CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
	e.GET("/", func(c echo.Context) error {
		description := "Simple website scraping API"
		return c.String(http.StatusOK, description)
	})
	// Live reload
	//e.Use(livereload.LiveReload())
	// start the server
	e.Logger.Fatal(e.Start(":8000"))
}
