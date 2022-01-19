package controller

import (
	"aquiladb/src/model"
	moduledb "aquiladb/src/module_db"
	"aquiladb/src/service"
	"fmt"
	"math/rand"
	"net/http"
	"strings"

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
// @Produce  json
// @Success 200 {string} secret_key 1
// @Success 200 {string} token 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /customer [post]
func (c CustomerTempAuthController) CreateTempCustomer(ctx *gin.Context) {

	var customer model.CustomerTemp
	randomAdjective := service.ADJECTIVES[rand.Intn(len(service.ADJECTIVES))]
	randomNoun := service.NOUNS[rand.Intn(len(service.NOUNS))]

	customer.FirstName = strings.Title(randomAdjective)
	customer.LastName = strings.Title(randomNoun)
	customer.SecretKey = service.KeyGenerate(14)

	createAquilaDb := &moduledb.CreateDbStruct{
		Data: moduledb.DataStructCreateDb{
			Schema: moduledb.SchemaStruct{
				Description: fmt.Sprintf("Database of %v %v", customer.FirstName, customer.LastName),
				Unique:      customer.SecretKey,
				Encoder:     "strn:msmarco-distilbert-base-tas-b",
				Codelen:     768,
				Metadata: moduledb.MetadataStructCreateDb{
					Name: "string",
					Age:  "number",
				},
			},
		},
		Signature: "secret",
	}

	// create aquila database
	responseAquilaDb := moduledb.CreateAquilaDatabase(createAquilaDb)

	customer.AquilaDb = responseAquilaDb.DatabaseName

	// get newest temporary created customer
	customer, err := c.service.CreateTempCustomer(customer)
	if err != nil {
		NewErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"secret_key":    customer.SecretKey,
		"token":         customer.Token,
		"database_name": responseAquilaDb.DatabaseName,
	})
}
