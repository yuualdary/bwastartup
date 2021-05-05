package handler

import (
	"bwastartup/auth"
	"bwastartup/helper"
	"bwastartup/models"
	"bwastartup/user"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)


type userHandler struct{
	
	userService user.Service
	AuthService auth.Service
}

func NewUserHandler(userService user.Service, AuthService auth.Service) *userHandler {
	return &userHandler{userService,AuthService}
	
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

	token,err := h.AuthService.GenerateToken(int(newUser.ID))

	if err !=nil {
		response := helper.APIResponse("Fail Add Data", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(newUser,token)


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

	token, err := h.AuthService.GenerateToken(int(newLogin.ID))

	if err !=nil {

		ErrorMessage := gin.H{
			"error" : err.Error(),
		}
		response := helper.APIResponse("Failed Login", http.StatusBadRequest, "error", ErrorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := user.FormatUser(newLogin, token)
	fmt.Println(formatter)

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


func (h *userHandler) UploadAvatar(c *gin.Context){

	//input user
	// simpan gambar pada folder images/
	// service kita panggil repo
	// jwt
	// repo ambil data user yg ID = 1
	// repo update data user ke db, dan simpan file ke folder
	file, err := c.FormFile("avatar")

	if err != nil{
		
		data := gin.H{"is_uploaded" : false}

		response := helper.APIResponse("Failed to upload an image", http.StatusBadRequest,"error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}
	//anggap dapat dari jwt
	CurrentUser :=c.MustGet("CurrentUser").(models.Users)
	userID := int(CurrentUser.ID)

	path := fmt.Sprintf("images/%d-%s",userID,file.Filename)
	//save to patg
	err = c.SaveUploadedFile(file,path)

	if err != nil{
		
		data := gin.H{"is_uploaded" : false}

		response := helper.APIResponse("Failed to upload an image", http.StatusBadRequest,"error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	// get user id dan path untuk disimpan kedalam databaase

	_, err = h.userService.SaveAvatar(userID,path)

	if err != nil{
		data := gin.H{"is_uploaded" : false}

		response := helper.APIResponse("Failed to upload an image", http.StatusBadRequest,"error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	
	}
	data := gin.H{"is_uploaded" : true}

	response := helper.APIResponse("Success to upload an image", http.StatusOK,"error", data)

	c.JSON(http.StatusOK, response)
	


}
