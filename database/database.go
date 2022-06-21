package database

// Veritabanının kurulumu ve ayarlanması için oluşturulmuş bir dosya.

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var dbuser = "root"
var dbpass = "123"
var dburi = "todo-db"
var dbport = "3306"
var dbname = "todos"

func ConnectDB() error {

	connUri := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		dbuser, dbpass, dburi, dbport, dbname)

	dsn := connUri + "?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	return nil
}
