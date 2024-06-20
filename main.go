package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	AuthorName string `json:"authorName"`
	Quantity   string `json:"quantity"`
}

var books = []Book{{
	ID: "1", Title: "The Go Programming Language", AuthorName: "Alan A. A. Donovan", Quantity: "10",
}, {
	ID: "2", Title: "Go in Action", AuthorName: "William Kennedy", Quantity: "5",
}, {
	ID: "3", Title: "Go Programming Blueprints", AuthorName: "Mat Ryer", Quantity: "7",
}, {
	ID: "4", Title: "Concurrency in Go", AuthorName: "Katherine Cox-Buday", Quantity: "3",
}, {
	ID: "5", Title: "Introducing Go", AuthorName: "Caleb Doxsey", Quantity: "8",
}}

type getBooksResponse struct {
	Data []Book `json:"data"`
}

func getBooks(c *gin.Context) {
	var response getBooksResponse
	response.Data = books
	c.IndentedJSON(http.StatusOK, response)
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.Run("localhost:8080")
}
