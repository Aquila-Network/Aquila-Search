package repository

import (
	"aquiladb/src/model"
	"fmt"

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

func (c CustomerTempAuthPostgres) RegisterTempCustomer(customer model.CustomerTemp) (string, error) {

	var id int

	row := c.db.QueryRow(
		"INSERT INTO customers_temp (first_name, last_name, secret_key) values ($1, $2, $3) RETURNING id",
		customer.FirstName,
		customer.LastName,
		customer.SecretKey,
	)
	if err := row.Scan(&id); err != nil {
		fmt.Println("Error create customer.")
		return err.Error(), err
	}

	successMessage := fmt.Sprintf("Customer created wiht id = %d", id)

	return successMessage, nil
}
