package service

import "aquiladb/src/repository"

type CustomerTempAuth struct {
	repo repository.CustomerTempAuthRepositoryInterface
}

func NewCustomerTempAuthService(repo repository.CustomerTempAuthRepositoryInterface) *CustomerTempAuth {
	return &CustomerTempAuth{
		repo: repo,
	}
}

func (c CustomerTempAuth) CreateTempCustomer() (string, error) {
	return c.repo.RegisterTempCustomer()
}
