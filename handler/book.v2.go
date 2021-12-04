package handler

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"webapi/moduls"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func BookPostHandler2(c *gin.Context) {
    var book moduls.BookInput
    err := c.ShouldBindJSON(&book)

    if err != nil {
        eMsg := []string{} 
        for _,e := range err.(validator.ValidationErrors) {
            msg := fmt.Sprintf("Error on field '%s', condition: '%s'", e.Field(), e.ActualTag)
            eMsg = append(eMsg, msg)
        }

        c.JSON(http.StatusBadRequest, gin.H{
            "error": eMsg,
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "title": book.Title,
        "sub_title": book.SubTitle,
        "price": book.Price,
        "description": book.Description,
        "release_date": book.Release_date,
    })
}

func CreateBookHandler(c *gin.Context) {
    var book moduls.Book
    err := c.ShouldBindJSON(&book)

    if err != nil {
        eMsg := []string{}
        for _, e:= range err.(validator.ValidationErrors) {
            msg := fmt.Sprintf("Field error :'%s', condition: '%s'", e.Field(), e.ActualTag())
            eMsg = append(eMsg, msg)
        }

        c.JSON(http.StatusBadRequest, gin.H{
            "error": eMsg,
        })

        return
    }
    log.Println(book)
    log.Println("Memasuki tahap memasukkan data")
    // Insert row into table
    moduls.DB.Create(&book)
    c.JSON(http.StatusOK, gin.H{
        "title": book.Title,
        "description": book.Description,
        "price": book.Price,
        "rating": book.Rating,
    })
}

func ShowBooksHandler(c *gin.Context) {
    var book moduls.Book
    books := book.All()

    c.JSON(http.StatusOK, books)
}

func SearchBookHandler(c *gin.Context) {
    // var book moduls.Book
    param := c.Request.URL.Query()

    var qCondition string
    var qValue string
    for i, value := range param {
        if len(param) > 1 {
            qCondition = fmt.Sprintf("%s AND %s ?", qCondition, i) 
        } else {
            if len(value) > 1 {
                qCondition = fmt.Sprintf("%s in ?", i)
            } else {
                qCondition = fmt.Sprintf("%s = ?", i)
            }
        }
        // log.Println("Panjangnya: ",len(value))
        if len(value) > 1 {
            for _, data := range value {
                // log.Println(i,":", data)
                // qValue =
            }
        } else {
            // log.Println("Test ", value[0])
            qValue = ""
        }
    }
}
