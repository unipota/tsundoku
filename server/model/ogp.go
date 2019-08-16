package model

import (
	"github.com/google/uuid"
)

func NewShare(share *Share) (*Share, error) {
	if err := db.Create(share).Error; err != nil {
		return nil, err
	}
	return share, nil
}

func GetShareByshareID(shareID uuid.UUID) (*Share, error) {
	share := &Share{}
	if err := db.Where("id = ?", shareID).First(share).Error; err != nil {
		return nil, err
	}
	return share, nil
}

func IsExistShareID(shareID uuid.UUID) bool {
	count := 0
	if err := db.Table("shares").Where("id = ?", shareID).Count(&count).Error; err != nil {
		return false
	}
	return count > 0
}
