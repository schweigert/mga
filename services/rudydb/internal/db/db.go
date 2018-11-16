package db

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func Connect() *gorm.DB {
	db, err := gorm.Open("postgres", os.Getenv("DB_CONFIG"))
	if err != nil {
		log.Panic(err)
	}

	db.LogMode(true)
	return db
}

func SafeClose(dbc *gorm.DB) {
	err := dbc.Close()
	if err != nil {
		log.Panic(err)
	}
}
