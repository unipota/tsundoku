package router

import (
	"net/http"

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
