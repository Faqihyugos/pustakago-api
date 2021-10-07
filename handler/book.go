package handler

import (
	"fmt"
	"net/http"

	"github.com/Faqihyugos/pustakago-api/book"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// handler function
func RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "Faqih Yugo Susilo",
		"bio": "Software Engineer",
	})
}

func HelloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"title": "Hello World",
		"subtitle": "Belajar Golang",
	})
}

// handler func parameter/path variable
func BooksHandler(c *gin.Context) {
	id := c.Param("id")
	title := c.Param("title")

	c.JSON(http.StatusOK, gin.H{"id": id, "title": title})
}

// handler query
func QueryHandler(c *gin.Context) {
	title := c.Query("title")
	price := c.Query("price")

	c.JSON(http.StatusOK, gin.H{"title": title, "price": price})
}



// post books
func PostBooksHandler(c *gin.Context) {
	// title, price
	var bookInput book.BookInput
	
	err := c.ShouldBindJSON(&bookInput)
	if err != nil {

		var errorMessages []string
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage) 
		}
		
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"title": bookInput.Title,
		"price": bookInput.Price,
	})
}