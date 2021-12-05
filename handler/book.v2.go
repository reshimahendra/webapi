package handler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
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
    var book moduls.Book
    // var judul string;

    log.Println(c.Request.URL.Query().Get("title"))
    if c.Request.URL.Query().Has("title") {
        e := moduls.DB.Debug().Where("title = ?", c.Request.URL.Query().Get("title")).Find(&book).Error
        if e != nil {
            log.Println("Ada kesalahan pencarian data")
            return
        } 
        c.JSON(http.StatusOK, book)
    } else {
        c.JSON(http.StatusNotFound, gin.H{
            "error": "No title field data",
        })
    }
}

func BookDetailHandlerv2(c *gin.Context) {
    var b moduls.Book
    bookID, err := strconv.Atoi(c.Param("id"))
    if err != nil || bookID <= 0 {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Data not found",
        })
        return
    }

    err = moduls.DB.Debug().Where("id = ?", bookID).First(&b).Error
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Error while finding data",
        })
        return
    }

    c.JSON(http.StatusOK, b)
}

func BookUpdateHandler(c *gin.Context) {
    var b moduls.Book

    bookID, err := strconv.Atoi(c.Param("id"))
    if err != nil || bookID <= 0 {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Data not found",
        })
        return
    }

    err = c.ShouldBindJSON(&b)
    if err != nil {
        eMsg := []string{}
        for _, e := range err.(validator.ValidationErrors) {
            msg := fmt.Sprintf("Field error :'%s', condition: '%s'", e.Field(), e.ActualTag())
            eMsg = append(eMsg, msg)
        }

        c.JSON(http.StatusBadRequest, gin.H{
            "error": eMsg,
        })
        return
    }

    // Update sesuai dengan nilai (where) book id
    b.ID = bookID
    b.UpdatedAt = time.Now()
    moduls.DB.Updates(&b)

    // kirim data json perihal data terkait ke client
    err = moduls.DB.First(&b, bookID).Error
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Data not found",
        })
        return
    }

    // return the book data
    c.JSON(http.StatusOK, b)
}

func BookDeleteHandler(c *gin.Context) {
    var b moduls.Book

    bookID, err := strconv.Atoi(c.Param("id"))
    if err != nil || bookID <= 0 {
        c.JSON(http.StatusBadRequest, gin.H{
            "error":"Data not found",
        })
        return
    }

    moduls.DB.Where("deleted_at IS NULL AND id = ?", bookID).First(&b)
    if b.ID != bookID {
        c.JSON(http.StatusOK, gin.H{
            "info": fmt.Sprintf("Book %d already deleted. Operation canceled.", bookID),
        })
        return
    }

    err = moduls.DB.Delete(&b, bookID).Error
    if err != nil {
        c.JSON(http.StatusNoContent, gin.H{
            "error": "Data not found",
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "book": fmt.Sprintf("Book %d has been deleted.", bookID),
    })

    
}
