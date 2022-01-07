package repository

import (
	"errors"

	"github.com/jmoiron/sqlx"
)

type CustomerAuthPostgres struct {
	db *sqlx.DB
}

func NewCustomerAuthRepository(db *sqlx.DB) *CustomerAuthPostgres {
	return &CustomerAuthPostgres{
		db: db,
	}
}

func (c *CustomerAuthPostgres) CreatePermanentCustomer() (string, error) {
	return "Hello from repository", errors.New("Bob")
}
