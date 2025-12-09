package controllers

import (
	"net/http"

	"github.com/dimasawp/go-cinema-api/config"
	"github.com/dimasawp/go-cinema-api/models"
	"github.com/labstack/echo/v4"
)

func GetShowtimes(c echo.Context) error {
    var showtimes []models.Showtime
    err := config.DB.Select(&showtimes, "SELECT * FROM showtimes")
    if err != nil {
        return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
    }
    return c.JSON(http.StatusOK, showtimes)
}

func CreateShowtime(c echo.Context) error {
    var input models.Showtime
    if err := c.Bind(&input); err != nil {
        return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
    }
    input.Status = "active"

    query := `INSERT INTO showtimes (movie_id, auditorium_id, start_time, end_time, status)
              VALUES ($1, $2, $3, $4, $5) RETURNING id`
    err := config.DB.QueryRow(query, input.MovieID, input.AuditoriumID, input.StartTime, input.EndTime, input.Status).Scan(&input.ID)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
    }
    return c.JSON(http.StatusCreated, input)
}

func UpdateShowtime(c echo.Context) error {
    id := c.Param("id")
    var input models.Showtime
    if err := c.Bind(&input); err != nil {
        return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
    }
    query := `UPDATE showtimes SET movie_id=$1, auditorium_id=$2, start_time=$3, end_time=$4, status=$5 WHERE id=$6`
    _, err := config.DB.Exec(query, input.MovieID, input.AuditoriumID, input.StartTime, input.EndTime, input.Status, id)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
    }
    return c.JSON(http.StatusOK, input)
}

func DeleteShowtime(c echo.Context) error {
    id := c.Param("id")
    _, err := config.DB.Exec("DELETE FROM showtimes WHERE id=$1", id)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
    }
    return c.JSON(http.StatusOK, echo.Map{"message": "deleted"})
}
