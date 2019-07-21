package router

import (
	"net/http"
	"net/url"

	"github.com/labstack/echo/v4"
	"github.com/unipota/tsundoku/server/model"
)

type ISBN struct {
	ISBN string `json:"isbn"`
}

func SearchWithISBN(c echo.Context) error {
	isbn := c.QueryParam("isbn")

	values := url.Values{}
	values.Add("q", "isbn:"+isbn)

	volumes, err := model.SearchBooks(values)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	bookRecords := model.VolumesToBookRecords(volumes)

	return c.JSON(http.StatusOK, bookRecords)
}

func SearchWithWord(c echo.Context) error {
	search := c.QueryParam("search")

	values := url.Values{}
	values.Add("q", search)

	volumes, err := model.SearchBooks(values)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	bookRecords := model.VolumesToBookRecords(volumes)

	return c.JSON(http.StatusOK, bookRecords)

}
