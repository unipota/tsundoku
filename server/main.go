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
	e.Use(middleware.Logger())
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:  "./static",
		HTML5: true,
	}))

	api := e.Group("/api", session.Middleware(store), router.IdentifyMiddleware)
	api.GET("/ping", router.Ping)

	api.GET("/search/isbn", router.SearchWithISBN)
	api.GET("/search", router.SearchWithWord)

	api.GET("/books", router.GetBookListHandler)
	api.POST("/books", router.PostNewBookHandler)
	api.PUT("/books/:id", router.PutUpdateBookHandler)
	api.DELETE("/books/:id", router.DeleteBookHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Fatal(e.Start(":" + port))
}
