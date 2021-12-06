package handler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"webapi/moduls"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type bookHandler struct {
    service moduls.Service
}

func NewBookHandler(bookService moduls.Service) *bookHandler {
    return &bookHandler{bookService}
}

func (h *bookHandler) BookPostHandler2(c *gin.Context) {
    var book moduls.BookRequest
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
        "price": book.Price,
        "description": book.Description,
        "rating": book.Rating,
    })
}

func (h *bookHandler) CreateBookHandler(c *gin.Context) {
    var bookRequest moduls.BookRequest

    err := c.ShouldBindJSON(&bookRequest)

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

    book, err := h.service.Create(bookRequest)
    if err != nil {
        log.Println("Error occur while creating new record")
        c.JSON(http.StatusBadRequest, gin.H{
            "error":"Error while creating record",
        })
        return
    }
    c.JSON(http.StatusOK, gin.H{
        "title": book.Title,
        "description": book.Description,
        "price": book.Price,
        "rating": book.Rating,
    })
}

func (h *bookHandler) ShowBooksHandler(c *gin.Context) {
    books, err := h.service.All()

    log.Println("Masuk show all")
    if err != nil {
        c.JSON(http.StatusNoContent, gin.H{
            "error":"Error retrieving data",
        })
        return
    }

    c.JSON(http.StatusOK, books)
}

func (h *bookHandler) SearchBookHandler(c *gin.Context) {
    log.Println(c.Request.URL.Query().Get("id"))
    if c.Request.URL.Query().Has("id") {
        id, _ := strconv.Atoi(c.Request.URL.Query().Get("id"))
        book, err := h.service.FindByID(id)

        // e := moduls.DB.Debug().Where("title = ?", c.Request.URL.Query().Get("title")).Find(&book).Error
        if err!= nil {
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

func (h *bookHandler) BookDetailHandlerv2(c *gin.Context) {
    bookID, err := strconv.Atoi(c.Param("id"))
    if err != nil || bookID <= 0 {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Data not found",
        })
        return
    }

    book, err := h.service.FindByID(bookID) 
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Error while finding data",
        })
        return
    }

    c.JSON(http.StatusOK, book)
}

func (h *bookHandler) BookUpdateHandler(c *gin.Context) {
    var bookRequest moduls.BookRequest

    bookID, err := strconv.Atoi(c.Param("id"))
    if err != nil || bookID <= 0 {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Data not found",
        })
        return
    }

    err = c.ShouldBindJSON(&bookRequest)
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
    // b.ID = bookID
    // b.UpdatedAt = time.Now()
    // moduls.DB.Updates(&b)
    b, err := h.service.Update (bookID, bookRequest)

    // kirim data json perihal data terkait ke client
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Data not found",
        })
        return
    }

    // return the book data
    c.JSON(http.StatusOK, b)
}

func (h *bookHandler) BookDeleteHandler(c *gin.Context) {
    bookID, err := strconv.Atoi(c.Param("id"))
    if err != nil || bookID <= 0 {
        c.JSON(http.StatusBadRequest, gin.H{
            "error":"Data not found",
        })
        return
    }

    _, err = h.service.FindByID(bookID)
    if err != nil {
        c.JSON(http.StatusOK, gin.H{
            "info": fmt.Sprintf("Book %d already deleted. Operation canceled.", bookID),
        })
        return
    }
    // moduls.DB.Where("deleted_at IS NULL AND id = ?", bookID).First(&b)
    // if b.ID != bookID {
    //     c.JSON(http.StatusOK, gin.H{
    //         "info": fmt.Sprintf("Book %d already deleted. Operation canceled.", bookID),
    //     })
    //     return
    // }

    err = h.service.Delete(bookID)
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
