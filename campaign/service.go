package campaign

import (
	"bwastartup/models"
	"errors"
	"fmt"
	"strconv"

	"github.com/gosimple/slug"
)


type Service interface {
	GetCampaign(UserID int) ([]models.Campaigns, error)
	GetCampaignByID(input GetCampaignDetailInput) (models.Campaigns, error)
	CreateCampaign(input CreateCampaignInput) (models.Campaigns, error)
	UpadateCampaign(CampaignID GetCampaignDetailInput, DetailCampaign CreateCampaignInput) (models.Campaigns, error)

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

func (s *service) GetCampaignByID(input GetCampaignDetailInput) (models.Campaigns, error){

	GetCampaign, err:= s.repository.FindDetailCampaign(input.ID)

	if err != nil{
		return GetCampaign, err
	}

	return GetCampaign, nil

}

func (s *service) CreateCampaign(input CreateCampaignInput) (models.Campaigns, error){
	
	campaign := models.Campaigns{}
	campaign.Name = input.Name
	campaign.Description = input.Description
	campaign.Perks = input.Perks
	campaign.Goal_amount = input.Goal_amount
	campaign.UsersID = int(input.User.ID)
	campaign.Slug = slug.Make(input.Name)

	conv:= strconv.Itoa(int(input.User.ID))
	SlugCandidate := fmt.Sprintf("%s %s", input.Name, conv)//format slugnya nama=campaign-IDuser
	campaign.Slug = slug.Make(SlugCandidate)

	NewCampaign, err := s.repository.Save(campaign)

	if err != nil {
		return NewCampaign, err
	}

	return NewCampaign, nil
}

func (s *service) UpadateCampaign(CampaignID GetCampaignDetailInput, DetailCampaign CreateCampaignInput) (models.Campaigns, error){

	GetCampaign, err:= s.repository.FindDetailCampaign(CampaignID.ID)

	if err != nil{
		return GetCampaign, err
	}
	//get data dari hasil pencarian GetCampaign yg hasilnya []campaign
	//validasi jika campaign yang di db USERID nya = current user ID yang sudah dimasukkan pada handler dengan MustGet
	if GetCampaign.UsersID != int(DetailCampaign.User.ID){

		return GetCampaign, errors.New("You Are Not The Owner")
	}

	GetCampaign.Name = DetailCampaign.Name
	GetCampaign.Description = DetailCampaign.Description
	GetCampaign.Goal_amount = DetailCampaign.Goal_amount
	GetCampaign.Perks = DetailCampaign.Perks
	
	UpdateCampaign, err := s.repository.UpdateCampaign(GetCampaign)

	if err != nil{
		return UpdateCampaign, err
	}
	
	return UpdateCampaign, nil



}

