package handler

import (
	"bwastartup/campaign"
	"bwastartup/helper"
	"bwastartup/models"
	"fmt"
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
	err = c.ShouldBindJSON(&CampaignDetail)// json from input

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

	UpdateCampaign, err := h.CampaignService.UpdateCampaign(CampaignID, CampaignDetail)

	if err != nil{
		response := helper.APIResponse("Fail To Update Campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success Update Campaign Data", http.StatusOK, "success", UpdateCampaign)
	c.JSON(http.StatusBadRequest, response)


}

func (h *CampaignHandler) UploadCampaignPhoto(c *gin.Context){

	var GetCampaign campaign.CreateCampaignPhotoInput

	err := c.ShouldBind(&GetCampaign)

	if err != nil{
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{
			"errors": errors,
		}
		response := helper.APIResponse("Fail To Upload Data", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	file, err := c.FormFile("file")

	if err != nil{
		
		Data := gin.H{
			"is_uploaded" : false,
		}
		response := helper.APIResponse("Fail To Upload Data", http.StatusBadRequest, "error", Data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	CurrentUser :=c.MustGet("CurrentUser").(models.Users)//get current user
	GetCampaign.User = CurrentUser
	GetUserID := CurrentUser.ID

	path := fmt.Sprintf("campaignphotos/%d-%s", GetUserID,file.Filename)

	err = c.SaveUploadedFile(file, path)//save upload file ngambil dari gin

	if err != nil{
		Data := gin.H{
			"is_uploaded" : false,
		}
		response := helper.APIResponse("Fail To Upload Data", http.StatusBadRequest, "error", Data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.CampaignService.CreateCampaignPhoto(GetCampaign, path)
	if err != nil{
		Data := gin.H{
			"is_uploaded" : false,
		}
		response := helper.APIResponse("Fail To Upload Data", http.StatusBadRequest, "error", Data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	Data := gin.H{"is_uploaded" : true,}
	response := helper.APIResponse("Success Upload Data", http.StatusBadRequest, "error", Data)
	c.JSON(http.StatusBadRequest, response)








}