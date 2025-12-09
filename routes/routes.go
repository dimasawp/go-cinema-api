package routes

import (
	"github.com/dimasawp/go-cinema-api/controllers"
	"github.com/dimasawp/go-cinema-api/middlewares"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
    e.POST("/login", controllers.Login)

	showtime := e.Group("/showtimes")
    showtime.Use(middlewares.JWTMiddleware)
    showtime.GET("", controllers.GetShowtimes)
    showtime.POST("", controllers.CreateShowtime)
    showtime.PUT("/:id", controllers.UpdateShowtime)
    showtime.DELETE("/:id", controllers.DeleteShowtime)
}
