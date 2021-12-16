package main

import (
	"webapi/handler"
	"webapi/moduls"
    "webapi/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
    gin.SetMode(gin.ReleaseMode)

    db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        panic("Failed to connect to DB")
    }

    bookRepo    := moduls.NewRepo(db)
    bookSvc     := moduls.NewService(bookRepo)
    bookHandler := handler.NewBookHandler(bookSvc)

    router := gin.Default()
    router.Use(middleware.CORSMiddleware())
    router.SetTrustedProxies([]string{"192.168.43.26"})

    // membuat api versioning dengan nama v1 (api version 1)
    v1 := router.Group("/v1")

    v1.GET("/hello", handler.HelloHandler) 
    v1.GET("/book/:id/:title", handler.BookDetailHandler)
    v1.POST("/books", handler.BookPostHandler)
    // v1.GET("/search", handler.SearchQueryHandler)
    v1.GET("/", handler.RootHandler)

    v2 := router.Group("/v2")
    v2.POST("/book-add", bookHandler.CreateBookHandler)
    v2.GET("/books", bookHandler.ShowBooksHandler)
    v2.GET("/book/:id", bookHandler.BookDetailHandlerv2)
    v2.PUT("/book/:id", bookHandler.BookUpdateHandler)
    v2.DELETE("/book/:id", bookHandler.BookDeleteHandler)
    v2.GET("/book", bookHandler.SearchBookHandler)

    router.Run(":3888")
}

