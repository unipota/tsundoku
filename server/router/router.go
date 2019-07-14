package router

import (
	"github.com/google/uuid"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"net/http"
)

type H struct {
	Message string `json: "message"`
}

func IdentifyMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sess, _ := session.Get("session", c)

		if sess.IsNew {
			sess.Options = &sessions.Options{
				Path:     "/",
				MaxAge:   60 * 60 * 24 * 14,
				HttpOnly: true,
			}
			sess.Values["device_id"] = uuid.New().String()
		}

		deviceID, ok := sess.Values["device_id"]

		if !ok {
			return c.JSON(http.StatusInternalServerError, H{"Can't Parse deviceID"})
		}

		deviceUUID, err := uuid.Parse(deviceID.(string))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, H{"Can't Parse deviceID"})
		}

		c.Set("deviceID", deviceUUID)
		sess.Save(c.Request(), c.Response())
		return next(c)
	}
}

func Ping(c echo.Context) error {
	deviceUUID, ok := c.Get("deviceID").(uuid.UUID)
	if !ok {
		return c.JSON(http.StatusInternalServerError, H{"Can't Get deviceID"})
	}

	return c.String(http.StatusOK, "pong"+deviceUUID.String())
}
