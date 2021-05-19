package models

import "gorm.io/gorm"

type Transactions struct {
	gorm.Model
	CampaignsID int
	UsersID     int
	Amount      int
	Status      string
	Code        string
	User Users `gorm:"foreignKey:UsersID"`

}