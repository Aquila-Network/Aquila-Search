package controller

import (
	"aquiladb/src/model"
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
	var customer model.Customer

	if err := ctx.ShouldBindJSON(&customer); err != nil {
		NewErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	response, err := c.service.CreatePermanentCustomer(customer)
	if err != nil {
		NewErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"customer": response,
	})
}
