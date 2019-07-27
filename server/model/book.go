package model

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

func NewBook(book *Book) (*Book, error) {
	book.ReadHistories = []ReadHistory{
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
	err := db.Where("device_id IN ?", sub2).Or("device_id = ?", deviceID).Preload("ReadHistories", func(db *gorm.DB) *gorm.DB {
		return db.Order("read_histories.created_at DESC")
	}).Find(&books).Error
	if err != nil {
		return nil, err
	}
	return books, nil
}

func GetBookByBookID(bookID uuid.UUID) (*Book, error) {
	book := &Book{}
	if err := db.Where("id = ?", bookID).Preload("ReadHistories", func(db *gorm.DB) *gorm.DB {
		return db.Order("read_histories.created_at DESC")
	}).First(book).Error; err != nil {
		return nil, err
	}

	return book, nil
}

func IsOwnBook(bookID, deviceID uuid.UUID) bool {
	count := 0
	sub1 := db.Table("device_users").Where("device_id = ?", deviceID).Select("user_id").SubQuery()
	sub2 := db.Table("device_users").Where("user_id = ?", sub1).Select("device_id").SubQuery()
	if err := db.Table("books").Where("device_id IN ? OR device_id = ?", sub2, deviceID).Where("id = ?", bookID).Count(&count).Error; err != nil {
		return false
	}
	return count > 0
}
