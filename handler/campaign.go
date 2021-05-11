package handler

import (
	"bwastartup/campaign"
	"bwastartup/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)


type CampaignHandler struct {
	CampaignService campaign.Service
}

func NewCampaignHandler(CampaignService campaign.Service)*CampaignHandler{
	return &CampaignHandler{CampaignService}
}


func (h *CampaignHandler) GetCampaign(c *gin.Context){
	UserID, _ := strconv.Atoi(c.Query("user_id"))

	GetUserCampaign, err := h.CampaignService.GetCampaign(UserID)

	if err != nil{
		response := helper.APIResponse("Fail Get Campaign Data", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("List Of Campaign Data", http.StatusBadRequest, "success",campaign.FormatCampaigns(GetUserCampaign))
	c.JSON(http.StatusBadRequest, response)
}

func (h *CampaignHandler) GetDetailCampaign(c *gin.Context){

	var input campaign.GetCampaignDetailInput//isinya id struct untuk ngebaca id brp /:id

	err := c.ShouldBindUri(&input)//isinya id yang sudah di convert cmmiw
	
	if err != nil{
		response := helper.APIResponse("Fail Get Bind Data", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	GetUserCampaign, err := h.CampaignService.GetCampaignByID(input)

	if err != nil{
		response := helper.APIResponse("Fail Get Detail Campaign Data", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}


	response := helper.APIResponse("List Of Detail Campaign Data", http.StatusBadRequest, "success",campaign.FormatCampaignDetail(GetUserCampaign))
	c.JSON(http.StatusBadRequest, response)
}