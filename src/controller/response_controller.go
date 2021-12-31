package controller

import "github.com/gin-gonic/gin"

type errorResponse struct {
	Status  int    `json:"status_code"`
	Message string `json:"error"`
}

type successResponse struct {
	Message string `json:"message"`
}

func NewErrorResponse(c *gin.Context, statusCode int, message string) {
	c.AbortWithStatusJSON(statusCode, errorResponse{statusCode, message})
}

func NewSuccessResponse(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, successResponse{message})
}
