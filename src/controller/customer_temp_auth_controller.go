package controller

import (
	"aquiladb/src/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CustomerTempAuthController struct {
	service service.CustomerTempAuthServiceInterface
}

func NewCustomerTempAuthController(service service.CustomerTempAuthServiceInterface) *CustomerTempAuthController {
	return &CustomerTempAuthController{
		service: service,
	}
}

// @Summary Create temp customer
// @Tags create temp customer
// @Description Create temp customer
// @Router /customer [post]
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
