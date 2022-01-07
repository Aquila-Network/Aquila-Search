package controller

import (
	"aquiladb/src/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CustomerTempAuthController struct {
	service service.CustomerTempAuthInterface
}

func NewCustomerTempAuthController(service service.CustomerTempAuthInterface) *CustomerTempAuthController {
	return &CustomerTempAuthController{
		service: service,
	}
}

func (c CustomerTempAuthController) CreateTempCustomer(ctx *gin.Context) {
	customer, err := c.service.CreateTempCustomer()
	if err != nil {
		NewErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"secret_key": customer.SecretKey,
		"token":      customer.Token,
	})
}
