package transaction

import "bwastartup/models"

type GetTransactionInput struct {
	ID   int `uri:"id" binding:"required"`
	User models.Users
}

type CreateTransactionInput struct {
	Amount int `json:"amount" binding:"required"`
	CampaignID int `json:"campaign_id" binding:"required"`
	User models.Users
	//ini hanya inputan, kalau default langsung dari service aja
}
