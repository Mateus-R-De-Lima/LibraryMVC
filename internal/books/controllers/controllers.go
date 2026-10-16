package books

import (
	"net/http"
	"strconv"

	"github.com/Mateus-R-De-Lima/Library-MVC/internal/books/models"
	"github.com/gin-gonic/gin"
)

type BooksController struct {
	bookService models.BookService
}

func NewBooksController(bookService models.BookService) *BooksController {
	return &BooksController{bookService: bookService}
}

func (bc *BooksController) RegisterRoutes(r *gin.Engine) {
	books := r.Group("/books")
	{
		books.GET("/", bc.GetAllBooks)
		books.GET("/:id", bc.GetBook)
		books.POST("", bc.CreateBook)
		books.PUT("/:id", bc.UpdateBook)
		books.DELETE("/:id", bc.DeleteBook)
	}
}

func parseBookID(ctx *gin.Context) (int64, bool) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return 0, false
	}
	return id, true
}

func writeError(ctx *gin.Context, statusCode int, err error) bool {
	if err == nil {
		return false
	}
	ctx.JSON(statusCode, gin.H{"error": err.Error()})
	return true
}

func (bc *BooksController) CreateBook(ctx *gin.Context) {
	var book models.Book

	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if err := bc.bookService.CreateBook(&book); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, book)
}

func (bc *BooksController) GetBook(ctx *gin.Context) {
	id, ok := parseBookID(ctx)
	if !ok {
		return
	}
	book, err := bc.bookService.GetBook(id)
	if writeError(ctx, http.StatusInternalServerError, err) {
		return
	}

	ctx.JSON(http.StatusOK, book)
}

func (bc *BooksController) GetAllBooks(ctx *gin.Context) {
	books, err := bc.bookService.GetAllBooks()
	if writeError(ctx, http.StatusInternalServerError, err) {
		return
	}

	ctx.JSON(http.StatusOK, books)
}

func (bc *BooksController) UpdateBook(ctx *gin.Context) {
	id, ok := parseBookID(ctx)
	if !ok {
		return
	}
	var book models.Book
	if err := bc.bookService.UpdateBook(id, &book); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, nil)
}

func (bc *BooksController) DeleteBook(ctx *gin.Context) {
	id, ok := parseBookID(ctx)
	if !ok {
		return
	}

	if err := bc.bookService.DeleteBook(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, nil)

}
