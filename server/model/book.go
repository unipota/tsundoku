package model

import (
	"strings"

	"github.com/google/uuid"
)

func PostNewBook(bookRecord BookRecord, deviceID uuid.UUID) error {
	bookData := &Book{
		ISBN:           bookRecord.ISBN,
		Title:          bookRecord.Title,
		TotalPages:     int(bookRecord.TotalPages),
		Caption:        bookRecord.Caption,
		Publisher:      bookRecord.Publisher,
		CoverImageUrl:  bookRecord.CoverImageURL,
		Memo:           bookRecord.Memo,
		PurchasedPrice: bookRecord.Price,
		DeviceID:       deviceID.String(),
	}

	bookData.RegularPrice = GetPriceWithISBN(bookRecord.ISBN)
	bookData.Author = strings.Join(bookRecord.Author, ",")

	deviceCount := 0
	err := db.Table("devices").Count(&deviceCount).Error
	if err != nil {
		return err
	}

	if deviceCount == 0 {
		device := Device{}
		device.Base.ID = deviceID
		err = db.Table("devices").Create(device).Error
		if err != nil {
			return err
		}
	}

	err = db.Table("books").Create(bookData).Error
	if err != nil {
		return err
	}

	return nil
}
