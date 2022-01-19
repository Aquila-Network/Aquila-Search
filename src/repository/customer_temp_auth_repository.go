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

func (c CustomerTempAuthPostgres) RegisterTempCustomer(customer model.CustomerTemp) (*model.CustomerTemp, error) {

	row := c.db.QueryRow(
		"INSERT INTO customers_temp (first_name, last_name, secret_key, aquila_db_database_name) values ($1, $2, $3, $4) RETURNING id, customer_id",
		customer.FirstName,
		customer.LastName,
		customer.SecretKey,
		customer.AquilaDb,
	)
	if err := row.Scan(&customer.Id, &customer.CustomerId); err != nil {
		fmt.Println("Error create customer.")
		return nil, err
	}

	successMessage := fmt.Sprintf("Customer created wiht id = %d, customer = %v", customer.Id, customer.CustomerId)
	fmt.Println(successMessage)

	return &customer, nil
}
