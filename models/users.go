package models

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	Name             string
	Email            string
	Occupation string
	Password_hash    string
	Avatar_file_name string
	Role             string
	Token            string
	Campaigns []Campaigns `gorm:"many2many:user_tasks;"`
	Transactions []Transactions `gorm:"many2many:user_tasks;"`

}

type UserCampaign struct{
	UsersID int `gorm:"primaryKey"`
	CampaignsID int `gorm:"primaryKey"`

}

type UsersTransaction struct{
	UsersID int `gorm:"primaryKey"`
	TransactionsID int `gorm:"primaryKey"`

}

