package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	BookID string `json:"id"` // json tag
	Title  string `json:"title"`   // json tag
	Author string `json:"author"`  // json tag
	Desc   string `json:"desc"`    // json tag
}

var books = []Book{}

func CreateBook(ctx *gin.Context) {
	var newBook Book

	if err := ctx.ShouldBindJSON(&newBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return 
	}
	newBook.BookID = fmt.Sprintf("%d", len(books)+1)
	books = append(books, newBook)

	ctx.JSON(http.StatusCreated, gin.H{"book": newBook,
		})
}

func UpdateBook (ctx *gin.Context) {
	bookID := ctx.Param("bookID")
	condition := false
	var updatedBook Book
	
	if err := ctx.ShouldBindJSON(&updatedBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return 
	}
	
	for i, book := range books {
		if bookID == book.BookID {
			condition = true
			books[i] = updatedBook
			books[i].BookID = bookID
			break
		}
	}
	
	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status": "Data Not Found",
			"error_message": fmt.Sprintf("Book with id %s not found", bookID),
	})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message" : fmt.Sprintf("Book with id %s has been successfully updated", bookID),
	})
}

func GetBook (ctx *gin.Context) {
	bookID := ctx.Param("bookID")
	condition := false
	var bookData Book
	
	for i, book := range books {
		if bookID == book.BookID {
			condition = true
			bookData = books[i]
			break
		}
	}
	
	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status": "Data Not Found",
			"error_message": fmt.Sprintf("Book with id %s not found", bookID),
	})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"book" : bookData,
	})
}

func DeleteBook (ctx *gin.Context) {
	bookID := ctx.Param("bookID")
	condition := false
	var bookIndex int
	
	for i, book := range books {
		if bookID == book.BookID {
			condition = true
			bookIndex = i
			break
		}
	}
	
	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status": "Data Not Found",
			"error_message": fmt.Sprintf("Book with id %s not found", bookID),
	})
		return
	}

	copy(books[bookIndex:], books[bookIndex+1:])
	books[len(books)-1] = Book{}
	books = books[:len(books)-1]

	ctx.JSON(http.StatusOK, gin.H{
		"message" : fmt.Sprintf("Book with id %s has been successfully deleted", bookID),
	})
}

func GetAllBooks (ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"books" : books,
	})
}

