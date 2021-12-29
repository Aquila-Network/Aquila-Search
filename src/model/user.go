package model

type User struct {
	Id        int32  `json:"id" db:"id"`
	FirstName string `json:"first_name" db:"first_name"`
	LastName  string `json:"last_name" db:"last_name"`
	Email     string `json:"email" validate:"email"`
	Password  string `json:"password"`
	IsAdmin   bool   `json:"is_admin" db:"is_admin"`
	IsActive  bool   `json:"is_active" db:"is_active"`
}

type LoginUser struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" binding:"required"`
}
