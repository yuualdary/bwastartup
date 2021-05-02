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


func (h *userHandler) LoginUser(c *gin.Context){


	var input user.LoginUserInput

	err := c.ShouldBindJSON(&input)

	//err adalah value dari input ->mendapatkan semua isi dan yang tidak diisi, jika ada tidak diisi masuk ke validasi (err)

	//validasi form
	if err != nil {

		errors := helper.FormatValidationError(err)

		ErrorMessage := gin.H{
			"error" : errors,
		}
		response := helper.APIResponse("Login Failed", http.StatusBadRequest, "error", ErrorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	//validasi hasil inputan
	newLogin, err := h.userService.LoginUser(input)

	if err !=nil {

		ErrorMessage := gin.H{
			"error" : err.Error(),
		}
		response := helper.APIResponse("Failed Login", http.StatusBadRequest, "error", ErrorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := user.FormatUser(newLogin,"tokentoken")


	response := helper.APIResponse("Success Login", http.StatusOK,"success",formatter)
	c.JSON(http.StatusOK,response)


}


func (h *userHandler) CheckEmailIsExist(c *gin.Context){

	//ada input email dari user
	//input email di-mapping ke struct input
	//struct input di passing ke service
	//service akan memanggil repository - is mail exist
	//repository query ke db
	
	var input user.CheckMailInput

	err := c.ShouldBindJSON(&input)

	//err adalah value dari input ->mendapatkan semua isi dan yang tidak diisi, jika ada tidak diisi masuk ke validasi (err)

	//validasi form
	if err != nil {

		errors := helper.FormatValidationError(err)

		ErrorMessage := gin.H{
			"error" : errors,
		}
		response := helper.APIResponse("Email Check Failed", http.StatusBadRequest, "error", ErrorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	CheckMail, err := h.userService.CheckMailUser(input)
	
	if err !=nil {

		ErrorMessage := gin.H{
			"error" : err.Error(),
		}

		response := helper.APIResponse("Email Check Failed", http.StatusBadRequest, "error", ErrorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	//buat balikin ke jsonnya
	data := gin.H{
		"Is_available":CheckMail,
	}

	var metamessage string
	
	if CheckMail{
		metamessage = "Email Available"
	}else{
		metamessage = "Email Has Been Registered"
	}

	response := helper.APIResponse(metamessage, http.StatusBadRequest, "error", data)
	c.JSON(http.StatusBadRequest, response)
	




}
