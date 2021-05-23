package models

import (
	"time"

	"gorm.io/gorm"
)


type Transactions struct {
	gorm.Model
	CampaignsID int
	UsersID     int
	Amount      int
	Status      string
	Code        string
	User Users `gorm:"foreignKey:UsersID"`
	Campaign Campaigns `gorm:"foreignKey:CampaignsID"`
	CreatedAt time.Time
	UpdatedAt time.Time

}