package middlewares

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/dimasawp/go-cinema-api/utils"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        authHeader := c.Request().Header.Get("Authorization")
        if (authHeader == "") {
            return c.JSON(http.StatusUnauthorized, echo.Map{"error": "missing token"})
        }

        tokenStr := strings.Split(authHeader, "Bearer ")[1]
        token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
            return utils.JwtSecret, nil
        })

        if err != nil || !token.Valid {
            return c.JSON(http.StatusUnauthorized, echo.Map{"error": "invalid token"})
        }

        claims := token.Claims.(jwt.MapClaims)
        c.Set("user_id", int64(claims["user_id"].(float64)))

        return next(c)
    }
}
