package main

import (
	"github.com/labstack/echo/v4"
	"github.com/pol-cova/go-url-shortener-api/db"
	"github.com/pol-cova/go-url-shortener-api/routes"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	app := echo.New()
	db.InitDB()
	routes.Router(app)
	err := app.Start(":" + port)
	if err != nil {
		panic("could not start the server")
	}
}
