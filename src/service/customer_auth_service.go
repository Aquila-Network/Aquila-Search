package service

import (
	"aquiladb/src/model"
	"aquiladb/src/repository"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type CustomerAuth struct {
	repo repository.CustomerAuthRepositoryInterface
}

type TokenClaimsCustomer struct {
	jwt.StandardClaims
	CustomerUuid string `json:"customer_uuid"`
	IsPermanent  bool   `json:"is_permanent"`
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

	customerTemp, err := c.repo.FindTempCustomerBySecretKey(customer.SecretKey)
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

func (c *CustomerAuth) GetCustomer(customerUUID string) (model.Customer, error) {
	return c.repo.GetCustomerByUUID(customerUUID)
}

func (c *CustomerAuth) Auth(secretKey string) (string, error) {

	if len(secretKey) < 1 {
		return "", errors.New("Secret key is required 'secret_key'.")
	}

	customer, err := c.repo.GetCustomerBySecretKey(secretKey)
	if err != nil {
		return "", err
	}

	customer.IsPermanent = true
	token, err := GenerateTokenCustomer(customer)
	if err != nil {
		return "", err
	}

	return token, nil
}

func GenerateTokenCustomer(customer model.Customer) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaimsCustomer{
		jwt.StandardClaims{
			// it should be lifetime
			// ExpiresAt: time.Now().Add(tokenTTlForTemporaryUser).Unix(),
			IssuedAt: time.Now().Unix(),
		},
		customer.CustomerId,
		customer.IsPermanent,
	})

	return token.SignedString([]byte(signingKey))
}
