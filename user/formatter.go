package user

import "bwastartup/models"

type UserFormatter struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Occupation string `json:"occupation"`
	Email      string `json:"email"`
	Token      string `json:"token"`
}

func FormatUser(user models.Users, token string) UserFormatter {
	formatter := UserFormatter{
		ID: int(user.ID),
		Name : user.Name,
		Occupation: user.Occupation,
		Email : user.Email,
		Token:  token,

	}
	return formatter
}