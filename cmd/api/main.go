package main

import (
	"log"

	books "github.com/Mateus-R-De-Lima/Library-MVC/internal/books/controllers"
	loans "github.com/Mateus-R-De-Lima/Library-MVC/internal/loans/controllers"
	users "github.com/Mateus-R-De-Lima/Library-MVC/internal/users/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	userController := users.NewUserController()
	booksController := books.NewBooksController()
	leansController := loans.NewLoanController()

	userController.RegisterRoutes(router)
	booksController.RegisterRoutes(router)
	leansController.RegisterRoutes(router)

	if err := router.Run(); err != nil {
		log.Fatal(err)
	}

}
