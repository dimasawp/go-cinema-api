package routes

import (
	"github.com/dimasawp/go-cinema-api/controllers"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
    e.POST("/login", controllers.Login)
}
