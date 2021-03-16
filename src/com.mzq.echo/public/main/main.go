package main

import (
	"com.mzq.echo/app/core"

	"com.mzq.echo/app/core/response"
	router "com.mzq.echo/routes"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := core.App()

	e.HTTPErrorHandler = response.MzqHTTPErrorHandler
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	router.Router(e)


	// Start server
	e.Logger.Fatal(e.Start(":1323"))

}

