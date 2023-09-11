package main

import (
	"github.com/gin-gonic/gin"
	"github.com/j-jayes/go-backend-test/pkg/api"
)

func main() {
	router := gin.Default()

	router.GET("/books", api.GetBooks)
	router.GET("/books/:id", api.BookById)
	router.POST("/books", api.CreateBook)
	router.PATCH("/checkout", api.CheckoutBook)
	router.PATCH("/return", api.ReturnBook)

	router.Run("localhost:8090")
}
