package main

import (
	"log"

	users "github.com/Mateus-R-De-Lima/Library-MVC/internal/users/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	userController := users.NewUserController()
	userController.RegisterRoutes(router)

	if err := router.Run(); err != nil {
		log.Fatal(err)
	}

}
