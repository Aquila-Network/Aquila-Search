package controller

import (
	"aquiladb/src/model"
	"aquiladb/src/service"
	"fmt"
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

func (c *CustomerAuthController) GetCustomer(ctx *gin.Context) {
	customerId, exist := ctx.Get("customer_uuid")
	if !exist {
		NewErrorResponse(ctx, http.StatusUnauthorized, "Can not find customer id.")
	}

	customer, err := c.service.GetCustomer(fmt.Sprintf("%v", customerId))
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"customer": customer,
	})
}
