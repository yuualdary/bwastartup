package transaction

import "bwastartup/models"

type GetTransactionInput struct {
	ID   int `uri:"id" binding:"required"`
	User models.Users
}
