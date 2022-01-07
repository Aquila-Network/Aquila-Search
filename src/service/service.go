package service

import (
	"aquiladb/src/model"
	"aquiladb/src/repository"
)

type AuthServiceInterface interface {
	Register(model.User) (string, error)
	Login(model.LoginUser) (string, error)
}

type CustomerTempAuthServiceInterface interface {
	CreateTempCustomer() (model.CustomerTemp, error)
}

type CustomerAuthServiceInterface interface {
	CreatePermanentCustomer() (string, error)
}

type Service struct {
	AuthServiceInterface
	CustomerTempAuthServiceInterface
	CustomerAuthServiceInterface
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		AuthServiceInterface:             NewAuthService(repos.AuthRepositoryInterface),
		CustomerTempAuthServiceInterface: NewCustomerTempAuthService(repos.CustomerTempAuthRepositoryInterface),
		CustomerAuthServiceInterface:     NewCustomerAuthService(repos.CustomerAuthRepositoryInterface),
	}
}
