package repository

import "github.com/jmoiron/sqlx"

type CustomerTempAuthPostgres struct {
	db *sqlx.DB
}

func NewCustomerTempAuthRepository(db *sqlx.DB) *CustomerTempAuthPostgres {
	return &CustomerTempAuthPostgres{
		db: db,
	}
}

func (c CustomerTempAuthPostgres) RegisterTempCustomer() (string, error) {
	return "Hello from repository.", nil
}
