package loans

import (
	"net/http"
	"strconv"

	"github.com/Mateus-R-De-Lima/Library-MVC/internal/loans/models"
	"github.com/gin-gonic/gin"
)

type LoanController struct {
	loanService models.LoanService
}

func NewLoanController(loanService models.LoanService) *LoanController {
	return &LoanController{loanService: loanService}
}

func (lc *LoanController) RegisterRoutes(r *gin.Engine) {
	loans := r.Group("/loans")
	{
		loans.GET("/", lc.GetAllLoans)
		loans.GET("/:id", lc.GetLoan)
		loans.POST("", lc.CreateLoan)
		loans.GET("", lc.GetAllLoans)
		loans.GET("/:id/return", lc.ReturnBook)
	}

	users := r.Group("/loans/users")
	{
		users.GET("/:userId/loans", lc.GetUserLoans)
	}
}

func parseID(ctx *gin.Context, info string) (int64, bool) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid " + info + " ID"})
		return 0, false
	}

	return id, true
}

func (lc *LoanController) CreateLoan(ctx *gin.Context) {
	var request struct {
		BookID int64 `json:"bookID"`
		UserID int64 `json:"userID"`
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	loan, err := lc.loanService.CreateLoan(request.BookID, request.UserID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, loan)
}

func (lc *LoanController) GetLoan(ctx *gin.Context) {
	id, ok := parseID(ctx, "loan")
	if !ok {
		return
	}

	loan, err := lc.loanService.GetLoan(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, loan)
}

func (lc *LoanController) GetAllLoans(ctx *gin.Context) {
	loans, err := lc.loanService.GetAllLoans()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, loans)

}

func (lc *LoanController) GetUserLoans(ctx *gin.Context) {
	userId, ok := parseID(ctx, "user")
	if !ok {
		return
	}

	loan, err := lc.loanService.GetUserLoans(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, loan)
}
func (lc *LoanController) ReturnBook(ctx *gin.Context) {
	id, ok := parseID(ctx, "book")
	if !ok {
		return
	}
	err := lc.loanService.ReturnBook(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}
