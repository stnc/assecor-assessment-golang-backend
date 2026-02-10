package controller

import (
	"net/http"
	"stncCms/app/services"

	"github.com/gin-gonic/gin"
)

// Options constructor
type Options struct {
	OptionsApp services.OptionsAppInterface
}

const viewPathOptions = "admin/options/"

// InitOptions post controller constructor
func InitOptions(OptionsApp services.OptionsAppInterface) *Options {
	return &Options{
		OptionsApp: OptionsApp,
	}
}

type Book struct {
	Title  string
	Author string
}

// Index list
func (access *Options) Index(c *gin.Context) {

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

	// c.HTML(http.StatusOK, viewPathOptions+"index.html", gin.H{
	// 	"books": books,
	// })
}
