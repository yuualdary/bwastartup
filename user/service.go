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
	SaveAvatar(ID int, filelocation string) (models.Users, error) 
	GetUserById(ID int) (models.Users, error)


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

func (s *service) SaveAvatar(ID int, filelocation string) (models.Users, error){

	//get user by id
	//update avatar file name
	//save perubahan avatar file name
	user, err := s.repository.FindById(ID)

	if err !=nil{
		return user, err
	}

	user.Avatar_file_name = filelocation

	Update, err := s.repository.UpdateUser(user)

	if err != nil{
		return Update, err
	}

	return Update, nil
}


func (s *service)GetUserById(ID int) (models.Users, error){

	user, err := s.repository.FindById(ID)

	if err != nil{
		return user, err
	}

	if user.ID == 0{
		return user, errors.New("No user found")
	}
	return user, nil
}






