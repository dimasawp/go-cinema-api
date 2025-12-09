package main

import (
	"github.com/dimasawp/go-cinema-api/config"
	"github.com/dimasawp/go-cinema-api/routes"
	"github.com/labstack/echo/v4"
)

func main() {
    config.ConnectDB()
    e := echo.New()
    routes.RegisterRoutes(e)
    e.Logger.Fatal(e.Start(":8080"))
}