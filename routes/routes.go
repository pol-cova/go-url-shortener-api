package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/pol-cova/go-url-shortener-api/handlers"
	"github.com/pol-cova/go-url-shortener-api/middlewares"
)

func Router(app *echo.Echo) {
	// Auth routes
	auth := app.Group("/auth")
	auth.POST("/signup", handlers.Signup)
	auth.POST("/login", handlers.Login)

	// Url routes
	app.POST("/short", handlers.ShortUrl)
	app.GET("/:key", handlers.RedirectUrl)

	// User routes
	user := app.Group("/user")
	user.Use(middlewares.AuthMiddleware)
	user.GET("/home", handlers.Home)
	user.GET("/me", handlers.Profile)
	user.GET("/delete", handlers.DeleteAccount)
}
