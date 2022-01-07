package controller

import (
	"aquiladb/src/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CustomerAuthController struct {
	service service.CustomerAuthServiceInterface
}

func NewCustomerAuthController(service service.CustomerAuthServiceInterface) *CustomerAuthController {
	return &CustomerAuthController{
		service: service,
	}
}

func (c *CustomerAuthController) CreatePermanentCustomer(ctx *gin.Context) {
	a, _ := c.service.CreatePermanentCustomer()

	ctx.JSON(http.StatusOK, gin.H{
		"secret_key": a,
	})
}
