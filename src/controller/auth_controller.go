package controller

import (
	"aquiladb/src/model"
	"aquiladb/src/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	services service.AuthServiceInterface
}

func NewAuthController(services service.AuthServiceInterface) *AuthController {
	return &AuthController{
		services: services,
	}
}

func (c *AuthController) Register(ctx *gin.Context) {
	var userData model.User

	if err := ctx.ShouldBindJSON(&userData); err != nil {
		NewErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	token, err := c.services.Register(userData)
	if err != nil {
		NewErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Registereg successfully.",
		"token":   token,
	})
}

func (c *AuthController) Login(ctx *gin.Context) {

	var loginUser model.LoginUser

	if err := ctx.ShouldBindJSON(&loginUser); err != nil {
		NewErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	token, err := c.services.Login(loginUser)
	if err != nil {
		NewErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Login successfully.",
		"token":   token,
	})
}
