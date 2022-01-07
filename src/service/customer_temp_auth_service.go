package service

import (
	"aquiladb/src/model"
	"aquiladb/src/repository"
	"math/rand"
	"time"
)

type CustomerTempAuth struct {
	repo repository.CustomerTempAuthRepositoryInterface
}

func NewCustomerTempAuthService(repo repository.CustomerTempAuthRepositoryInterface) *CustomerTempAuth {
	return &CustomerTempAuth{
		repo: repo,
	}
}

func (c CustomerTempAuth) CreateTempCustomer() (string, error) {

	var customer model.CustomerTemp

	customer.FirstName = "Bob"
	customer.LastName = "Stone"
	customer.SecretKey = KeyGenerate(15)

	return c.repo.RegisterTempCustomer(customer)
}

func KeyGenerate(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	var seededRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)

	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
