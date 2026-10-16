package users

import (
	"net/http"
	"strconv"

	"github.com/Mateus-R-De-Lima/Library-MVC/internal/users/models"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService models.UserService
}

func NewUserController(userService models.UserService) *UserController {
	return &UserController{userService: userService}
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

func parseUserID(ctx *gin.Context) (int64, bool) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
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

func (uc *UserController) CreateUser(ctx *gin.Context) {
	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if err := uc.userService.CreateUser(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, user)
}

func (uc *UserController) GetUser(ctx *gin.Context) {
	id, ok := parseUserID(ctx)
	if !ok {
		return
	}

	user, err := uc.userService.GetUser(id)
	if writeError(ctx, http.StatusInternalServerError, err) {
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (uc *UserController) GetAllUsers(ctx *gin.Context) {
	users, err := uc.userService.GetAllUsers()
	if writeError(ctx, http.StatusInternalServerError, err) {
		return
	}

	ctx.JSON(http.StatusOK, users)
}

func (uc *UserController) UpdateUser(ctx *gin.Context) {
	id, ok := parseUserID(ctx)
	if !ok {
		return
	}

	var user models.User

	if err := uc.userService.UpdateUser(id, &user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, nil)
}

func (uc *UserController) DeleteUser(ctx *gin.Context) {
	id, ok := parseUserID(ctx)
	if !ok {
		return
	}

	if err := uc.userService.DeleteUser(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, nil)
}
