package handler

import (
	"bwastartup/helper"
	"bwastartup/models"
	"bwastartup/transaction"
	"net/http"

	"github.com/gin-gonic/gin"
)


type TransactionHandler struct {
	TransactionService transaction.Service
}

func NewTransactionHandler(TransactionService transaction.Service) *TransactionHandler{
	return &TransactionHandler{TransactionService}
}

func(h *TransactionHandler) GetTransaction(c *gin.Context){

	var input transaction.GetTransactionInput

	err:= c.ShouldBindUri(&input)

	if err != nil{
		response := helper.APIResponse("Fail Get Bind Data", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	CurrentUser :=c.MustGet("CurrentUser").(models.Users)//get current user
	input.User = CurrentUser

	GetTransaction, err := h.TransactionService.GetTransactionByCampaignID(input)

	if err != nil{
		response := helper.APIResponse("Fail Get Transaction Data", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List Of Transaction Data", http.StatusBadRequest, "success",transaction.FormatTransactions(GetTransaction))
	c.JSON(http.StatusBadRequest, response)
}

func (h *TransactionHandler) GetUserTransaction(c *gin.Context){
	
	CurrentUser :=c.MustGet("CurrentUser").(models.Users)//get current user
	UserID := CurrentUser.ID

	transactions, err := h.TransactionService.GetTransactionByUserID(int(UserID))

	if err !=nil{
		response := helper.APIResponse("Fail Get User Transaction Data", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List Of User Transactions", http.StatusBadRequest, "success",transaction.ListFormatUserTransactions(transactions))
	c.JSON(http.StatusBadRequest, response)



	
}
func(h *TransactionHandler) CreateTransaction(c *gin.Context){

	var input transaction.CreateTransactionInput

	err := c.ShouldBindJSON(&input)
	
	if err != nil{

		
		errors := helper.FormatValidationError(err)

		ErrorMessage := gin.H{
			"error" : errors,
		}
		response := helper.APIResponse("Fail Create Transaction Data", http.StatusBadRequest, "error", ErrorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	CurrentUser :=c.MustGet("CurrentUser").(models.Users)//get current user
	input.User = CurrentUser

	//save data
	SaveTransaction, err := h.TransactionService.CreateTransaction(input)
	if err != nil{
		response := helper.APIResponse("Fail Create Transaction Data", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success Create Transaction Data", http.StatusOK, "success", transaction.FormatPaymentTransaction(SaveTransaction))
	c.JSON(http.StatusOK, response)
}

func (h *TransactionHandler)GetNotification(c *gin.Context){

	var input transaction.TransactionNotificationInput

	err := c.ShouldBindJSON(&input)

	if err != nil{

		response := helper.APIResponse("Fail Get Notification Data", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	err = h.TransactionService.ProcessPayment(input)
	
	if err != nil{

		response := helper.APIResponse("Fail Get Notification Data", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}


	c.JSON(http.StatusOK, input)



}