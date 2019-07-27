package router

import (
	"net/http"

	"github.com/gorilla/sessions"

	"github.com/labstack/echo-contrib/session"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/unipota/tsundoku/server/model"
)

type whoAmIResponse struct {
	Logined    bool   `json:"logined"`
	ScreenName string `json:"screenName,omitempty"`
	IconURL    string `json:"iconUrl,omitempty"`
}

func GetWhoAmIHandler(c echo.Context) error {
	deviceID := c.Get("deviceID").(uuid.UUID)

	user, err := model.GetUserByDeviceUUID(deviceID)

	if err != nil {
		if model.IsErrRecordNotFound(err) {
			return c.JSON(http.StatusOK, whoAmIResponse{false, "", ""})
		}
		return c.JSON(http.StatusInternalServerError, H{"Error with get user"})
	}

	return c.JSON(http.StatusOK, whoAmIResponse{true, user.Name, user.IconURL})
}

func PostLogoutHandler(c echo.Context) error {
	sess, _ := session.Get("session", c)

	sess.Options = &sessions.Options{MaxAge: -1, Path: "/"}
	if err := sess.Save(c.Request(), c.Response()); err != nil {
		return c.JSON(http.StatusInternalServerError, H{"Something error while save session"})
	}
	return c.NoContent(http.StatusOK)
}
