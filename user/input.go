package user

type RegisterUserInput struct {
	Name          string `json:"name" binding:"required"`
	Occupation    string `json:"occupation" binding:"required"`
	Email         string `json:"email" binding:"required,email"`
	Password_hash string `json:"password" binding:"required"`
}

type LoginUserInput struct {
	Email         string `json:"email" binding:"required,email"`
	Password_hash string `json:"password" binding:"required"`
}

type CheckMailInput struct {
	Email string `json:"email" binding:"required,email"`
}
