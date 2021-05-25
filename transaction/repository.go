package transaction

import (
	"bwastartup/models"

	"gorm.io/gorm"
)


type Repository interface {
	GetTransactionByCampaignID(CampaignID int) ([]models.Transactions, error)
	GetTransactionByUserID(UserID int) ([]models.Transactions, error)
	CreateTransaction(transaction models.Transactions)(models.Transactions, error)
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

func (r *repository) GetTransactionByUserID(UserID int) ([]models.Transactions, error){

	var transactions []models.Transactions

	err := r.db.Preload("User").Preload("Campaign.Campaign_photos","campaign_photos.is_primary = 1").Where("users_id = ?",UserID).Order("id desc").Find(&transactions).Error
	if err != nil{
		return transactions,err
	}

	return transactions,nil
}

func (r *repository)CreateTransaction(transaction models.Transactions)(models.Transactions, error){
	err := r.db.Create(&transaction).Error

	if err != nil{
		return transaction, err
	}

	return transaction, nil
}