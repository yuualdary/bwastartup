package models

import "gorm.io/gorm"

type Campaign_photo struct {
	gorm.Model
	CampaignsID int
	File_name string
	Is_primary bool

}

