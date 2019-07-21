package model

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/seihmd/openbd"
	"google.golang.org/api/books/v1"
)

type BookRecord struct {
	ID            string   `json:"id"`
	ISBN          string   `json:"isbn"`
	Title         string   `json:"title"`
	Author        []string `json:"author"`
	TotalPages    int64    `json:"total_pages"`
	Price         int      `json:"price"`
	Caption       string   `json:"caption"`
	Publisher     string   `json:"publisher"`
	CoverImageURL string   `json:"cover_image_url`
	ReadPages     int      `json:"read_pages"`
	Memo          string   `json:"memo"`
}

func SearchBooks(values url.Values) (*books.Volumes, error) {
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

func ConvertVolumeToBookRecord(volume *books.Volume) *BookRecord {
	bookRecord := &BookRecord{
		ID:            "",
		ISBN:          "",
		Title:         volume.VolumeInfo.Title,
		Author:        volume.VolumeInfo.Authors,
		TotalPages:    volume.VolumeInfo.PageCount,
		Price:         0,
		Caption:       "",
		Publisher:     volume.VolumeInfo.Publisher,
		CoverImageURL: "",
		ReadPages:     0,
		Memo:          "",
	}

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

func VolumesToBookRecords(volumes *books.Volumes) []*BookRecord {
	bookRecords := make([]*BookRecord, 0)

	for _, volume := range volumes.Items {
		bookRecord := ConvertVolumeToBookRecord(volume)
		bookRecord = UpdateBookRecordWithOpenDB(bookRecord)
		bookRecords = append(bookRecords, bookRecord)
	}

	return bookRecords
}

func UpdateBookRecordWithOpenDB(bookRecord *BookRecord) *BookRecord {
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

func GetPriceWithISBN(isbn string) int {
	reqularPrice := 0
	openBD := openbd.New()
	book, err := openBD.Get(isbn)
	if err != nil {
		return reqularPrice
	}

	supplyData := book.ProductSupply.SupplyDetail
	for _, priceData := range supplyData.Price {
		if priceData.CurrencyCode == "JPY" {
			price, err := priceData.PriceAmount.Int64()
			if err != nil {
				return reqularPrice
			}
			reqularPrice = int(price)
			break
		}
	}

	return reqularPrice
}
