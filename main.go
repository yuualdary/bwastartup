package main

import (
	"bwastartup/auth"
	"bwastartup/campaign"
	"bwastartup/config"
	"bwastartup/handler"
	"bwastartup/middleware"
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

	CampaignRepository := campaign.NewRepository(config.DB)
	CampaignService := campaign.NewService(CampaignRepository)

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
	CampaignHandler := handler.NewCampaignHandler(CampaignService)

	// campaigns, err := campaign.NewRepository(config.DB).FindByID(3)
 

	// fmt.Println("dbug")
	// fmt.Println("dbug")
	// fmt.Println("dbug")

	// for _, campaignlist := range campaigns{

	// 	fmt.Println(campaignlist.Name)

	// 	if len(campaignlist.Campaign_photos) > 0 {
	// 		fmt.Println(campaignlist.Campaign_photos[0].File_name)
	// 	}
	// }


	// token, err := AuthService.ValidateToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyNH0.eMpAU2bDKAbVLcnBxr3CEjf2gNr-aNsJSHMAWJr9QOo")
	// if err != nil{
	// 	fmt.Println("ERROR")
	// 	fmt.Println("ERROR")
	// 	fmt.Println("ERROR")
	// }

	// if token.Valid{
	// 	fmt.Println("VALID")
	// 	fmt.Println("VALID")
	// 	fmt.Println("VALID")
	// }else{
	// 	fmt.Println("INVALID")
	// 	fmt.Println("INVALID")
	// 	fmt.Println("INVALID")
	// }
	// userService.SaveAvatar(2,"images/1-profile.png")
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
	router.Static("/images","./images")//namafolder, nama file db
	v1 := router.Group("/api/v1")
	{
		v1.POST("/user/create", userHandler.RegisterUser)
		v1.POST("/user/login", userHandler.LoginUser)
		v1.POST("/user/checkmail", userHandler.CheckEmailIsExist)
		v1.POST("/user/avatar",middleware.AuthMiddleware(AuthService, userService) ,userHandler.UploadAvatar)


		v1.GET("/campaigns/all",CampaignHandler.GetCampaign)
		v1.GET("/campaigns/all/:id",CampaignHandler.GetDetailCampaign)



		

	}
	router.Run(":8000")

}