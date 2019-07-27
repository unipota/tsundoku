package router

import (
	"net/http"
	"time"

	"github.com/gorilla/sessions"

	"github.com/labstack/echo-contrib/session"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/unipota/tsundoku/server/model"
)

type whoAmIResponse struct {
	Logined    bool      `json:"logined"`
	ScreenName string    `json:"screenName,omitempty"`
	IconURL    string    `json:"iconUrl,omitempty"`
	CreatedAt  time.Time `json:"createdAt"`
}

func GetWhoAmIHandler(c echo.Context) error {
	deviceID := c.Get("deviceID").(uuid.UUID)

	user, err := model.GetUserByDeviceUUID(deviceID)

	if err != nil {
		if model.IsErrRecordNotFound(err) {
			device, err := model.GetDeviceByDeviceID(deviceID)
			if err != nil {
				c.Logger().Error(err)
				return c.JSON(http.StatusInternalServerError, H{"Error with get device"})
			}
			return c.JSON(http.StatusOK, whoAmIResponse{false, "", "", device.CreatedAt})
		}
		return c.JSON(http.StatusInternalServerError, H{"Error with get user"})
	}

	return c.JSON(http.StatusOK, whoAmIResponse{true, user.Name, user.IconURL, user.CreatedAt})
}

func PostLogoutHandler(c echo.Context) error {
	sess, _ := session.Get("session", c)

	sess.Options = &sessions.Options{MaxAge: -1, Path: "/"}
	if err := sess.Save(c.Request(), c.Response()); err != nil {
		return c.JSON(http.StatusInternalServerError, H{"Something error while save session"})
	}
	return c.NoContent(http.StatusOK)
}
