package router

import (
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/unipota/tsundoku/server/model"
)

type BookRecord struct {
	ID            string    `json:"id"`
	ISBN          string    `json:"isbn"`
	Title         string    `json:"title"`
	Author        []string  `json:"author"`
	TotalPages    int       `json:"totalPages"`
	Price         int       `json:"price"`
	Caption       string    `json:"caption"`
	Publisher     string    `json:"publisher"`
	CoverImageURL string    `json:"coverImageUrl"`
	ReadPages     int       `json:"readPages"`
	Memo          string    `json:"memo"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}

type ReadHistory struct {
	ReadPages int       `json:"readPages"`
	CreatedAt time.Time `json:"createdAt"`
}

type BookDetail struct {
	ID            string        `json:"id"`
	ISBN          string        `json:"isbn"`
	Title         string        `json:"title"`
	Author        []string      `json:"author"`
	TotalPages    int           `json:"totalPages"`
	Price         int           `json:"price"`
	Caption       string        `json:"caption"`
	Publisher     string        `json:"publisher"`
	CoverImageURL string        `json:"coverImageUrl"`
	ReadHistories []ReadHistory `json:"readHistories"`
	Memo          string        `json:"memo"`
	CreatedAt     time.Time     `json:"createdAt"`
	UpdatedAt     time.Time     `json:"updatedAt"`
}

type BookStat struct {
	ID            string        `json:"id"`
	Title         string        `json:"title"`
	TotalPages    int           `json:"totalPages"`
	Price         int           `json:"price"`
	ReadHistories []ReadHistory `json:"readHistories"`
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

func GetBookDetailHandler(c echo.Context) error {
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

	book, err := model.GetBookByBookID(bookID)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, H{"Error with get book"})
	}
	return c.JSON(http.StatusOK, book2BookDetail(*book))

}

func PostNewBookHandler(c echo.Context) error {
	searchedBook := SearchedBook{}
	if err := c.Bind(&searchedBook); err != nil {
		return c.JSON(http.StatusBadRequest, H{"Bad request"})
	}

	deviceID := c.Get("deviceID").(uuid.UUID)
	book := searchedBook2Book(searchedBook, deviceID)
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

	if bookRecord.ID != bookID.String() {
		return c.JSON(http.StatusBadRequest, H{"Bad request"})
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

func GetBookStatsHandler(c echo.Context) error {
	deviceID := c.Get("deviceID").(uuid.UUID)
	books, err := model.GetBooksByDeviceID(deviceID)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, H{"Error with get books"})
	}

	bookStats := make([]BookStat, 0, len(books))
	for _, book := range books {
		bookStats = append(bookStats, book2BookStat(*book))
	}

	return c.JSON(http.StatusOK, bookStats)
}

func bookRecord2Book(bookRecord BookRecord, deviceID uuid.UUID) model.Book {
	book := model.Book{
		ISBN:          bookRecord.ISBN,
		Title:         bookRecord.Title,
		Author:        strings.Join(bookRecord.Author, ","),
		TotalPages:    bookRecord.TotalPages,
		Caption:       bookRecord.Caption,
		Publisher:     bookRecord.Publisher,
		CoverImageUrl: bookRecord.CoverImageURL,
		Memo:          bookRecord.Memo,
		Price:         bookRecord.Price,
		DeviceID:      deviceID,
		ReadHistories: []model.ReadHistory{
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
		Title:         book.Title,
		Author:        strings.Split(book.Author, ","),
		TotalPages:    book.TotalPages,
		Price:         book.Price,
		Caption:       book.Caption,
		Publisher:     book.Publisher,
		CoverImageURL: book.CoverImageUrl,
		ReadPages:     book.ReadHistories[0].ReadPage,
		Memo:          book.Memo,
		CreatedAt:     book.CreatedAt,
		UpdatedAt:     book.UpdatedAt,
	}
}

func book2BookDetail(book model.Book) BookDetail {
	bookDetail := BookDetail{
		ID:            book.ID.String(),
		ISBN:          book.Title,
		Title:         book.ISBN,
		Author:        strings.Split(book.Author, ","),
		TotalPages:    book.TotalPages,
		Price:         book.Price,
		Caption:       book.Caption,
		Publisher:     book.Publisher,
		CoverImageURL: book.CoverImageUrl,
		Memo:          book.Memo,
		CreatedAt:     book.CreatedAt,
		UpdatedAt:     book.UpdatedAt,
	}

	for _, history := range book.ReadHistories {
		bookDetail.ReadHistories = append(bookDetail.ReadHistories, ReadHistory{
			history.ReadPage,
			history.CreatedAt,
		})
	}
	return bookDetail
}

func searchedBook2Book(searchedBook SearchedBook, deviceID uuid.UUID) model.Book {
	book := model.Book{
		ISBN:          searchedBook.ISBN,
		Title:         searchedBook.Title,
		Author:        strings.Join(searchedBook.Author, ","),
		TotalPages:    getTotalPages(searchedBook.ISBN),
		Caption:       searchedBook.Caption,
		Publisher:     searchedBook.Publisher,
		CoverImageUrl: searchedBook.CoverImageURL,
		Memo:          searchedBook.Memo,
		Price:         searchedBook.Price,
		DeviceID:      deviceID,
	}
	id, err := uuid.Parse(searchedBook.ID)
	if err != nil {
		return book
	}
	book.ID = id
	return book

}

func getTotalPages(isbn string) int {
	values := url.Values{}
	values.Add("q", "isbn:"+isbn)
	volumes, err := searchWithGoogle(values)

	if err != nil || volumes.TotalItems == 0 {
		return 0
	}

	return int(volumes.Items[0].VolumeInfo.PageCount)
}

func book2BookStat(book model.Book) BookStat {
	bookStat := BookStat{
		ID:         book.ID.String(),
		Title:      book.Title,
		TotalPages: book.TotalPages,
		Price:      book.Price,
	}

	bookStat.ReadHistories = make([]ReadHistory, 0)
	for _, fromBook := range book.ReadHistories {
		readHistory := ReadHistory{
			ReadPages: fromBook.ReadPage,
			CreatedAt: fromBook.CreatedAt,
		}
		bookStat.ReadHistories = append(bookStat.ReadHistories, readHistory)
	}

	return bookStat
}
