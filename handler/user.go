package handler

import (
	"bwastartup/helper"
	"bwastartup/user"
	"net/http"

	"github.com/gin-gonic/gin"
)


type userHandler struct{
	
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context){
	//tangkap input dari user
	var input user.RegisterUserInput

	

	err := c.ShouldBindJSON(&input)
	//err adalah value dari input ->mendapatkan semua isi dan yang tidak diisi, jika ada tidak diisi masuk ke validasi (err)

	if err != nil {

		errors := helper.FormatValidationError(err)

		ErrorMessage := gin.H{
			"error" : errors,
		}
		response := helper.APIResponse("Fail Add Data", http.StatusBadRequest, "error", ErrorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)

	if err !=nil {
		response := helper.APIResponse("Fail Add Data", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := user.FormatUser(newUser,"tokentoken")


	response := helper.APIResponse("Account Has Been Registered", http.StatusOK,"success",formatter)
	c.JSON(http.StatusOK,response)
}