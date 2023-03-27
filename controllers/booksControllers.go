package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Book struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

var BookDatas []Book

func CreateBooks(ctx *gin.Context) {
	var newBook Book

	if err := ctx.ShouldBindJSON(&newBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if BookDatas != nil {
		newBook.ID = BookDatas[len(BookDatas)-1].ID + 1
	} else {
		newBook.ID = 1
	}

	BookDatas = append(BookDatas, newBook)

	ctx.JSON(http.StatusCreated, "Created")
}

func GetAllBooks(ctx *gin.Context) {

	if BookDatas != nil {
		ctx.JSON(http.StatusOK, BookDatas)
	} else {
		ctx.JSON(http.StatusOK, []string{})
	}

}

func GetBookById(ctx *gin.Context) {
	bookID := ctx.Param("bookID")
	condition := false
	var bookData Book

	idBook, err := strconv.Atoi(bookID)
	if err != nil {
		log.Println("Invalid convert id")
		return
	}

	for i, v := range BookDatas {
		if idBook == v.ID {
			condition = true
			bookData = BookDatas[i]
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status_code": http.StatusNotFound,
			"message":     fmt.Sprintf("book with id %v not found", idBook),
		})
		return
	}

	ctx.JSON(http.StatusOK, bookData)
}

func UpdateBooks(ctx *gin.Context) {
	bookID := ctx.Param("bookID")
	condition := false
	var updateBook Book

	idBook, err := strconv.Atoi(bookID)
	if err != nil {
		log.Println("Invalid convert ID")
		return
	}

	if err := ctx.ShouldBindJSON(&updateBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	for i, v := range BookDatas {
		if idBook == v.ID {
			condition = true
			BookDatas[i] = updateBook
			BookDatas[i].ID = idBook
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status_code": http.StatusNotFound,
			"message":     fmt.Sprintf("book with id %v not found", idBook),
		})
		return
	}

	ctx.JSON(http.StatusOK, "Updated")
}

func DeleteBook(ctx *gin.Context) {
	bookID := ctx.Param("bookID")
	condition := false
	var bookIndex int

	idBook, err := strconv.Atoi(bookID)
	if err != nil {
		log.Println("Invalid convert id")
		return
	}

	for i, v := range BookDatas {
		if idBook == v.ID {
			condition = true
			bookIndex = i
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status_code": http.StatusNotFound,
			"message":     fmt.Sprintf("book with id %v not found", idBook),
		})
		return
	}

	copy(BookDatas[bookIndex:], BookDatas[bookIndex+1:])
	BookDatas[len(BookDatas)-1] = Book{}
	BookDatas = BookDatas[:len(BookDatas)-1]

	ctx.JSON(http.StatusOK, "Deleted")
}
