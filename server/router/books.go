package router

import (
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/unipota/tsundoku/server/model"
)

type BookRecord struct {
	ID            string   `json:"id"`
	ISBN          string   `json:"isbn"`
	Title         string   `json:"title"`
	Author        []string `json:"author"`
	TotalPages    int      `json:"totalPages"`
	Price         int      `json:"price"`
	Caption       string   `json:"caption"`
	Publisher     string   `json:"publisher"`
	CoverImageURL string   `json:"coverImageUrl"`
	ReadPages     int      `json:"readPages"`
	Memo          string   `json:"memo"`
}

func GetBookListHandler(c echo.Context) error {
	deviceID := c.Get("deviceID").(uuid.UUID)
	books, err := model.GetBooksByDeviceID(deviceID)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, H{"Error with get books"})
	}

	booksForResponse := make([]BookRecord, 0, len(books))
	for _, book := range books {
		booksForResponse = append(booksForResponse, book2BookRecord(*book))
	}

	return c.JSON(http.StatusOK, booksForResponse)
}

func PostNewBookHandler(c echo.Context) error {
	bookRecord := BookRecord{}
	if err := c.Bind(&bookRecord); err != nil {
		return c.JSON(http.StatusBadRequest, H{"Bad request"})
	}
	deviceID := c.Get("deviceID").(uuid.UUID)
	book := bookRecord2Book(bookRecord, deviceID)
	newBook, err := model.NewBook(&book)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, H{"Error with create book"})
	}
	return c.JSON(http.StatusCreated, book2BookRecord(*newBook))
}

func PutUpdateBookHandler(c echo.Context) error {
	bookRecord := BookRecord{}
	if err := c.Bind(&bookRecord); err != nil {
		return c.JSON(http.StatusBadRequest, H{"Bad request"})
	}
	deviceID := c.Get("deviceID").(uuid.UUID)
	pathParam := c.Param("bookID")
	bookID, err := uuid.Parse(pathParam)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusNotFound, H{"invalid uuid"})
	}

	if !model.IsOwnBook(bookID, deviceID) {
		return c.JSON(http.StatusNotFound, H{"book not found"})
	}

	book := bookRecord2Book(bookRecord, deviceID)
	updatedBook, err := model.UpdateBook(&book)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, H{"Error with update book"})
	}

	return c.JSON(http.StatusOK, book2BookRecord(*updatedBook))
}

func DeleteBookHandler(c echo.Context) error {
	deviceID := c.Get("deviceID").(uuid.UUID)
	pathParam := c.Param("bookID")
	bookID, err := uuid.Parse(pathParam)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusNotFound, H{"invalid uuid"})
	}

	if !model.IsOwnBook(bookID, deviceID) {
		return c.JSON(http.StatusNotFound, H{"book not found"})
	}

	if err := model.DeleteBook(bookID); err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, H{"Error with update book"})
	}
	return c.NoContent(http.StatusOK)
}

func bookRecord2Book(bookRecord BookRecord, deviceID uuid.UUID) model.Book {
	book := model.Book{
		ISBN:          bookRecord.ISBN,
		Title:         bookRecord.Title,
		Author:        strings.Join(bookRecord.Author, ","),
		TotalPages:    bookRecord.TotalPages,
		Caption:       bookRecord.Caption,
		Publisher:     bookRecord.Caption,
		CoverImageUrl: bookRecord.CoverImageURL,
		Memo:          bookRecord.Memo,
		Price:         bookRecord.Price,
		DeviceID:      deviceID,
		BookHistories: []model.BookHistory{
			{ReadPage: bookRecord.ReadPages},
		},
	}
	id, err := uuid.Parse(bookRecord.ID)
	if err != nil {
		return book
	}
	book.ID = id
	return book
}

func book2BookRecord(book model.Book) BookRecord {
	return BookRecord{
		ID:            book.ID.String(),
		ISBN:          book.ISBN,
		Title:         book.ISBN,
		Author:        strings.Split(book.Author, ","),
		TotalPages:    book.TotalPages,
		Price:         book.Price,
		Caption:       book.Caption,
		Publisher:     book.Publisher,
		CoverImageURL: book.CoverImageUrl,
		ReadPages:     book.BookHistories[0].ReadPage,
		Memo:          book.Memo,
	}
}
