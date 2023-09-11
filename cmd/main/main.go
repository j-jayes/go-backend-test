package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/j-jayes/go-backend-test/pkg/api"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // default to localhost if not set
		log.Println("Defaulting to port 8080 as no PORT environment variable was set.")
	}
	router := gin.Default()

	router.GET("/books", api.GetBooks)
	router.GET("/books/:id", api.BookById)
	router.POST("/books", api.CreateBook)
	router.PATCH("/checkout", api.CheckoutBook)
	router.PATCH("/return", api.ReturnBook)

	router.Run("0.0.0.0:" + port)
}
