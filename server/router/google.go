package router

import (
	"context"
	"errors"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/unipota/tsundoku/server/model"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"

	"golang.org/x/oauth2"
	v2 "google.golang.org/api/oauth2/v2"
)

const (
	authorizeEndpoint = "https://accounts.google.com/o/oauth2/v2/auth"
	tokenEndpoint     = "https://www.googleapis.com/oauth2/v4/token"
)

var (
	GOOGLE_CLIENT_ID     = os.Getenv("GOOGLE_CLIENT_ID")
	GOOGLE_CLIENT_SECRET = os.Getenv("GOOGLE_CLIENT_SECRET")
)

type googleCallbackRequest struct {
	Code  string `form:"code"`
	State string `form:"state"`
}

func GetGoogleAuthHandler(c echo.Context) error {
	config := getGoogleConnect()

	state := genState()
	sess, _ := session.Get("session", c)
	sess.Values["state"] = state

	sess.Save(c.Request(), c.Response())

	url := config.AuthCodeURL(state)

	return c.Redirect(http.StatusFound, url)
}

func GetGoogleCallbackHandler(c echo.Context) error {
	req := &googleCallbackRequest{}
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusInternalServerError, H{"Bind error"})
	}

	sess, _ := session.Get("session", c)
	state := sess.Values["state"].(string)

	if state != req.State {
		return c.JSON(http.StatusBadRequest, H{"Invalid state"})
	}

	config := getGoogleConnect()

	context := context.Background()

	tok, err := config.Exchange(context, req.Code)
	if err != nil {
		panic(err)
	}

	if tok.Valid() == false {
		panic(errors.New("vaild token"))
	}

	service, _ := v2.New(config.Client(context, tok))
	info, err := service.Userinfo.V2.Me.Get().Do()
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, H{"get user info error"})
	}

	deviceID := c.Get("deviceID").(uuid.UUID)
	if err := model.ModifyUserAccount(deviceID, "google", info.Id, info.Name, info.Picture); err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, H{"Error with modify user"})
	}
	return c.Redirect(http.StatusFound, "/")
}

func getGoogleConnect() *oauth2.Config {
	config := &oauth2.Config{
		ClientID:     GOOGLE_CLIENT_ID,
		ClientSecret: GOOGLE_CLIENT_SECRET,
		Endpoint: oauth2.Endpoint{
			AuthURL:  authorizeEndpoint,
			TokenURL: tokenEndpoint,
		},
		Scopes:      []string{"openid", "profile"},
		RedirectURL: BASE_URL + "/auth/google/callback",
	}

	return config
}
