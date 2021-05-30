package payment

import (
	"bwastartup/models"
	"log"
	"strconv"

	midtrans "github.com/veritrans/go-midtrans" //pakai alias karena gaboleh ngiimport ada tanda -
)


type Service interface {
	GetPaymentURL(transaction Transaction, user models.Users)(string, error)
}

type service struct {

}

func NewService()*service{

	return &service{ }
}

func (s *service) GetPaymentURL(transaction Transaction, user models.Users)(string, error){

	midclient := midtrans.NewClient()
    midclient.ServerKey = ""
    midclient.ClientKey = ""
    midclient.APIEnvType = midtrans.Sandbox

    var snapGateway midtrans.SnapGateway
    snapGateway = midtrans.SnapGateway{
        Client: midclient,
    }
    snapReq := &midtrans.SnapReq{
        TransactionDetails: midtrans.TransactionDetails{
            OrderID: strconv.Itoa(int(transaction.ID)),
            GrossAmt: int64(transaction.Amount),
        },
        CustomerDetail: &midtrans.CustDetail{
            FName: user.Name,
            Email: user.Email,
        },
       
        
    }

    log.Println("GetToken:")
    snapTokenResp, err := snapGateway.GetToken(snapReq)
	if err != nil{
		return "", err
	}

	return snapTokenResp.RedirectURL,nil
}


