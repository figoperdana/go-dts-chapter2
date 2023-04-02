package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"database/sql"
	"strconv"

	"tugas7/config"
)

type Book struct {
	BookID int    `json:"id"` // json tag
	Title  string `json:"title"`   // json tag
	Author string `json:"author"`  // json tag
	Description   string `json:"Description"`    // json tag
}

var books = []Book{}

func CreateBook(ctx *gin.Context) {
	var newBook Book

	if err := ctx.ShouldBindJSON(&newBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	db := config.Connect()
	defer db.Close()

	err := db.QueryRow("INSERT INTO books (title, author, Description) VALUES ($1, $2, $3) RETURNING id", newBook.Title, newBook.Author, newBook.Description).Scan(&newBook.BookID)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"book": newBook})
}

func UpdateBook(ctx *gin.Context) {
	bookID, err := strconv.Atoi(ctx.Param("bookID"))
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var updatedBook Book
	if err := ctx.ShouldBindJSON(&updatedBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	db := config.Connect()
	defer db.Close()

	result, err := db.Exec("UPDATE books SET title=$1, author=$2, Description=$3 WHERE id=$4", updatedBook.Title, updatedBook.Author, updatedBook.Description, bookID)
	if err != nil {
	ctx.AbortWithError(http.StatusInternalServerError, err)
	return
	}

	rowsAffected, err := result.RowsAffected()
	if rowsAffected == 0 {
	ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
		"error_status":  "Data Not Found",
		"error_message": fmt.Sprintf("Book with id %d not found", bookID),
	})
	return
}

ctx.JSON(http.StatusOK, gin.H{
	"message": fmt.Sprintf("Book with id %d has been successfully updated", bookID),
})
}

func GetBook(ctx *gin.Context) {
bookID, err := strconv.Atoi(ctx.Param("bookID"))
if err != nil {
ctx.AbortWithError(http.StatusBadRequest, err)
return
}

db := config.Connect()
defer db.Close()

var bookData Book
err = db.QueryRow("SELECT id, title, author, Description FROM books WHERE id=$1", bookID).Scan(&bookData.BookID, &bookData.Title, &bookData.Author, &bookData.Description)
if err == sql.ErrNoRows {
	ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
		"error_status":  "Data Not Found",
		"error_message": fmt.Sprintf("Book with id %d not found", bookID),
	})
	return
} else if err != nil {
	ctx.AbortWithError(http.StatusInternalServerError, err)
	return
}

ctx.JSON(http.StatusOK, gin.H{
	"book": bookData,
})
}

func DeleteBook(ctx *gin.Context) {
bookID, err := strconv.Atoi(ctx.Param("bookID"))
if err != nil {
ctx.AbortWithError(http.StatusBadRequest, err)
return
}

db := config.Connect()
defer db.Close()

result, err := db.Exec("DELETE FROM books WHERE id=$1", bookID)
if err != nil {
	ctx.AbortWithError(http.StatusInternalServerError, err)
	return
}

rowsAffected, err := result.RowsAffected()
if err != nil {
	ctx.AbortWithError(http.StatusInternalServerError, err)
	return
}

if rowsAffected == 0 {
	ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
		"error_status":  "Data Not Found",
		"error_message": fmt.Sprintf("Book with id %d not found", bookID),
	})
	return
}

ctx.JSON(http.StatusOK, gin.H{
	"message": fmt.Sprintf("Book with id %d has been successfully deleted", bookID),
})
}

func GetAllBooks(ctx *gin.Context) {
db := config.Connect()
defer db.Close()
rows, err := db.Query("SELECT id, title, author, Description FROM books")
if err != nil {
	ctx.AbortWithError(http.StatusInternalServerError, err)
	return
}
defer rows.Close()

var books []Book
for rows.Next() {
	var book Book
	err := rows.Scan(&book.BookID, &book.Title, &book.Author, &book.Description)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	books = append(books, book)
}

if err := rows.Err(); err != nil {
	ctx.AbortWithError(http.StatusInternalServerError, err)
	return
}

ctx.JSON(http.StatusOK, gin.H{
	"books": books,
})
}