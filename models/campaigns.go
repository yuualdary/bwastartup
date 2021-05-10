package models

import (
	"gorm.io/gorm"
)


type Campaigns struct {
	gorm.Model
	UsersID int
	Name string
	Campaign_photos []Campaign_photo //`gorm:"foreginKey:CampaignsID;"`
	Description string
	Goal_amount int
	Current_amount int
	Perks string
	Backer_count int 
	Slug string
	Transactions []Transactions //`gorm:"many2many:campaign_transaction;"`


}


// type CampaignsToPhoto struct{

// 	CampaignsID int `gorm:"primaryKey"`
// 	CampaignsToPhoto int `gorm:"primaryKey"`
// }

// type CampaignsToTransaction struct{
// 	CampaignsID int `gorm:"primaryKey"`
// 	TransactionsID int `gorm:"primaryKey"`
	
// }