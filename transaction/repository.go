package transaction

import (
	"bwastartup/models"

	"gorm.io/gorm"
)


type Repository interface {
	GetTransactionByCampaignID(CampaignID int) ([]models.Transactions, error)
}

type repository struct {
	db *gorm.DB
}


func NewRepository(db *gorm.DB) *repository{

	return &repository{db}
}

func (r *repository)GetTransactionByCampaignID(CampaignID int) ([]models.Transactions, error){

	var transactions []models.Transactions

	err := r.db.Preload("User").Where("campaigns_id = ? ",CampaignID).Order("id desc").Find(&transactions).Error

	if err != nil{

		return transactions, err
	}

	return transactions, nil



}
