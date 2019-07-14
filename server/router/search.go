package router

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/labstack/echo/v4"
	"google.golang.org/api/books/v1"
)

type (
	ISBN struct {
		ISBN string `json:"isbn"`
	}
)

func SearchWithISBN(c echo.Context) error {
	isbn := c.QueryParam("isbn")

	values := url.Values{}
	values.Add("q", "isbn:"+isbn)
	resp, err := http.Get("https://www.googleapis.com/books/v1/volumes" + "?" + values.Encode())

	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	volumes := books.Volumes{}
	json.Unmarshal(body, &volumes)

	return c.JSON(http.StatusOK, volumes)
}
