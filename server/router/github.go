package router

import (
	"context"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo-contrib/session"

	"github.com/google/uuid"
	"github.com/unipota/tsundoku/server/model"

	gh "github.com/google/go-github/github"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
	githuboauth "golang.org/x/oauth2/github"
)

// CallbackRequest コールバックリクエスト
type githubCallbackRequest struct {
	Code  string `form:"code"`
	State string `form:"state"`
}

var (
	GITHUB_CLIENT_ID     = os.Getenv("GITHUB_CLIENT_ID")
	GITHUB_CLIENT_SECRET = os.Getenv("GITHUB_CLIENT_SECRET")
)

// Get 認証する
func GetGitHubAuthHandler(c echo.Context) error {
	config := getGitHubConnect()

	state := genState()
	sess, _ := session.Get("session", c)
	sess.Values["state"] = state

	sess.Save(c.Request(), c.Response())

	url := config.AuthCodeURL(state)

	return c.Redirect(http.StatusFound, url)
}

// Get コールバックする
func GetGitHubCallbackHandler(c echo.Context) error {
	req := &githubCallbackRequest{}
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusInternalServerError, H{"Bind error"})
	}
	sess, _ := session.Get("session", c)
	state := sess.Values["state"].(string)

	if state != req.State {
		return c.JSON(http.StatusBadRequest, H{"Invalid state"})
	}

	config := getGitHubConnect()

	background := context.Background()
	tok, err := config.Exchange(oauth2.NoContext, req.Code)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, H{"get access token error"})
	}

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: tok.AccessToken},
	)
	tc := oauth2.NewClient(background, ts)

	client := gh.NewClient(tc)

	u, _, err := client.Users.Get(background, "")
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, H{"get user info error"})
	}

	deviceID := c.Get("deviceID").(uuid.UUID)
	if err := model.ModifyUserAccount(deviceID, "github", strconv.Itoa(int(*u.ID)), *u.Login, *u.AvatarURL); err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, H{"Error with modify user"})
	}
	return c.Redirect(http.StatusFound, "/")
}

func getGitHubConnect() *oauth2.Config {
	config := &oauth2.Config{
		ClientID:     GITHUB_CLIENT_ID,
		ClientSecret: GITHUB_CLIENT_SECRET,
		Endpoint:     githuboauth.Endpoint,
		Scopes:       []string{"user:email"},
		RedirectURL:  BASE_URL + "/auth/github/callback",
	}

	return config
}
