package main

import (
	"webapi/moduls"
	"webapi/handler"

	"github.com/gin-gonic/gin"
)

func main() {
    moduls.DB_Conn()
    router := gin.Default()
    router.SetTrustedProxies([]string{"192.168.43.26"})

    // membuat api versioning dengan nama v1 (api version 1)
    v1 := router.Group("/v1")

    v1.GET("/hello", handler.HelloHandler) 
    v1.GET("/book/:id/:title", handler.BookDetailHandler)
    v1.POST("/books", handler.BookPostHandler)
    // v1.GET("/search", handler.SearchQueryHandler)
    v1.GET("/", handler.RootHandler)

    v2 := router.Group("/v2")
    v2.POST("/books", handler.BookPostHandler2)
    v2.POST("/books2", handler.CreateBookHandler)
    v2.GET("/books", handler.ShowBooksHandler)
    v2.GET("/book/:id", handler.BookDetailHandlerv2)
    v2.PUT("/book/:id", handler.BookUpdateHandler)
    v2.DELETE("/book/:id", handler.BookDeleteHandler)
    v2.GET("/book", handler.SearchBookHandler)

    router.Run(":3888")
}

