package models

import (
	_ "github.com/go-sql-driver/mysql"

	"github.com/jinzhu/gorm"
)

// InitDB creates and migrates the database
func InitDB() (*gorm.DB, error) {
	var err error
	dbUser := "root"
	dbPassword := ""
	dbName := "go-task"
	dsn := dbUser + ":" + dbPassword + "@/" + dbName + "?charset=utf8&parseTime=True"
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	db.LogMode(true)
	db.AutoMigrate(&User{}, &Task{}, &Status{})

	return db, nil
}
