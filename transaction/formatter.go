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
type UserTransactionFormatter struct{
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Amount    int    `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
	Campaign CampaignUserFormatter `json:"campaign"`

}
type CampaignUserFormatter struct{

	Name string `json:"name"`
	ImageUrl string `json:"image_url"`
}

func FormatUserTransaction(transaction models.Transactions) UserTransactionFormatter{

	formatter := UserTransactionFormatter{}
	formatter.ID = int(transaction.ID)
	formatter.Name = transaction.User.Name
	formatter.Amount = transaction.Amount
	formatter.CreatedAt = transaction.CreatedAt
	

	CampaignFormatter := CampaignUserFormatter{}
	CampaignFormatter.Name =  transaction.Campaign.Name
	CampaignFormatter.ImageUrl = ""

	if len(transaction.Campaign.Campaign_photos) > 0 {
		CampaignFormatter.ImageUrl = transaction.Campaign.Campaign_photos[0].File_name
	}
	
	formatter.Campaign = CampaignFormatter

	return formatter

}

func ListFormatUserTransactions(transactions []models.Transactions) []UserTransactionFormatter{

	if len(transactions) == 0 {
		return []UserTransactionFormatter{}
	}

	var ListTransactionFormatter []UserTransactionFormatter//single object

	for _, GetTransaction := range transactions{

		formatter := FormatUserTransaction(GetTransaction)//get user transaction formatter
		ListTransactionFormatter = append(ListTransactionFormatter, formatter)//di append ke list []


	}
	return ListTransactionFormatter

}



type PaymentTransactionFormatter struct{
	ID        int    `json:"id"`
	CampaignID int `json:"campaign_id"`
	UserID int `json:"user_id"`
	Amount    int    `json:"amount"`
	Status string `json:"status"`
	Code string `json:"code"`
	PaymentURL string `json:"payment_url"`

}
func FormatPaymentTransaction(transaction models.Transactions) PaymentTransactionFormatter{
	formatter := PaymentTransactionFormatter{}
	formatter.ID = int(transaction.ID)
	formatter.CampaignID = transaction.CampaignsID
	formatter.UserID = transaction.UsersID
	formatter.Amount = transaction.Amount
	formatter.Status = transaction.Status
	formatter.Code = transaction.Code
	formatter.PaymentURL = transaction.PaymentURL

	 return formatter
}