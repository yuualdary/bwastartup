package models

import (
	"gorm.io/gorm"
)


type Campaigns struct {
	gorm.Model
	UsersID int
	Name string
	Campaign_photo []Campaign_photo `gorm:"many2many:user_tasks;"`
	Description string
	Goal_amount int
	Current_amount int
	Perks string
	Backer_count int 
	Slug string
	Transactions []Transactions `gorm:"many2many:user_tasks;"`


}
type CampaignsToPhoto struct{

	CampaignsID int `gorm:"primaryKey"`
	CampaignsToPhoto int `gorm:"primaryKey"`
}

type CampaignsToTransaction struct{
	CampaignsID int `gorm:"primaryKey"`
	TransactionsID int `gorm:"primaryKey"`
	
}