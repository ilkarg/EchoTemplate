package main

import (
	middlewares "api/middlewares"
	routes "api/routes"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	middlewares.InitMiddlewares(e)
	routes.InitRoutes(e)
	e.Logger.Fatal(e.Start(":8001"))
}
