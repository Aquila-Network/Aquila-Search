package controller

import (
	"aquiladb/src/service"

	"github.com/gin-gonic/gin"
)

type AuthControllerInterface interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
}

type Controller struct {
	AuthControllerInterface
}

func NewController(services *service.Service) *Controller {
	return &Controller{
		AuthControllerInterface: NewAuthController(services.AuthServiceInterface),
	}
}
