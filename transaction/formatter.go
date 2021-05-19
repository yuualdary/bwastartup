package transaction

import (
	"bwastartup/models"
	"time"
)


type TransactionFormatter struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Amount    int    `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

func FormatTransaction(transaction models.Transactions) TransactionFormatter{
	formatter := TransactionFormatter{}
	formatter.ID = int(transaction.ID)
	formatter.Name = transaction.User.Name
	formatter.Amount = transaction.Amount
	formatter.CreatedAt = transaction.CreatedAt
	 return formatter
}

func FormatTransactions(ListTransaction []models.Transactions) []TransactionFormatter{

	if len(ListTransaction) == 0 {
		return []TransactionFormatter{}
	}

	var ListTransactionFormatter []TransactionFormatter

	for _, transaction := range ListTransaction{
		formatter := FormatTransaction(transaction)
		ListTransactionFormatter = append(ListTransactionFormatter, formatter)
	}

	return ListTransactionFormatter
}