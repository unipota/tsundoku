package router

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"os"

	"github.com/google/uuid"

	"github.com/unipota/tsundoku/server/model"

	"github.com/labstack/echo-contrib/session"

	"github.com/labstack/echo/v4"

	"github.com/gomodule/oauth1/oauth"
)

// Account アカウント
type Account struct {
	ID              string `json:"id_str"`
	ScreenName      string `json:"screen_name"`
	ProfileImageURL string `json:"profile_image_url_https"`
	Email           string `json:"email"`
}

var (
	TEMP_CREDENTIAL_KEY     = os.Getenv("TEMP_CREDENTIAL_KEY")
	TWITTER_CONSUMER_SECRET = os.Getenv("TWITTER_CONSUMER_SECRET")
	CALLBACK_URL            = os.Getenv("CALLBACK_URL")
)

type twitterCallbackRequest struct {
	Token    string `query:"oauth_token"`
	Verifier string `query:"oauth_verifier"`
}

func GetTwitterAuthHandler(c echo.Context) error {
	config := getTwitterConnect()
	if CALLBACK_URL == "" {
		CALLBACK_URL = "http://127.0.0.1/auth/twitter/callback"
	}
	rt, err := config.RequestTemporaryCredentials(nil, CALLBACK_URL, nil)
	if err != nil {
		panic(err)
	}

	sess, err := session.Get("session", c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, H{"Session get error"})
	}

	sess.Values["request_token"] = rt.Token
	sess.Values["request_token_secret"] = rt.Secret

	if err := sess.Save(c.Request(), c.Response()); err != nil {
		return c.JSON(http.StatusInternalServerError, H{"Session save error"})
	}

	redirectUrl := config.AuthorizationURL(rt, nil)
	return c.Redirect(http.StatusFound, redirectUrl)
}

func GetTwitterCallbackHandler(c echo.Context) error {
	req := &twitterCallbackRequest{}
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusInternalServerError, H{"Bind error"})
	}

	sess, err := session.Get("session", c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, H{"Session get error"})
	}

	at, err := getAccessToken(
		&oauth.Credentials{
			Token:  sess.Values["request_token"].(string),
			Secret: sess.Values["request_token_secret"].(string),
		},
		req.Verifier,
	)

	if err != nil {
		panic(err)
	}

	sess.Values["oauth_secret"] = at.Secret
	sess.Values["oauth_token"] = at.Token

	if err := sess.Save(c.Request(), c.Response()); err != nil {
		return c.JSON(http.StatusInternalServerError, H{"Session save error"})
	}

	account := Account{}
	if err = getMe(at, &account); err != nil {
		panic(err)
	}

	deviceID := c.Get("deviceID").(uuid.UUID)

	if err := model.ModifyUserAccount(deviceID, "twitter", account.ID, account.ScreenName, account.ProfileImageURL); err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, H{"Error with modify user"})
	}

	return c.Redirect(http.StatusFound, "/")
}

func getTwitterConnect() *oauth.Client {
	return &oauth.Client{
		TemporaryCredentialRequestURI: "https://api.twitter.com/oauth/request_token",
		ResourceOwnerAuthorizationURI: "https://api.twitter.com/oauth/authorize",
		TokenRequestURI:               "https://api.twitter.com/oauth/access_token",
		Credentials: oauth.Credentials{
			Token:  TEMP_CREDENTIAL_KEY,
			Secret: TWITTER_CONSUMER_SECRET,
		},
	}
}

// getAccessToken アクセストークンを取得する
func getAccessToken(rt *oauth.Credentials, oauthVerifier string) (*oauth.Credentials, error) {
	oc := getTwitterConnect()
	at, _, err := oc.RequestToken(nil, rt, oauthVerifier)

	return at, err
}

// getMe 自身を取得する
func getMe(at *oauth.Credentials, user *Account) error {
	oc := getTwitterConnect()

	v := url.Values{}
	v.Set("include_email", "true")

	resp, err := oc.Get(nil, at, "https://api.twitter.com/1.1/account/verify_credentials.json", v)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 500 {
		return errors.New("Twitter is unavailable")
	}

	if resp.StatusCode >= 400 {
		return errors.New("Twitter request is invalid")
	}

	err = json.NewDecoder(resp.Body).Decode(user)
	if err != nil {
		return err
	}

	return nil

}
