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

    bookResponse := bookToResponse(book)

    c.JSON(http.StatusOK, gin.H{
        "data": bookResponse, 
    })
}

func (h *bookHandler) ShowBooksHandler(c *gin.Context) {
    books, err := h.service.All()

    if err != nil {
        c.JSON(http.StatusNoContent, gin.H{
            "error":"Error retrieving data",
        })
        return
    }
    var bookResponse []moduls.BookResponse
    for _, b := range books {
        bookResponse = append(bookResponse, bookToResponse(b))
    }

    c.JSON(http.StatusOK, gin.H{
        "data": bookResponse,
    })
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
        bookResponse := bookToResponse(book)
        c.JSON(http.StatusOK, gin.H{
            "data": bookResponse,
        })
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

    bookResponse := bookToResponse(book)

    c.JSON(http.StatusOK, gin.H{
        "data": bookResponse,
    })
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
    b, err := h.service.Update (bookID, bookRequest)

    // kirim data json perihal data terkait ke client
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Data not found",
        })
        return
    }
    // Karena struct B tidak berisi ID, maka perlu ditambahkan secara manual
    b.ID = bookID

    bookResponse := bookToResponse(b)

    // return the book data
    c.JSON(http.StatusOK, gin.H{
        "data": bookResponse,
    })
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

func bookToResponse(b moduls.Book) moduls.BookResponse{
    return moduls.BookResponse{
        ID:          b.ID,
        Title:       b.Title,
        Description: b.Description,
        Price:       b.Price,
        Rating:      b.Rating,
    }
}
