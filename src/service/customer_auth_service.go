package service

import (
	"aquiladb/src/repository"
)

type CustomerAuth struct {
	repo repository.CustomerAuthRepositoryInterface
}

func NewCustomerAuthService(repo repository.CustomerAuthRepositoryInterface) *CustomerAuth {
	return &CustomerAuth{
		repo: repo,
	}
}

func (c *CustomerAuth) CreatePermanentCustomer() (string, error) {
	return c.repo.CreatePermanentCustomer()
}
