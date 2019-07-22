package router

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/unipota/tsundoku/server/model"
)

func GetBookListHandler(c echo.Context) error {
	return c.NoContent(http.StatusNotImplemented)
}

func PostNewBookHandler(c echo.Context) error {
	bookRecord := model.BookRecord{}
	c.Bind(&bookRecord)
	deviceID := c.Get("deviceID").(uuid.UUID)
	err := model.AddNewBook(bookRecord, deviceID)
	if err != nil {
		fmt.Println(err)
		return c.NoContent(http.StatusInternalServerError)
	}
	return c.NoContent(http.StatusCreated)
}

func PutUpdateBookHandler(c echo.Context) error {
	return c.NoContent(http.StatusNotImplemented)
}

func DeleteBookHandler(c echo.Context) error {
	return c.NoContent(http.StatusNotImplemented)
}
