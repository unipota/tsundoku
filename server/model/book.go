package model

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

func GetBooksByDeviceID(deviceID uuid.UUID) (*[]Book, error) {
	books := &[]Book{}
	if err := db.Where("device_id IN ?", db.Table("device_users").Where("user_id = ?", db.Table("device_users").Where("device_id = ?", deviceID).Select("user_id").SubQuery()).Select("device_id").SubQuery()).Preload("BookHistories", func(db *gorm.DB) *gorm.DB {
		return db.Order("book_histories.created_at DESC")
	}).Find(books).Error; err != nil {
		return nil, err
	}
	return books, nil
}
