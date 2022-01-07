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

type Controller struct {
	AuthControllerInterface
	CustomerTempAuthInterface
}

func NewController(services *service.Service) *Controller {
	return &Controller{
		AuthControllerInterface:   NewAuthController(services.AuthServiceInterface),
		CustomerTempAuthInterface: NewCustomerTempAuthController(services.CustomerTempAuthInterface),
	}
}
