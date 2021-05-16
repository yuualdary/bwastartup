package campaign

import "bwastartup/models"

type GetCampaignDetailInput struct {
	ID int `uri:"id" binding:"required"`
}
type CreateCampaignInput struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Goal_amount int    `json:"goal_amount" binding:"required"`
	Perks       string `json:"perks" binding:"required"`
	User        models.Users
}