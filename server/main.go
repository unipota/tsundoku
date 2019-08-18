package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/srinathgs/mysqlstore"
	"github.com/unipota/tsundoku/server/model"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/unipota/tsundoku/server/router"
)

func main() {
	db, err := model.EstablishConnection()
	if err != nil {
		panic(err)
	}

	for {
		if err := db.DB().Ping(); err != nil {
			fmt.Println(err)
			time.Sleep(time.Second * 3)
		} else {
			break
		}
	}

	if err := model.Migration(); err != nil {
		panic(err)
	}

	cookieSecret := os.Getenv("COOKIE_SECRET")
	if cookieSecret == "" {
		cookieSecret = "tsundoku"
	}

	store, err := mysqlstore.NewMySQLStoreFromConnection(db.DB(), "session", "/", 60*60*24*14, []byte(cookieSecret))
	if err != nil {
		panic(err)
	}

	e := echo.New()
	e.Debug = true
	e.Use(middleware.Logger())

	// middleware.StaticWithConfigからほぼコピペ・一部改変
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			p := c.Request().URL.Path
			if strings.HasSuffix(c.Path(), "*") { // When serving from a group, e.g. `/static*`.
				p = c.Param("*")
			}
			p, err = url.PathUnescape(p)
			if err != nil {
				return
			}
			name := filepath.Join("./static", path.Clean("/"+p)) // "/"+ for security

			fi, err := os.Stat(name)
			if err != nil {
				if os.IsNotExist(err) {
					if err = next(c); err != nil {
						if he, ok := err.(*echo.HTTPError); ok {
							if he.Code == http.StatusNotFound {
								// TODO: OGP周り
								return c.File("./static/index.html")
							}
						}
						return
					}
				}
				return
			}

			if fi.IsDir() {
				// トップページ('/')にアクセスしてきた時
				index := filepath.Join(name, "index.html")
				fi, err = os.Stat(index)

				if err != nil {
					if os.IsNotExist(err) {
						return next(c)
					}
					return
				}
				// TODO: OGP周り
				return c.File(index)
			}

			return c.File(name)
		}
	})

	auth := e.Group("/auth", session.Middleware(store), router.IdentifyMiddleware, router.LoginedUserRedirect)
	auth.GET("/twitter/oauth", router.GetTwitterAuthHandler)
	auth.GET("/twitter/callback", router.GetTwitterCallbackHandler)

	auth.GET("/google/oauth", router.GetGoogleAuthHandler)
	auth.GET("/google/callback", router.GetGoogleCallbackHandler)

	auth.GET("/github/oauth", router.GetGitHubAuthHandler)
	auth.GET("/github/callback", router.GetGitHubCallbackHandler)

	api := e.Group("/api", session.Middleware(store), router.IdentifyMiddleware)
	api.GET("/ping", router.Ping)
	api.POST("/logout", router.PostLogoutHandler)
	api.GET("/logout", router.PostLogoutHandler)

	api.GET("/users/me", router.GetWhoAmIHandler)

	api.GET("/search/isbn", router.SearchWithISBN)
	api.GET("/search", router.SearchWithWord)

	api.GET("/books", router.GetBookListHandler)
	api.POST("/books", router.PostNewBookHandler)
	api.GET("/books/:bookID", router.GetBookDetailHandler)
	api.PUT("/books/:bookID", router.PutUpdateBookHandler)
	api.DELETE("/books/:bookID", router.DeleteBookHandler)
	api.GET("/books/statistics", router.GetBookStatsHandler)

	api.GET("/share", router.GetShareURLHandler)

	e.GET("/share/:shareID", router.GetSharePageHandler)
	e.GET("/share/image/:shareID", router.GetOGPImageHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Fatal(e.Start(":" + port))
}
