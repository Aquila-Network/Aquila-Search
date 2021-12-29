package repository

import (
	"aquiladb/src/model"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{
		db: db,
	}
}

func (a *AuthPostgres) Register(user model.User) (model.User, error) {

	a.db.Get(&user, "SELECT id FROM users WHERE email=$1", user.Email)

	if user.Id != 0 {
		return user, errors.New("User already exists")
	}

	var id int32

	row := a.db.QueryRow("INSERT INTO users (first_name, last_name, email, password) values ($1, $2, $3, $4) RETURNING id", user.FirstName, user.LastName, user.Email, user.Password)
	if err := row.Scan(&id); err != nil {
		return user, err
	}

	user.Id = id
	return user, nil
}

func (a *AuthPostgres) Login(email, password string) (model.User, error) {

	var user model.User

	query := fmt.Sprintf("SELECT id, first_name, last_name, email, is_admin, is_active FROM users WHERE email=$1 AND password=$2")
	err := a.db.Get(&user, query, email, password)
	if err != nil {
		return user, errors.New("User not found.")
	}

	return user, nil
}
