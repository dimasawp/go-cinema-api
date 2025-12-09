package controllers

import (
	"net/http"

	"github.com/dimasawp/go-cinema-api/config"
	"github.com/dimasawp/go-cinema-api/models"
	"github.com/dimasawp/go-cinema-api/utils"
	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
    type LoginInput struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }

    input := new(LoginInput)
    if err := c.Bind(input); err != nil {
        return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
    }

    var user models.User
    err := config.DB.Get(&user, "SELECT * FROM users WHERE email=$1", input.Email)
    if err != nil {
        return c.JSON(http.StatusUnauthorized, echo.Map{"error": "invalid credentials"})
    }

    if !utils.CheckPasswordHash(input.Password, user.PasswordHash) {
        return c.JSON(http.StatusUnauthorized, echo.Map{"error": "invalid credentials"})
    }

    token, _ := utils.GenerateJWT(user.ID)
    return c.JSON(http.StatusOK, echo.Map{
        "token": token,
        "user":  user.FullName,
    })
}
