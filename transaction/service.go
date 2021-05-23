package transaction

import (
	"bwastartup/campaign"
	"bwastartup/models"
	"errors"
)


type Service interface {
	GetTransactionByCampaignID(CampaignID GetTransactionInput) ([]models.Transactions, error)
	GetTransactionByUserID(UserID int) ([]models.Transactions, error)
}

type service struct {
	repository Repository
	campaignRepository campaign.Repository
}

func NewService(repository Repository, 	campaignRepository campaign.Repository) *service {
	return &service{repository,campaignRepository}
}


func (s* service)GetTransactionByCampaignID(CampaignID GetTransactionInput) ([]models.Transactions, error){

	Campaign, err := s.campaignRepository.FindDetailCampaign(CampaignID.ID)
	if err != nil{
		return []models.Transactions{}, err
	}
	//yang bisa akses hanya pemilik dari campaign tersebut (bukan transaksinya !)
	//get detail campaign dari ID yang di tuju cari detail dari ID tersebut , 
	//kemudian, temukan ID nya
	//dari user id yang didapat dari detail campaign, bandingkan 
	//apakah detail campaign.userid pada db sesuai atau tidak dengan id yang sekarang sedang login 
	if uint(Campaign.UsersID) != CampaignID.User.ID{

		return []models.Transactions{}, errors.New("Not An owner")

	}

	GetTransaction, err := s.repository.GetTransactionByCampaignID(CampaignID.ID)

	if err != nil{
		return GetTransaction, err
	}

	return GetTransaction, nil
	
}

func (s *service)GetTransactionByUserID(UserID int) ([]models.Transactions, error){

	User, err := s.repository.GetTransactionByUserID(UserID)

	if err !=nil{
		return User,err
	}
	return User, nil





}


