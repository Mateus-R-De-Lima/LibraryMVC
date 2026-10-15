package books

import "github.com/gin-gonic/gin"

type BooksController struct {
}

func NewBooksController() *BooksController {
	return &BooksController{}
}

func (uc *BooksController) RegisterRoutes(r *gin.Engine) {
	users := r.Group("/users")
	{
		users.GET("/", uc.GetAllBooks)
		users.GET("/:id", uc.GetBook)
		users.POST("", uc.CreateBook)
		users.PUT("/:id", uc.UpdateBook)
		users.DELETE("/:id", uc.DeleteBook)
	}
}

func (uc *BooksController) CreateBook(ctx *gin.Context) {

}

func (uc *BooksController) GetBook(ctx *gin.Context) {

}

func (uc *BooksController) GetAllBooks(ctx *gin.Context) {

}

func (uc *BooksController) UpdateBook(ctx *gin.Context) {

}

func (uc *BooksController) DeleteBook(ctx *gin.Context) {

}
