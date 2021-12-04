package moduls

import (
	"log"
	"reflect"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DB_Conn() {
    // Connect to SQLITE DB server
    var err error
    DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        panic("Failed to connect to DB")
    }

    log.Println(reflect.TypeOf(DB))
    DB.AutoMigrate(Book{})
}
