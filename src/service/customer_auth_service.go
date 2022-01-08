package service

import (
	"aquiladb/src/model"
	"aquiladb/src/repository"
	"errors"
	"fmt"
)

type CustomerAuth struct {
	repo repository.CustomerAuthRepositoryInterface
}

func NewCustomerAuthService(repo repository.CustomerAuthRepositoryInterface) *CustomerAuth {
	return &CustomerAuth{
		repo: repo,
	}
}

func (c *CustomerAuth) CreatePermanentCustomer(customer model.Customer) (model.Customer, error) {

	if len(customer.FirstName) < 1 {
		return customer, errors.New("First name is required 'first_name'.")
	}

	if len(customer.LastName) < 1 {
		return customer, errors.New("Last name is required 'last_name'.")
	}

	if len(customer.Email) < 1 {
		return customer, errors.New("Email is required 'email'.")
	}

	if IsEmailValid(customer.Email) == false {
		return customer, errors.New("Email is not valid.")
	}

	if len(customer.SecretKey) < 1 {
		return customer, errors.New("Secret key is required 'secret_key'.")
	}

	customerTemp, err := c.repo.FindCustomerBySecretKey(customer.SecretKey)
	if err != nil {
		return customer, errors.New("Custoemer with this secret key not found.")
	}

	customer.CustomerId = customerTemp.CustomerId
	customer.AquilaDb = customerTemp.AquilaDb
	customer.SharableHash = customerTemp.SharableHash
	customer.IsSharable = customerTemp.IsSharable
	customer.DocumentNumber = customerTemp.DocumentNumber

	response, err := c.repo.CreatePermanentCustomer(customer)
	if err != nil {
		return customer, err
	}

	fmt.Println(response)

	return customer, nil
}
