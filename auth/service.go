package auth

import "github.com/dgrijalva/jwt-go"

type Service interface {
	GenerateToken(userID int) (string, error)
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