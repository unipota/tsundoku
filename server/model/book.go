package model

import (
	"github.com/google/uuid"
)

func GetBooksByDeviceID(deviceID uuid.UUID) (*[]Book, error) {
	books := &[]Book{}
	if err := db.Where("device_id IN ?", db.Table("device_users").Where("user_id = ?", db.Table("device_users").Where("device_id = ?", deviceID).Select("user_id").SubQuery()).Select("device_id").SubQuery()).Preload("BookHistories").Find(books).Error; err != nil {
		return nil, err
	}
	return books, nil
}
