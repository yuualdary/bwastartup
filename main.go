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

		

	}
	router.Run()

}