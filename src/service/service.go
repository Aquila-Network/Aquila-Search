package service

import (
	"aquiladb/src/model"
	"aquiladb/src/repository"
)

type AuthServiceInterface interface {
	Register(model.User) (string, error)
	Login(model.LoginUser) (string, error)
}

type Service struct {
	AuthServiceInterface
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		AuthServiceInterface: NewAuthService(repos.AuthRepositoryInterface),
	}
}
