package main

import (
	"fmt"
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
		log.Fatal("Do conection error")
	}
	db.AutoMigrate(&book.Book{})

	//CRUD
	// Create
	//book := book.Book{}
	//book.Title = "Atomic Habbits"
	//book.Price = 150000
	//book.Rating = 5
	//book.Description = "Buku self development tentang membangu kebiasaan baik dan menghilangkan kebiasaan buruk"
	//
	//err = db.Create(&book).Error
	//if err != nil {
	//	fmt.Println("===========================")
	//	fmt.Println("Error creating book record")
	//	fmt.Println("===========================")
	//}

	//========
	// Read
	//========

	var book book.Book

	err = db.Debug().Where("id = ?", 1).First(&book).Error
	if err != nil {
		fmt.Println("===========================")
		fmt.Println("Error finding book record")
		fmt.Println("===========================")
	}

	book.Title = "Design Thinking 2 edition"
	err = db.Save(&book).Error
	if err != nil {
		fmt.Println("===========================")
		fmt.Println("Error updating book record")
		fmt.Println("===========================")
	}

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
