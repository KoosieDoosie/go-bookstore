package config

import (
	"github.com/jinzhu/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "KoosieDoosie:Ss27262726/simplerest?charset=utf8&parseTime=True&loc=Local") // Open connection with the database
	if err != nil {
		panic(err)
	}
	db = d

}

func GetDB() *gorm.DB {
	return db
}
