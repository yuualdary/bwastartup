package handler

import (
	"bwastartup/campaign"
	"bwastartup/helper"
	"bwastartup/models"
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

func(h *CampaignHandler) CreateCampaign(c *gin.Context){

	var input campaign.CreateCampaignInput

	err := c.ShouldBindJSON(&input)
	
	if err != nil{

		
		errors := helper.FormatValidationError(err)

		ErrorMessage := gin.H{
			"error" : errors,
		}
		response := helper.APIResponse("Fail Create Campaign Data", http.StatusBadRequest, "error", ErrorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	CurrentUser :=c.MustGet("CurrentUser").(models.Users)//get current user
	input.User = CurrentUser

	//save data
	NewCampaign, err := h.CampaignService.CreateCampaign(input)
	if err != nil{
		response := helper.APIResponse("Fail Create Campaign Data", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success Create Campaign Data", http.StatusOK, "success", NewCampaign)
	c.JSON(http.StatusBadRequest, response)

}

func (h *CampaignHandler) UpdateCampaign(c *gin.Context){

	var CampaignID campaign.GetCampaignDetailInput//isinya id struct untuk ngebaca id brp /:id

	err := c.ShouldBindUri(&CampaignID)//isinya id yang sudah di convert cmmiw
	
	if err != nil{


		response := helper.APIResponse("Fail To Update Data", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	

	var CampaignDetail campaign.CreateCampaignInput
	err = c.ShouldBindJSON(&CampaignDetail)

	if err != nil{

		
		errors := helper.FormatValidationError(err)

		ErrorMessage := gin.H{
			"error" : errors,
		}
		response := helper.APIResponse("Fail Get Form Campaign Data", http.StatusBadRequest, "error", ErrorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	CurrentUser :=c.MustGet("CurrentUser").(models.Users)//get current user
	CampaignDetail.User = CurrentUser

	UpdateCampaign, err := h.CampaignService.UpadateCampaign(CampaignID, CampaignDetail)

	if err != nil{
		response := helper.APIResponse("Fail To Update Campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success Update Campaign Data", http.StatusOK, "success", UpdateCampaign)
	c.JSON(http.StatusBadRequest, response)


}