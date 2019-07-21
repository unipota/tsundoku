package router

import (
	"encoding/json"
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

	volumes, err := searchBooks(values)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, volumes)
}

func SearchWithWord(c echo.Context) error {
	search := c.QueryParam("search")

	values := url.Values{}
	values.Add("q", search)

	volumes, err := searchBooks(values)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, volumes)

}

func searchBooks(values url.Values) (*books.Volumes, error) {
	resp, err := http.Get("https://www.googleapis.com/books/v1/volumes" + "?" + values.Encode())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	volumes := &books.Volumes{}
	json.Unmarshal(body, volumes)

	return volumes, nil
}
