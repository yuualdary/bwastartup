package campaign

import (
	"bwastartup/models"
	"fmt"
)


type CampaignFormatter struct {
	ID             int    `json:"id"`
	UsersID        int    `json:"user_id"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	ImageUrl       string `json:"image_url"`
	Goal_amount    int    `json:"goal_amount"`
	Current_amount int    `json:"current_amount"`
	Slug string   `json:"slug"`

}

func FormatCampaign(campaign models.Campaigns) CampaignFormatter{
	CampaignFormatter := CampaignFormatter{}
	CampaignFormatter.ID = int(campaign.ID)
	CampaignFormatter.UsersID = campaign.UsersID
	CampaignFormatter.Name = campaign.Name
	CampaignFormatter.Description = campaign.Description
	CampaignFormatter.Slug = campaign.Slug
	CampaignFormatter.ImageUrl = ""

	fmt.Println(len(campaign.Campaign_photos))
	if len(campaign.Campaign_photos) > 0 {
		CampaignFormatter.ImageUrl = campaign.Campaign_photos[0].File_name
	}

	return CampaignFormatter

}

func FormatCampaigns(ListCampaigns []models.Campaigns) []CampaignFormatter{

	
	 ListCampaignsFormatter := []CampaignFormatter{}//ini buat default, dmn nilai default dari formatter adalah struct kosong

	for _, campaign := range ListCampaigns {

		CampaignFormatter := FormatCampaign(campaign)
		ListCampaignsFormatter = append(ListCampaignsFormatter, CampaignFormatter)

	}
	return ListCampaignsFormatter
}