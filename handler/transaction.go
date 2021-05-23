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