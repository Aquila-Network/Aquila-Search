package repository

import "github.com/jmoiron/sqlx"

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{
		db: db,
	}
}

func (a *UserPostgres) GetAllUsers() {

}

func (a *UserPostgres) GetUserById() {

}
