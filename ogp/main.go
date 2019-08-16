package main

import (
	"log"
	"os"

	"github.com/labstack/echo/v4"

	"github.com/unipota/tsundoku/ogp/router"
)

func main() {
	e := echo.New()
	e.GET("/ss", router.Screenshot)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3001"
	}
	log.Fatal(e.Start(":" + port))
}
