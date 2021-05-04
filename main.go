package main

import (
	"bwastartup/auth"
	"bwastartup/config"
	"bwastartup/controllers"
	"bwastartup/handler"
	"bwastartup/user"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main(){

	router := gin.Default()
	config.ConnectDatabase()

	userRepository := user.NewRepository(config.DB)
	userService := user.NewService(userRepository)
	AuthService := auth.NewService()

	fmt.Println(AuthService.GenerateToken(1001))
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
	userHandler := handler.NewUserHandler(userService, AuthService)
	userService.SaveAvatar(2,"images/1-profile.png")
	// login, err := userRepository.FindByEmail("email@domain.com")

	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	// if login.ID == 0{
	// 	fmt.Println("User Not Found")
	// }else{
	// 	fmt.Println(login.Name)
	// }

	// input := user.LoginUserInput{
	// 	Email : "email@domain.com",
	// 	Password_hash: "asa",
	// }

	// user, err := userService.LoginUser(input)
	// if err != nil{
	// 	fmt.Println("Something Went Wrong")
	// 	fmt.Println(err.Error())
	// }

	// fmt.Println(user.Name)
	// fmt.Println(user.Email)
	v1 := router.Group("/api/v1")
	{
		v1.GET("/user/all", controllers.GetAllUser)
		v1.POST("/user/create", userHandler.RegisterUser)
		v1.POST("/user/login", userHandler.LoginUser)
		v1.POST("/user/checkmail", userHandler.CheckEmailIsExist)
		v1.POST("/user/avatar", userHandler.UploadAvatar)
		


		

	}
	router.Run()

}