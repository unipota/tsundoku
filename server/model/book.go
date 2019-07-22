package model

import (
	"strings"

	"github.com/google/uuid"
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

func AddNewBook(bookRecord BookRecord, deviceID uuid.UUID) error {
	bookData := &Book{
		ISBN:          bookRecord.ISBN,
		Title:         bookRecord.Title,
		TotalPages:    int(bookRecord.TotalPages),
		Caption:       bookRecord.Caption,
		Publisher:     bookRecord.Publisher,
		CoverImageUrl: bookRecord.CoverImageURL,
		Memo:          bookRecord.Memo,
		Price:         bookRecord.Price,
		DeviceID:      deviceID,
	}

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

func GetBooksByDeviceID(deviceID uuid.UUID) (*[]Book, error) {
	books := &[]Book{}
	if err := db.Where("device_id IN ?", db.Table("device_users").Where("user_id = ?", db.Table("device_users").Where("device_id = ?", deviceID).Select("user_id").SubQuery()).Select("device_id").SubQuery()).Preload("BookHistories").Find(books).Error; err != nil {
		return nil, err
	}
	return books, nil
}
