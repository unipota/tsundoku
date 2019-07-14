package model

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Book struct {
	ID            string `gorm:"type:char(36);not null;primary_key"`
	ISBN          string `gorm:"type:char(13)"`
	Title         string `gorm:"type:char(60);not null"`
	Author        string `gorm:"type:char(60)"`
	TotalPages    int    `gorm:""`
	RegularPrice  int    `gorm:""`
	Caption       string `gorm:"type:TEXT"`
	Publisher     string `gorm:"type:char(60)"`
	CoverImageUrl string `gorm:"type:char(200)"`
}

func (book *Book) BeforeCreate(scope gorm.Scope) error {
	scope.SetColumn("ID", uuid.New())
	return nil
}
