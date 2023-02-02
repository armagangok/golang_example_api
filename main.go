package main


import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	
)

type Book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []Book{
	{ID: "1", Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 2},
	{ID: "2", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 5},
	{ID: "3", Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 6},
}

func fetchBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func bookById(c *gin.Context) {
	id := c.Param("id")
	book, err := fetchBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found."})
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

func fetchBookById(id string) (*Book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}

	return nil, errors.New("book not found")
}

// func fetchBookByID(id string)(*Book, error) {
// 	for i,b := range books {
// 		b.ID == id {
// 			 return &books[i], nil
// 		}

// 	}

// 	return nil, errors.New("book not found")
// }

func storeBooks(c *gin.Context) {
	var newBook Book
	if errorOccured := c.BindJSON(&newBook); errorOccured != nil {
		return
	}

	books = append(books, newBook)

	c.IndentedJSON(http.StatusCreated, books)
}

func router() {
	router := gin.Default()
	router.GET("/books", fetchBooks)
	router.GET("/books/:id", bookById)
	router.POST("/books", storeBooks)
	router.Run("localhost:8080")

}


func main() {
	router()
}
