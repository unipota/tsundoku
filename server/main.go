package main

import (
	"fmt"
	"log"
	"os"
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
	e.GET("/ogp/:shareID", router.GetOGPImageHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Fatal(e.Start(":" + port))
}
