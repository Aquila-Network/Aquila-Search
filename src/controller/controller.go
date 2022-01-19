package controller

import (
	"aquiladb/src/service"

	"github.com/gin-gonic/gin"
)

type AuthControllerInterface interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
}

type CustomerTempAuthInterface interface {
	CreateTempCustomer(c *gin.Context)
}

type CustomerAuthInterface interface {
	CreatePermanentCustomer(c *gin.Context)
	GetCustomer(c *gin.Context)
	Auth(c *gin.Context)
}

type AquilaDBInterface interface {
	CreateAquilaDB(c *gin.Context)
	DocInsert(c *gin.Context)
}

type Controller struct {
	AuthControllerInterface
	CustomerTempAuthInterface
	CustomerAuthInterface
	AquilaDBInterface
}

func NewController(services *service.Service) *Controller {
	return &Controller{
		AuthControllerInterface:   NewAuthController(services.AuthServiceInterface),
		CustomerTempAuthInterface: NewCustomerTempAuthController(services.CustomerTempAuthServiceInterface),
		CustomerAuthInterface:     NewCustomerAuthController(services.CustomerAuthServiceInterface),
		AquilaDBInterface:         NewAquilaDBController(),
	}
}
