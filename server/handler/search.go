package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/labstack/echo"
	"google.golang.org/api/books/v1"
)

type (
	ISBN struct {
		ISBN string `json:"isbn"`
	}
)

func SearchWithISBN(c echo.Context) error {
	isbn := ISBN{}
	c.Bind(&isbn)

	values := url.Values{}
	values.Add("q", "isbn:"+isbn.ISBN)
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
