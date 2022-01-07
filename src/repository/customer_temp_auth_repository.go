package repository

import (
	"aquiladb/src/model"

	"github.com/jmoiron/sqlx"
)

type CustomerTempAuthPostgres struct {
	db *sqlx.DB
}

func NewCustomerTempAuthRepository(db *sqlx.DB) *CustomerTempAuthPostgres {
	return &CustomerTempAuthPostgres{
		db: db,
	}
}

func (c CustomerTempAuthPostgres) RegisterTempCustomer(custoemr model.CustomerTemp) (string, error) {
	return "Hello from repository.", nil
}
