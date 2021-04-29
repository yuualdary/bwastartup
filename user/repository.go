package user

import (
	"bwastartup/models"
	"fmt"

	"gorm.io/gorm"
)



type Repository interface {
	Save(user models.Users) (models.Users, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository{
	return &repository{db}
}

func (r *repository) Save(user models.Users) (models.Users, error){
		err := r.db.Create(&user).Error
		fmt.Println(user)
		if err != nil{
			return user,err
		}
		return user, nil
}