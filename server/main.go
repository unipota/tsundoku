package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "HelloWorld")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	e.Start(port)
}
