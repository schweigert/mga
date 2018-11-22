package db

import (
	"log"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var dbcChan chan *gorm.DB

func init() {
	dbcChan = make(chan *gorm.DB, 5)
	go dbcChanProducer()
}

func dbcChanProducer() {
	for {
		dbc, err := gorm.Open("postgres", os.Getenv("DB_CONFIG"))
		if err != nil {
			log.Println(err)
			time.Sleep(2 * time.Second)
		} else {
			dbc.LogMode(true)
			dbcChan <- dbc
		}
	}
}

func Connect() *gorm.DB {
	return <-dbcChan
}

func SafeClose(dbc *gorm.DB) {
	err := dbc.Close()
	if err != nil {
		log.Panic(err)
	}
}
