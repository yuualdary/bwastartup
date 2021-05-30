package main

import (
	"bwastartup/auth"
	"bwastartup/campaign"
	"bwastartup/config"
	"bwastartup/handler"
	"bwastartup/middleware"
	"bwastartup/payment"
	"bwastartup/transaction"
	"bwastartup/user"
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main(){

	router := gin.Default()
	router.Use(cors.Default())

	config.ConnectDatabase()

	userRepository := user.NewRepository(config.DB)
	userService := user.NewService(userRepository)
	AuthService := auth.NewService()

	CampaignRepository := campaign.NewRepository(config.DB)
	CampaignService := campaign.NewService(CampaignRepository)
	

	TransactionRepository := transaction.NewRepository(config.DB)
	paymentService := payment.NewService()

	TransactionService := transaction.NewService(TransactionRepository,CampaignRepository,paymentService)

	// user,_ := userService.GetUserById(3)

	// input := transaction.CreateTransactionInput{
	// 	CampaignID: 3 ,
	// 	Amount: 500000,
	// 	User: user,
	// }

	// TransactionService.CreateTransaction(input)
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
	TransactionHandler := handler.NewTransactionHandler(TransactionService)

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
		v1.GET("/user/fetch",middleware.AuthMiddleware(AuthService, userService) ,userHandler.FetchUser)


		v1.GET("/campaigns/all",CampaignHandler.GetCampaign)
		v1.GET("/campaigns/all/:id",CampaignHandler.GetDetailCampaign)
		v1.POST("/campaigns/create",middleware.AuthMiddleware(AuthService, userService),CampaignHandler.CreateCampaign)// kenapa pakai middleware, agar yang mau membuat campaign sudah login dan kita bisa mendapatkan ID nya dia(yang buat campaign)
		v1.PUT("/campaigns/update/:id",middleware.AuthMiddleware(AuthService, userService),CampaignHandler.UpdateCampaign)// kenapa pakai middleware, agar yang mau membuat campaign sudah login dan kita bisa mendapatkan ID nya dia(yang buat campaign)
		v1.POST("/campaigns-images/upload",middleware.AuthMiddleware(AuthService, userService),CampaignHandler.UploadCampaignPhoto)// kenapa pakai middleware, agar yang mau membuat campaign sudah login dan kita bisa mendapatkan ID nya dia(yang buat campaign)


		v1.GET("/campaigns/:id/transactions", middleware.AuthMiddleware(AuthService, userService), TransactionHandler.GetTransaction)
		v1.GET("/transactions/users", middleware.AuthMiddleware(AuthService, userService), TransactionHandler.GetUserTransaction)
		v1.POST("/transactions/create", middleware.AuthMiddleware(AuthService, userService), TransactionHandler.CreateTransaction)
		v1.POST("/transactions/notification", TransactionHandler.GetNotification)

	}
	router.Run(":8000")

}