package users

import "github.com/gin-gonic/gin"

type UserController struct {
}

func NewUserController() *UserController {
	return &UserController{}
}

func (uc *UserController) RegisterRoutes(r *gin.Engine) {
	users := r.Group("/users")
	{
		users.GET("/", uc.GetAllUsers)
		users.GET("/:id", uc.GetUser)
		users.POST("", uc.CreateUser)
		users.PUT("/:id", uc.UpdateUser)
		users.DELETE("/:id", uc.DeleteUser)

	}
}

func (uc *UserController) CreateUser(ctx *gin.Context) {

}

func (uc *UserController) GetUser(ctx *gin.Context) {
	ctx.String(200, "FUNCIONOU")
}

func (uc *UserController) GetAllUsers(ctx *gin.Context) {

}

func (uc *UserController) UpdateUser(ctx *gin.Context) {

}

func (uc *UserController) DeleteUser(ctx *gin.Context) {

}
