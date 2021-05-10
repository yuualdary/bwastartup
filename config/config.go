package config

import (
	"bwastartup/models"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


var DB *gorm.DB

func ConnectDatabase(){

	var(
		UserCampaign = models.UserCampaign{}
		UsersTransaction = models.UsersTransaction{}
		// CampaignsToPhoto = models.CampaignsToPhoto{}
		// CampaignsToTransaction = models.CampaignsToTransaction{}
		User = models.Users{}
		Campaign = models.Campaigns{}
		Campaign_photo = models.Campaign_photo{}
		Transaction = models.Transactions{}

	)

	dsn := "root:@tcp(127.0.0.1:3306)/bb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil{
		log.Fatal(err.Error())
	}

	// db.SetupJoinTable(Campaign,"campaign_photos",CampaignsToPhoto)
	db.SetupJoinTable(Campaign,"campaign_photos", Campaign_photo)

	db.SetupJoinTable(User,"campaigns",UserCampaign)
	db.SetupJoinTable(User,"transactions",UsersTransaction)

	db.AutoMigrate(&User,&Campaign,&Campaign_photo,&Transaction)

	fmt.Println("Connecting To Database...")

	DB = db



}