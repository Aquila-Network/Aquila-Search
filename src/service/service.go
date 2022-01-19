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
	CreateTempCustomer(model.CustomerTemp) (model.CustomerTemp, error)
}

type CustomerAuthServiceInterface interface {
	CreatePermanentCustomer(model.Customer) (model.Customer, error)
	GetCustomer(custoemrUUID string) (model.Customer, error)
	Auth(secretKey string) (string, error)
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
