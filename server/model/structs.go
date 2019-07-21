package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Base struct {
	ID        uuid.UUID  `gorm:"type:char(36);primary_key;"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"update_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}

func (base *Base) BeforeCreate(scope *gorm.Scope) error {
	id := uuid.New()
	return scope.SetColumn("ID", id)
}

type Book struct {
	Base
	ISBN           string `gorm:"type:char(13)"`
	Title          string `gorm:"type:char(60) not null;"`
	Author         string `gorm:"type:char(60);"`
	TotalPages     int    `gorm:""`
	RegularPrice   int    `gorm:""`
	Caption        string `gorm:"type:TEXT;"`
	Publisher      string `gorm:"type:char(60);"`
	CoverImageUrl  string `gorm:"type:char(200);"`
	Memo           string `gorm:"type:text;"`
	PurchasedPrice int    `gorm:""`
	DeviceID       string `gorm:"char(36);not null"`

	BookHistories []BookHistory
}

type Device struct {
	Base
}

type DeviceUser struct {
	Base
	DeviceID uuid.UUID `gorm:"type:char(36);not null;unique_index:device_user;primary_key"`
	UserID   uuid.UUID `gorm:"type:char(36);not null;unique_index:device_user;primary_key"`
}

type User struct {
	Base
	Name string `gorm:"type:char(40);not null;"`
}

type Social struct {
	Base
	Type       string    `gorm:"type:char(40);not null;"`
	Identifier string    `gorm:"type:char(40);"`
	UserID     uuid.UUID `gorm:"type:char(36);not null;"`
}

type BookHistory struct {
	Base
	BookID   uuid.UUID `gorm:"type:char(36);not null;"`
	ReadPage int       `gorm:""`
}
