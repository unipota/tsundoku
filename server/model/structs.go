package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Book struct {
	ID             string    `gorm:"type:char(36);not null;primary_key"`
	ISBN           string    `gorm:"type:char(13)"`
	Title          string    `gorm:"type:char(60) not null;"`
	Author         string    `gorm:"type:char(60);"`
	TotalPages     int       `gorm:""`
	RegularPrice   int       `gorm:""`
	Caption        string    `gorm:"type:TEXT;"`
	Publisher      string    `gorm:"type:char(60);"`
	CoverImageUrl  string    `gorm:"type:char(200);"`
	Memo           string    `gorm:"type:text;"`
	PurchasedPrice int       `gorm:""`
	DeviceID       string    `gorm:"char(36);not null"`
	CreatedAt      time.Time `gorm:"precision:6"`
	UpdatedAt      time.Time `gorm:"precision:6"`
}

func (book *Book) BeforeCreate(scope gorm.Scope) error {
	scope.SetColumn("ID", uuid.New())
	return nil
}

type Device struct {
	ID        string    `gorm:"type:char(36);not null;primary_key"`
	CreatedAt time.Time `gorm:"precision:6"`
	UpdatedAt time.Time `gorm:"precision:6"`
}

func (device *Device) BeforeCreate(scope gorm.Scope) error {
	scope.SetColumn("ID", uuid.New())
	return nil
}

type DeviceUser struct {
	ID        string    `gorm:"type:char(36);not null;primary_key"`
	DeviceID  string    `gorm:"type:char(36);not null;primary_key"`
	UserID    string    `gorm:"type:char(36);not null;primary_key"`
	CreatedAt time.Time `gorm:"precision:6"`
	UpdatedAt time.Time `gorm:"precision:6"`
}

func (deviceUser *DeviceUser) BeforeCreate(scope gorm.Scope) error {
	scope.SetColumn("ID", uuid.New())
	return nil
}

type User struct {
	ID        string    `gorm:"type:char(36);not null;primary_key"`
	Name      string    `gorm:"type:char(40);not null;"`
	CreatedAt time.Time `gorm:"precision:6"`
	UpdatedAt time.Time `gorm:"precision:6"`
}

func (user *User) BeforeCreate(scope gorm.Scope) error {
	scope.SetColumn("ID", uuid.New())
	return nil
}

type Social struct {
	ID         string    `gorm:"type:char(36);not null;primary_key"`
	Type       string    `gorm:"type:char(40);not null;"`
	Identifier string    `gorm:"type:char(40);"`
	UserID     string    `gorm:"type:char(36);not null;"`
	CreatedAt  time.Time `gorm:"precision:6"`
	UpdatedAt  time.Time `gorm:"precision:6"`
}

func (social *Social) BeforeCreate(scope gorm.Scope) error {
	scope.SetColumn("ID", uuid.New())
	return nil
}

type BookHistory struct {
	ID        string    `gorm:"type:char(36);not null;primary_key"`
	BookID    string    `gorm:"type:char(36);not null;"`
	CreatedAt time.Time `gorm:"precision:6"`
	UpdatedAt time.Time `gorm:"precision:6"`
}

func (bookHistory *BookHistory) BeforeCreate(scope gorm.Scope) error {
	scope.SetColumn("ID", uuid.New())
	return nil
}
