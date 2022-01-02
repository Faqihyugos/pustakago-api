package main

import (
	"github.com/Faqihyugos/pustakago-api/book"
	"github.com/Faqihyugos/pustakago-api/handler"
	"github.com/gin-gonic/gin"
	mysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {
	// koneksi  db
	dsn := "root:Secret@tcp(127.0.0.1:3306)/pustaka-api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Do connection error")
	}
	db.AutoMigrate(&book.Book{})

	bookRepository := book.NewRepository(db)

	book := book.Book{
		Title:       "$100 Startup",
		Description: "Good book",
		Price:       150000,
		Rating:      5,
	}

	bookRepository.Crete(book)

	// router
	router := gin.Default()
	v1 := router.Group("/v1")
	v1.GET("/", handler.RootHandler)
	v1.GET("/hello", handler.HelloHandler)
	v1.GET("/books/:id/:title", handler.BooksHandler)
	v1.GET("/query", handler.QueryHandler)
	v1.POST("/books", handler.PostBooksHandler)

	router.Run(":8888")
}
