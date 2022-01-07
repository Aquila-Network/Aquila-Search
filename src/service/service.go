package service

import (
	"aquiladb/src/model"
	"aquiladb/src/repository"
)

type AuthServiceInterface interface {
	Register(model.User) (string, error)
	Login(model.LoginUser) (string, error)
}

type CustomerTempAuthInterface interface {
	CreateTempCustomer() (model.CustomerTemp, error)
}

type Service struct {
	AuthServiceInterface
	CustomerTempAuthInterface
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		AuthServiceInterface:      NewAuthService(repos.AuthRepositoryInterface),
		CustomerTempAuthInterface: NewCustomerTempAuthService(repos.CustomerTempAuthRepositoryInterface),
	}
}
