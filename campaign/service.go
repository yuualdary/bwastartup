package campaign

import "bwastartup/models"

type Service interface {
	GetCampaign(UserID int) ([]models.Campaigns, error)

}

type service struct{
	repository Repository
}

func NewService(repository Repository) *service{
	return &service{repository}
}

func (s *service) GetCampaign(UserID int) ([]models.Campaigns, error){//karena mau nampilin campaign yang sesuai id butuh slice untuk menangkap keseluruhan campaign yang connect ke user id
	
	if UserID != 0 {
		GetCampaignID, err := s.repository.FindByID(UserID)

		if err != nil{
			return GetCampaignID, err
		}

		return GetCampaignID, nil
	}

	GetAllCampaign, err := s.repository.FindAll()

		if err != nil{
			return GetAllCampaign, err
		}

	return GetAllCampaign, nil
}

