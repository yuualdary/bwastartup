package campaign

import (
	"bwastartup/models"
	"fmt"
	"strings"
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

type CampaignDetailFormatter struct{

	ID             int    `json:"id"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	ImageUrl       string `json:"image_url"`
	Goal_amount    int    `json:"goal_amount"`
	Current_amount int    `json:"current_amount"`
	UsersID        int    `json:"user_id"`
	Slug string   `json:"slug"`
	Perks []string `json:"perks"`
	User CampaignUserFormatter `json:"user"`
	Images []CampaignImageFormatter `json:"images"`
}

type CampaignUserFormatter struct{

	Name string `json:"name"`
	ImageUrl string `json:"image_url"`
}

type CampaignImageFormatter struct{

	ImageURL string `json:"image_url"`
	IsPrimary bool `json:"is_primary"`
}

func FormatCampaignDetail(campaign models.Campaigns) CampaignDetailFormatter{

	CampaignDetailFormatter := CampaignDetailFormatter{}
	CampaignDetailFormatter.ID = int(campaign.ID)
	CampaignDetailFormatter.UsersID = campaign.UsersID
	CampaignDetailFormatter.Name = campaign.Name
	CampaignDetailFormatter.Description = campaign.Description
	CampaignDetailFormatter.Slug = campaign.Slug

	CampaignDetailFormatter.ImageUrl = ""

	fmt.Println(len(campaign.Campaign_photos))
	if len(campaign.Campaign_photos) > 0 {
		CampaignDetailFormatter.ImageUrl = campaign.Campaign_photos[0].File_name
	}

	var perks []string
	for _, perk := range strings.Split(campaign.Perks, ","){
		perks = append(perks, strings.TrimSpace(perk))//trim space untuk mengapus space pada awal characther
	}
	CampaignDetailFormatter.Perks=perks

	user := campaign.User

	SetCampaignUserFormatter := CampaignUserFormatter{}
	SetCampaignUserFormatter.Name = user.Name
	SetCampaignUserFormatter.ImageUrl = user.Avatar_file_name
	CampaignDetailFormatter.User = SetCampaignUserFormatter

	images := []CampaignImageFormatter{}

	for _, image := range campaign.Campaign_photos{

		CampaignImageFormatter := CampaignImageFormatter{}
		CampaignImageFormatter.ImageURL = image.File_name
		CampaignImageFormatter.IsPrimary = image.Is_primary

		IsPrimary := false 

		if image.Is_primary == true {

			IsPrimary = true
		}

		CampaignImageFormatter.IsPrimary = IsPrimary
		
		images = append(images, CampaignImageFormatter)
		


	}

	CampaignDetailFormatter.Images = images
	return CampaignDetailFormatter
}