package model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"golang.org/x/xerrors"
	os "os"
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

	db.AutoMigrate(&Book{})
	return db, nil

}
