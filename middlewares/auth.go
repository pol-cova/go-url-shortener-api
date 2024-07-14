package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/pol-cova/go-url-shortener-api/utils"
	"net/http"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")
		if token == "" {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "unauthorized", "message": "token required"})
		}
		userId, err := utils.ValidateToken(token)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "unauthorized", "message": "invalid token"})
		}
		c.Set("userId", userId)

		return next(c)
	}
}
