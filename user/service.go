package user

import (
	"bwastartup/models"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)


type Service interface {
	RegisterUser(input RegisterUserInput) (models.Users, error)
	LoginUser(input LoginUserInput) (models.Users, error)
	CheckMailUser(input CheckMailInput) (bool, error)


}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(input RegisterUserInput) (models.Users, error) {
	user := models.Users{}
	user.Name = input.Name	
	user.Email = input.Email
	user.Occupation = input.Occupation
	
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password_hash),bcrypt.MinCost)

	if err != nil {
		return user, err
	}
	user.Password_hash = string(passwordHash)
	user.Role = "user"

	NewUser, err := s.repository.Save(user)
	if err != nil{
		return NewUser, err
	}
	return NewUser, nil

}

func (s *service) LoginUser(input LoginUserInput) (models.Users, error){

	email := input.Email
	password := input.Password_hash

	GetUser, err := s.repository.FindByEmail(email)

	if err != nil {
		
		return GetUser, nil
	}

	if GetUser.ID == 0 {
		return GetUser, errors.New("no user found on that email") //buat custom 
	}

	fmt.Println(password)

	err = bcrypt.CompareHashAndPassword([]byte(GetUser.Password_hash), []byte(password))


	if err != nil{
		return GetUser, err//jangan lupa, kalau mau ngehasilin otuput error harus return errornya (err)
	}
	return GetUser, nil
}

func (s *service) CheckMailUser(input CheckMailInput) (bool, error){
	email := input.Email

	user,err:=s.repository.FindByEmail(email)

	if err != nil{
		return false, err
	}
	//check ada error atau tidak
	
	//check jika email sudah digunakan atau tidak
	if user.ID == 0 {
		return true, nil
	}

	return false, nil

}




