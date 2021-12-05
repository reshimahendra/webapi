package handler

import (
	"fmt"
	"log"
	"net/http"
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
// func SearchBookHandler(c *gin.Context) {
//     // var book moduls.Book
//     param := c.Request.URL.Query()
//
//
//     var qCondition []string
//     var qValue string
//     var sqlCommand string
//     n := ""
//     for i, value := range param {
//         sqlCommand = fmt.Sprintf("%s, %s", sqlCommand, i)
//         qValue = ""
//         if len(value) > 1 {
//             log.Println("Length i more than 1", i)
//             switch i {
//             case "name","description":
//                 for _, n = range value {
//                     qValue = fmt.Sprintf("%s, %q", qValue,n)
//                 }
//                 qCondition = append(qCondition, fmt.Sprintf("%s IN (%s)", i, qValue))
//                 fmt.Println("%s IN (%s)", i, qValue)
//             case "id", "price", "rating":
//                 for _, n = range value {
//                     qValue = fmt.Sprintf("%s, %d", qValue,n)
//                 }
//                 fmt.Println("%s IN (%s)", i, qValue)
//                 qCondition = append(qCondition, fmt.Sprintf("%s IN (%s)", i, qValue))
//             default: log.Println(qCondition)
//                 fmt.Println("%s IN (%s)", i, qValue)
//             }
//         } else {
//             switch i {
//                 case "name": qCondition = append(qCondition, fmt.Sprintf("%s = %q", i, value[0]))
//                 case "id","price","rating": qCondition = append(qCondition, fmt.Sprintf("%s = %d", i, value[0]))
//                 default: log.Println(qCondition)
//             }
//         }
//
//     }
//     log.Println(sqlCommand)
//     log.Println("==================================================")
//
//     // // construct the sqlCommand
//     if len(qCondition) > 1 {
//         for k, l := range qCondition {
//             log.Println("k: %s l: %s", k, l)
//             if k == 0 {
//                 sqlCommand = l
//             } else {
//                 sqlCommand = fmt.Sprintf("%s AND %s", sqlCommand, l)
//             }
//         }
//     } //else {
//     //     sqlCommand = qCondition[0]
//     // }
//     log.Println(param)
// }
