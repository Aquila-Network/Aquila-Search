package middleware

import (
	"aquiladb/src/service"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userClaims"
)

func UserIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "empty auth header",
		})
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid auth header.",
		})
		return
	}

	if len(headerParts[1]) == 0 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Token is empty.",
		})
		return
	}

	tokenClaims, err := service.ParseTokenPermanentCustomer(headerParts[1])
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Set("customer_uuid", tokenClaims.CustomerUuid)
	c.Next()
}

// ???
func AdminIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	headerParts := strings.Split(header, " ")
	tokenClaims, err := service.ParseToken(headerParts[1])
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if !tokenClaims.IsAdmin {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"error": "You are not an admin. Get out.",
		})
	}

	c.Next()
}
