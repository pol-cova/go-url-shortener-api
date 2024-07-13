package main

import (
	"github.com/labstack/echo/v4"
	"github.com/pol-cova/go-url-shortener-api/db"
	"github.com/pol-cova/go-url-shortener-api/routes"
)

func main() {
	app := echo.New()
	db.InitDB()
	routes.Router(app)
	err := app.Start(":8080")
	if err != nil {
		panic("could not start the server")
	}
}
