package transaction

import (
	"bwastartup/GeneratorNumber"
	"bwastartup/campaign"
	"bwastartup/models"
	"bwastartup/payment"
	"strconv"

	"errors"
)


type Service interface {
	GetTransactionByCampaignID(CampaignID GetTransactionInput) ([]models.Transactions, error)
	GetTransactionByUserID(UserID int) ([]models.Transactions, error)
	CreateTransaction(input CreateTransactionInput) (models.Transactions, error)
	ProcessPayment(input TransactionNotificationInput) error
}

type service struct {
	repository Repository
	campaignRepository campaign.Repository
	paymentService payment.Service
	
}

func NewService(repository Repository, 	campaignRepository campaign.Repository, paymentService payment.Service) *service {
	return &service{repository,campaignRepository,paymentService}
}


func (s* service)GetTransactionByCampaignID(CampaignID GetTransactionInput) ([]models.Transactions, error){

	Campaign, err := s.campaignRepository.FindDetailCampaign(CampaignID.ID)
	if err != nil{
		return []models.Transactions{}, err
	}
	//yang bisa akses hanya pemilik dari campaign tersebut (bukan transaksinya !)
	//get detail campaign dari ID yang di tuju cari detail dari ID tersebut , 
	//kemudian, temukan ID nya
	//dari user id yang didapat dari detail campaign, bandingkan 
	//apakah detail campaign.userid pada db sesuai atau tidak dengan id yang sekarang sedang login 
	if uint(Campaign.UsersID) != CampaignID.User.ID{

		return []models.Transactions{}, errors.New("Not An owner")

	}

	GetTransaction, err := s.repository.GetTransactionByCampaignID(CampaignID.ID)

	if err != nil{
		return GetTransaction, err
	}

	return GetTransaction, nil
	
}

func (s *service)GetTransactionByUserID(UserID int) ([]models.Transactions, error){

	User, err := s.repository.GetTransactionByUserID(UserID)

	if err !=nil{
		return User,err
	}
	return User, nil


}

func(s * service)CreateTransaction(input CreateTransactionInput) (models.Transactions, error){

	transaction := models.Transactions{}
	transaction.Amount = input.Amount
	transaction.CampaignsID= input.CampaignID
	transaction.UsersID= int(input.User.ID)
	transaction.Code,_ = GeneratorNumber.NewUUID()
	transaction.Status ="pending"

	SaveTransaction, err := s.repository.CreateTransaction(transaction)

	if err != nil {

		return SaveTransaction,err
	}
	paymentTransaction:= payment.Transaction{
		ID: int(SaveTransaction.ID),
		Amount : SaveTransaction.Amount,
	}
	paymentURL, err := s.paymentService.GetPaymentURL(paymentTransaction, input.User)

	if err !=nil{
		return SaveTransaction, err
	}

	SaveTransaction.PaymentURL = paymentURL

	SaveTransaction,err = s.repository.UpdateTransaction(SaveTransaction)//diupdate untuk dapat data transaksinya

	if err !=nil{

		return SaveTransaction,err
	}
	return SaveTransaction,nil
}

func (s *service)  ProcessPayment(input TransactionNotificationInput) error{

    transaction_id,_:= strconv.Atoi(input.OrderID)

    transaction, err := s.repository.GetTransactionByID(transaction_id)
     if err !=nil{
         return err
     }

     if input.PaymentType=="credit_card" &&input.TransactionStatus =="capture" &&input.FraudStatus=="accept"{

        transaction.Status = "paid"
     } else if input.TransactionStatus == "settlement"{
        transaction.Status = "paid"
     }else if input.TransactionStatus == "deny"|| input.TransactionStatus == "expire" || input.TransactionStatus == "cancel"{
         transaction.Status="cancelled"
     }

     SaveTransaction, err := s.repository.UpdateTransaction(transaction)

     if err != nil {

        return  err
     }
     campaign, err := s.campaignRepository.FindDetailCampaign(SaveTransaction.CampaignsID)
     if err != nil {

        return  err
     }

     if SaveTransaction.Status == "paid"{
         campaign.Backer_count = campaign.Backer_count +1 
         campaign.Current_amount = campaign.Backer_count + SaveTransaction.Amount

         _, err := s.campaignRepository.UpdateCampaign(campaign)

         if err != nil{

            return err
         }
     }
     return nil
     
}


