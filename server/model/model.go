package model

import (
	"fmt"
	os "os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"golang.org/x/xerrors"
)

var (
	db *gorm.DB
)

func EstablishConnection() (*gorm.DB, error) {
	dbUser := os.Getenv("MYSQL_USERNAME")
	if dbUser == "" {
		dbUser = "root"
	}

	dbPass := os.Getenv("MYSQL_PASSWORD")
	if dbPass == "" {
		dbPass = "mysql"
	}

	dbHost := os.Getenv("MYSQL_HOSTNAME")
	if dbHost == "" {
		dbHost = "db"
	}

	dbPort := os.Getenv("MYSQL_PORT")
	if dbPort == "" {
		dbPort = "3306"
	}

	dbName := os.Getenv("MYSQL_DATABASE")
	if dbName == "" {
		dbName = "tsundoku"
	}

	_db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&collation=utf8mb4_general_ci&parseTime=true", dbUser, dbPass, dbHost, dbPort, dbName))
	if err != nil {
		return nil, xerrors.Errorf("Can't Connect to DATABASE: %w", err)
	}
	db = _db
	db = db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4")

	return db, nil
}

func Migration() error {
	if err := db.AutoMigrate(allTables...).Error; err != nil {
		return err
	}

	foreignKeys := [][5]string{
		{"books", "device_id", "devices(id)", "CASCADE", "CASCADE"},
		{"book_histories", "book_id", "books(id)", "CASCADE", "CASCADE"},
		{"device_users", "device_id", "devices(id)", "CASCADE", "CASCADE"},
		{"device_users", "user_id", "users(id)", "CASCADE", "CASCADE"},
		{"socials", "user_id", "users(id)", "CASCADE", "CASCADE"},
	}

	for _, c := range foreignKeys {
		fmt.Println(c)
		if err := db.Table(c[0]).AddForeignKey(c[1], c[2], c[3], c[4]).Error; err != nil {
			return err
		}
	}

	return nil
}

func IsErrRecordNotFound(err error) bool {
	return xerrors.Is(err, gorm.ErrRecordNotFound)
}

var allTables = []interface{}{
	&Book{},
	&Device{},
	&DeviceUser{},
	&User{},
	&Social{},
	&ReadHistory{},
}
