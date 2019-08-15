package router

import (
	"github.com/labstack/echo/v4"
)

func ScreenshotHandler(c echo.Context) error {
	return c.File("/tsundoku/ogp/screenshots/screenshot.png")
}
