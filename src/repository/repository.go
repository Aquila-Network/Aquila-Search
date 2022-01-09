package repository

import (
	"aquiladb/src/model"

	"github.com/jmoiron/sqlx"
)

type AuthRepositoryInterface interface {
	Register(model.User) (model.User, error)
	Login(email, password string) (model.User, error)
}

type UserRepositoryInterface interface {
	GetAllUsers()
	GetUserById()
}

type CustomerTempAuthRepositoryInterface interface {
	RegisterTempCustomer(model.CustomerTemp) (*model.CustomerTemp, error)
}

type CustomerAuthRepositoryInterface interface {
	CreatePermanentCustomer(model.Customer) (string, error)
	FindTempCustomerBySecretKey(secretKey string) (model.CustomerTemp, error)
	GetCustomerByUUID(customerUUID string) (model.Customer, error)
}

type Repository struct {
	AuthRepositoryInterface
	UserRepositoryInterface
	CustomerTempAuthRepositoryInterface
	CustomerAuthRepositoryInterface
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		AuthRepositoryInterface:             NewAuthPostgres(db),
		UserRepositoryInterface:             NewUserPostgres(db),
		CustomerTempAuthRepositoryInterface: NewCustomerTempAuthRepository(db),
		CustomerAuthRepositoryInterface:     NewCustomerAuthRepository(db),
	}
}
