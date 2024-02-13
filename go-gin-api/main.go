package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID     string `json:"id"`
	Title  string `json:"title"` 
	Author string `json:"author"` 
	Quantity int  `json:"quantity"`
}

var books = []book {
	{ID: "1", Title: "In Search of Good", Author: "Naveen", Quantity: 3},
	{ID: "2", Title: "I Found Bad", Author: "N", Quantity: 1},
	{ID: "3", Title: "Calm after Strom", Author: "NC", Quantity: 5},
}


func getBookById(id string) (*book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}
	return nil, errors.New("book not found")
}

// ********************************************************************

func getBooks(c *gin.Context) {
	// * WARNING: It is recommended to use this only for development purposes 
	// * Since printing pretty JSON is more CPU and bandwidth consuming. Use Context.JSON() instead.
	c.IndentedJSON(http.StatusOK, books)
}

// ********************************************************************

func checkoutBook(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing ID query parameter"})
		return
	}

	book, err := getBookById(id)
	if err != nil {
        c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
        return
    }

	if book.Quantity <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Book not aailable"})
		return
	}

	book.Quantity -= 1
	c.IndentedJSON(http.StatusOK, book)
}

// ********************************************************************

func returnBook(c *gin.Context) {
	id, ok := c.GetQuery("id")
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing ID query parameter"})
		return
	}

	book, err := getBookById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}
	book.Quantity += 1
	
	c.IndentedJSON(http.StatusOK, book)
}

// ********************************************************************

func bookById(c *gin.Context) {
	id := c.Param("id")
	book, err := getBookById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return 
	}
	c.IndentedJSON(http.StatusOK, book)
}

// ********************************************************************

func createBook(c *gin.Context) {
	var newBook book
	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

// ********************************************************************


func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.GET("/books/:id", bookById)
	router.POST("/books", createBook)
	router.PATCH("/checkout", checkoutBook)
	router.PATCH("/return", returnBook)

	router.Run("localhost:8020")

}

