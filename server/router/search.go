package router

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/labstack/echo/v4"
	"github.com/seihmd/openbd"
	"google.golang.org/api/books/v1"
)

type ISBN struct {
	ISBN string `json:"isbn"`
}

func SearchWithISBN(c echo.Context) error {
	isbn := c.QueryParam("isbn")

	values := url.Values{}
	values.Add("q", "isbn:"+isbn)

	volumes, err := searchBooks(values)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	bookRecords := volumesToBookRecords(volumes)
	return c.JSON(http.StatusOK, bookRecords)
}

func SearchWithWord(c echo.Context) error {
	search := c.QueryParam("search")

	values := url.Values{}
	values.Add("q", search)

	volumes, err := searchBooks(values)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	bookRecords := volumesToBookRecords(volumes)

	return c.JSON(http.StatusOK, bookRecords)

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
	if err := json.Unmarshal(body, volumes); err != nil {
		return nil, err
	}

	return volumes, nil
}

func convertVolumeToBookRecord(volume *books.Volume) *BookRecord {
	bookRecord := &BookRecord{
		ID:            "",
		ISBN:          "",
		Title:         volume.VolumeInfo.Title,
		Author:        volume.VolumeInfo.Authors,
		TotalPages:    int(volume.VolumeInfo.PageCount),
		Price:         0,
		Caption:       "",
		Publisher:     volume.VolumeInfo.Publisher,
		CoverImageURL: "",
		ReadPages:     0,
		Memo:          "",
	}

	if volume.VolumeInfo.ImageLinks != nil {
		imageLinks := [...]string{
			volume.VolumeInfo.ImageLinks.ExtraLarge,
			volume.VolumeInfo.ImageLinks.Large,
			volume.VolumeInfo.ImageLinks.Medium,
			volume.VolumeInfo.ImageLinks.Small,
			volume.VolumeInfo.ImageLinks.Thumbnail,
			volume.VolumeInfo.ImageLinks.SmallThumbnail,
		}
		for _, imageLink := range imageLinks {
			if imageLink != "" {
				bookRecord.CoverImageURL = imageLink
				break
			}
		}
	}

	industryIdentifiers := volume.VolumeInfo.IndustryIdentifiers
	isbn10 := ""
	isbn13 := ""
	for _, idetifier := range industryIdentifiers {
		if idetifier.Type == "ISBN_10" {
			isbn10 = idetifier.Identifier
		} else if idetifier.Type == "ISBN_13" {
			isbn13 = idetifier.Identifier
		}
	}
	if isbn13 != "" {
		bookRecord.ISBN = isbn13
	} else if isbn10 != "" {
		bookRecord.ISBN = isbn10
	}

	return bookRecord
}

func volumesToBookRecords(volumes *books.Volumes) []*BookRecord {
	bookRecords := make([]*BookRecord, 0)

	for _, volume := range volumes.Items {
		bookRecord := convertVolumeToBookRecord(volume)
		bookRecord = updateBookRecordWithOpenDB(bookRecord)
		bookRecords = append(bookRecords, bookRecord)
	}

	return bookRecords
}

func updateBookRecordWithOpenDB(bookRecord *BookRecord) *BookRecord {
	openBD := openbd.New()
	book, err := openBD.Get(bookRecord.ISBN)
	if err != nil {
		return bookRecord
	}

	if bookRecord.Publisher == "" {
		bookRecord.Publisher = book.Publisher
	}

	supplyData := book.ProductSupply.SupplyDetail
	for _, priceData := range supplyData.Price {
		if priceData.CurrencyCode == "JPY" {
			price, err := priceData.PriceAmount.Int64()
			if err != nil {
				return bookRecord
			}
			bookRecord.Price = int(price)
			break
		}
	}

	return bookRecord
}
