package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/pol-cova/go-url-shortener-api/handlers"
)

func Router(app *echo.Echo) {
	app.POST("/signup", handlers.Signup)
}
