package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:  "./static",
		HTML5: true,
	}))
	e.GET("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "HelloWorld")
	})

	e.GET("/*", func(c echo.Context) error {
		return c.File("./dist/index.html")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	e.Start(":" + port)
}
