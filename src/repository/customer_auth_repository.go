package repository

import (
	"aquiladb/src/model"

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

func (c *CustomerAuthPostgres) CreatePermanentCustomer(customer model.Customer) (string, error) {
	var id int
	row := c.db.QueryRow(
		"INSERT INTO customers (customer_id, first_name, last_name, email, description, secret_key, aquila_db_database_name, shared_hash, is_sharable, document_number) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id",
		customer.CustomerId,
		customer.FirstName,
		customer.LastName,
		customer.Email,
		customer.Description,
		customer.SecretKey,
		customer.AquilaDb,
		customer.SharableHash,
		customer.IsSharable,
		customer.DocumentNumber,
	)
	if err := row.Scan(&id); err != nil {
		return "", err
	}

	return "Hello from repository", nil
}

func (c *CustomerAuthPostgres) FindTempCustomerBySecretKey(secretKey string) (model.CustomerTemp, error) {

	var custoemerTemp model.CustomerTemp

	err := c.db.Get(
		&custoemerTemp,
		"SELECT id, customer_id, aquila_db_database_name, shared_hash, is_sharable, document_number, is_active FROM customers_temp WHERE secret_key = $1",
		secretKey,
	)

	return custoemerTemp, err
}

func (c *CustomerAuthPostgres) GetCustomerByUUID(customerUUID string) (model.Customer, error) {

	var customer model.Customer

	err := c.db.Get(
		&customer,
		"SELECT customer_id, first_name, last_name, email, description, secret_key, aquila_db_database_name, shared_hash, is_sharable, document_number, created_at FROM customers WHERE customer_id=$1",
		customerUUID,
	)
	if err != nil {
		// for debugging
		return customer, err
		// return customer, errors.New("Customer not found.")
	}

	return customer, nil
}
