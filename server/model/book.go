package model

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

func NewBook(book *Book) (*Book, error) {
	book.BookHistories = []BookHistory{
		{},
	}
	if err := db.Create(book).Error; err != nil {
		return nil, err
	}

	return book, nil
}

func UpdateBook(book *Book) (*Book, error) {
	if err := db.Save(book).Error; err != nil {
		return nil, err
	}

	return book, nil
}

func DeleteBook(bookID uuid.UUID) error {
	return db.Delete(Book{Base: Base{ID: bookID}}).Error
}

func GetBooksByDeviceID(deviceID uuid.UUID) ([]*Book, error) {
	books := []*Book{}
	sub1 := db.Table("device_users").Where("device_id = ?", deviceID).Select("user_id").SubQuery()
	sub2 := db.Table("device_users").Where("user_id = ?", sub1).Select("device_id").SubQuery()
	err := db.Where("device_id IN ?", sub2).Or("device_id = ?", deviceID).Preload("BookHistories", func(db *gorm.DB) *gorm.DB {
		return db.Order("book_histories.created_at DESC")
	}).Find(&books).Error
	if err != nil {
		return nil, err
	}
	return books, nil
}

func IsOwnBook(bookID, deviceID uuid.UUID) bool {
	count := 0
	if err := db.Table("books").Where("device_id IN ?", db.Table("device_users").Where("user_id = ?", db.Table("device_users").Where("device_id = ?", deviceID).Select("user_id").SubQuery()).Select("device_id").SubQuery()).Or("device_id = ?", deviceID).Where("id = ?", bookID).Count(&count).Error; err != nil {
		return false
	}
	return count > 0
}
