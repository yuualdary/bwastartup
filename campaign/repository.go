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
	Save(campaign models.Campaigns) (models.Campaigns, error)
	UpdateCampaign(campaign models.Campaigns) (models.Campaigns, error)
	UploadImage(CampaignPhoto models.Campaign_photo) (models.Campaign_photo, error)
	SetAllCampaignPhotoDefault(campaignID int)(bool, error)
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

func (r *repository) Save(campaign models.Campaigns) (models.Campaigns, error){
	
	//create hanya untuk data baru!
	err := r.db.Create(&campaign).Error

	if err != nil {
		return campaign,err
	}

	return campaign, nil
}

func (r *repository) UpdateCampaign(campaign models.Campaigns) (models.Campaigns, error){
	err := r.db.Save(&campaign).Error

	if err != nil{
		return campaign,err 
	}

	return campaign, nil
}

func (r *repository) UploadImage(CampaignPhoto models.Campaign_photo) (models.Campaign_photo, error){

	err:= r.db.Create(&CampaignPhoto).Error

	if err != nil{

		return CampaignPhoto, err
	}

	return CampaignPhoto, nil
}

func (r *repository) SetAllCampaignPhotoDefault(campaignID int)(bool, error){

	//Update campaignphoto, ganti default atau (isprimary) ke false dimana campaign_id sesuai dengan campaign yang dipilih 
	err := r.db.Model(&models.Campaign_photo{}).Where("campaigns_id = ?",campaignID).Update("is_primary",false).Error
	//ini pakai metode gorm, jadi langsung tau model atau table mana yang akan di eksekusi
	
	if err != nil{
		//jika ada data yang salah return error
		return false,err
	}
	return true, nil
}


