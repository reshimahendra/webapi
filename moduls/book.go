package moduls

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

// InitDB sets up setting up the connection pool global variable.
// func InitDB(dataSourceName string) error {
//     var err error
//
//     db, err = sql.Open("sqlite3", dataSourceName)
//     if err != nil {
//         return err
//     }
//
//     return db.Ping()
// }

type Table interface {
    First()
    Last()
    All()
}

type BookInput struct{
    Title        string `json:"title" binding:"required"`
    Price        json.Number `json:"price" binding:"required,number,gt=0"`
    SubTitle     string `json:"sub_title" binding:"required"`
    Description  string
    Release_date string
}

type Book struct {
    gorm.Model
    ID          int `json: "id"`
    Title       string `json: "title" binding:"required"` 
    Description string `json: "description"`
    Price       int `json: "price" binding: "required, number, gte=0"`
    Rating      int8 `json: "rating" binding: "required, number, gte=0, lte=5"`
    CreatedAt   time.Time `json: "created_at"`
    UpdatedAt   time.Time `json: : "updated_at"`
}

func (b Book) First() (Book) {
    DB.First(&b)
    return b
}

func (b Book) Last() (Book) {
    DB.Last(&b)
    return b
}

func (b Book) All() ([]Book) {
    var books []Book
    DB.Debug().Find(&books)
    return books
}

