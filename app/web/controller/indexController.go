package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const viewPathIndex = "admin/index/"

// Index all list f
func Index(c *gin.Context) {
	books := make([]Book, 0)
	books = append(books, Book{
		Title:  "Title 1",
		Author: "Author 1",
	})
	books = append(books, Book{
		Title:  "Title 2",
		Author: "Author 2",
	})
	c.JSON(
		http.StatusOK,
		books,
	)
}

//OptionsDefault all list f
