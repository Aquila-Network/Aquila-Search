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
	a, _ := c.service.CreateTempCustomer()
	ctx.JSON(http.StatusOK, gin.H{
		"token": a,
	})
}
