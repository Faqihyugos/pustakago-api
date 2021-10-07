package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/",func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"name": "Faqih Yugo Susilo",
			"bio": "Software Engineer",
		})
	})

	router.GET("/hello",func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"title": "Hello World",
			"subtitle": "Belajar Golang",
		})
	})


	router.Run()
}