package handlers

import (
	"net/http"
	"nhannguyen/gorest/internal/routes/v1/models"

	"github.com/gin-gonic/gin"
)

var books = []models.Book{
	{ID: 1, Title: "1984", Author: "George Orwell"},
	{ID: 2, Title: "Brave New World", Author: "Aldous Huxley"},
}

// GetBooks godoc
// @Summary Get all books
// @Description get books
// @Tags Books
// @Produce json
// @Success 200 {object} models.BookResponse
// @Router /api/v1/books [get]
func GetBooks(c *gin.Context) {
	data := models.BookResponse{
		Success: true,
		Data:    books,
	}
	c.JSON(http.StatusOK, data)
}

// CreateBook godoc
// @Summary Create a book
// @Description create by JSON book
// @Tags Books
// @Accept json
// @Produce json
// @Param book body models.Book true "Add book"
// @Success 200 {object} models.Book
// @Router /api/v1/books [post]
func CreateBook(c *gin.Context) {
	var newBook models.Book
	if err := c.ShouldBindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newBook.ID = len(books) + 1
	books = append(books, newBook)
	c.JSON(http.StatusOK, newBook)
}
