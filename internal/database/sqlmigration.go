package database

import (
	"github.com/jinzhu/gorm"
	"log"
)

func ConnectMysql(databaseConnectionString string) *gorm.DB {
	db, err := gorm.Open("mysql", databaseConnectionString)
	if err != nil {
		log.Print(err)
	}
	return db
}

func SqlMigrate(db *gorm.DB){
	db.AutoMigrate(&ShopUsers{})
	db.AutoMigrate(&CustomersData{})
}
