package campaign

import (
	"bwastartup/models"
	"fmt"

	"gorm.io/gorm"
)


type Repository interface {
	FindAll() ([]models.Campaigns, error)
	FindByID(UserID int) ([]models.Campaigns, error)
	FindDetailCampaign(CampaignID int) (models.Campaigns, error)
}


type repository struct{

	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository{

	return &repository{db}
}

func (r *repository) FindAll() ([]models.Campaigns, error){
	var campaigns []models.Campaigns

	err:=r.db.Preload("Campaign_photos").Find(&campaigns).Error
	if err != nil{
		return campaigns, err
	}

	return campaigns, nil
}

func (r *repository)  FindByID(UserID int) ([]models.Campaigns, error){

	var campaigns []models.Campaigns

	err := r.db.Where("users_id = ?", UserID).Preload("Campaign_photos", "campaign_photos.is_primary = 1").Find(&campaigns).Error
	fmt.Println(err)
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}


func (r *repository) FindDetailCampaign(CampaignID int) (models.Campaigns, error){

	var DetailCampaign models.Campaigns

	err := r.db.Preload("User").Preload("Campaign_photos").Where("id = ?", CampaignID).Find(&DetailCampaign).Error

	if err !=nil {
		return DetailCampaign, err
	}

	return DetailCampaign, nil
}



