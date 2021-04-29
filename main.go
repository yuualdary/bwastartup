package main

import (
	"bwastartup/config"
	"bwastartup/controllers"
	"bwastartup/handler"
	"bwastartup/user"

	"github.com/gin-gonic/gin"
)

func main(){

	router := gin.Default()
	config.ConnectDatabase()

	userRepository := user.NewRepository(config.DB)
	userService := user.NewService(userRepository)
	// userInput := user.RegisterUserInput{}

	// userInput.Name = "Test Simpan"
	// userInput.Email = "test@test.co"
	// userInput.Occupation = "Tukang Ciduk"
	// userInput.Password = "123456"

	// userService.RegisterUser(userInput)

	// user := models.Users{
	// 	Name: "Test Nama",
	// }
	// fmt.Println(user)
	// userRepository.Save(user)
	userHandler := handler.NewUserHandler(userService)

	v1 := router.Group("/api/v1")
	{
		v1.GET("/user/all", controllers.GetAllUser)
		v1.POST("/user/create", userHandler.RegisterUser)

	}
	router.Run(":3000")

}