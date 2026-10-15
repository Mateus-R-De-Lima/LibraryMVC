package loans

import "github.com/gin-gonic/gin"

type LoansController struct {
}

func NewLoanController() *LoansController {
	return &LoansController{}
}

func (uc *LoansController) RegisterRoutes(r *gin.Engine) {
	users := r.Group("/loans")
	{
		users.GET("/", uc.GetAllLoans)
		users.GET("/:id", uc.GetLoan)
		users.POST("", uc.CreateLoan)
		users.PUT("/:id", uc.UpdateLoan)
		users.DELETE("/:id", uc.DeleteLoan)
	}
}

func (uc *LoansController) CreateLoan(ctx *gin.Context) {

}

func (uc *LoansController) GetLoan(ctx *gin.Context) {

}

func (uc *LoansController) GetAllLoans(ctx *gin.Context) {

}

func (uc *LoansController) UpdateLoan(ctx *gin.Context) {

}

func (uc *LoansController) DeleteLoan(ctx *gin.Context) {

}
