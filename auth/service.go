package auth

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
)


type Service interface {
	GenerateToken(userID int) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtService struct {
}

var SECRET_KEY = []byte("BWASTARTUP_s3cr3T_k3Y")//ini bahaya, harus diketahui dan hanya diketahui oleh developer

func NewService() *jwtService{

	return &jwtService{}
}

func (s *jwtService) GenerateToken(userID int) (string, error){
	
	claim := jwt.MapClaims{}
	claim["user_id"]=userID //value dari user 
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	//token valid jika, dibuat dengan secret key
	SignedToken, err := token.SignedString(SECRET_KEY)

	if err != nil {
		return SignedToken, err
	}

	return SignedToken, nil
}

func (s *jwtService) ValidateToken(EncodeToken string) (*jwt.Token, error){
	token, err := jwt.Parse(EncodeToken, func(token *jwt.Token) (interface{}, error) {//func ny bawaan
		//jadi fungsi func mengecek apakah token yang dibuat sesuai dengan secret_key yand kita buat
		_, ok := token.Method.(*jwt.SigningMethodHMAC)//tipenya HMAC karena diatas pake HS256

		if !ok{

			return nil, errors.New("invalid token")
		}
		return []byte(SECRET_KEY),nil
	})

	if err != nil{
		return token, err
	}

	return token, nil

}



