package api

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/yourproject/pkg/models"
)

var books = []models.Book{
	{ID: "1", Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 2},
	//... other initial books ...
}

func GetBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

//... other handler functions like `BookById`, `CheckoutBook`, `ReturnBook`, `CreateBook`, etc ...

func getBookById(id string) (*models.Book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}

	return nil, errors.New("book not found")
}
