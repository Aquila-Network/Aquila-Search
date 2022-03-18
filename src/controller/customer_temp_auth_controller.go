package controller

import (
	"aquiladb/src/model"
	"aquiladb/src/service"
	"fmt"
	"math/rand"
	"net/http"
	"strings"

	moduleDb "github.com/Aquila-Network/go-aquila"
	moduleDbSrc "github.com/Aquila-Network/go-aquila/src"
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

	createAquilaDb := &moduleDbSrc.DataStructCreateDb{
		Schema: moduleDbSrc.SchemaStruct{
			Description: fmt.Sprintf("Database of %v %v", customer.FirstName, customer.LastName),
			Unique:      customer.SecretKey,
			Encoder:     "strn:msmarco-distilbert-base-tas-b",
			Codelen:     768,
			Metadata: moduleDbSrc.MetadataStructCreateDb{
				Name: "string",
				Age:  "number",
			},
		},
	}

	// create url for aquila db
	createURL := fmt.Sprintf("http://%v:%v/db/create",
		configEnv.AquilaDB.Host,
		configEnv.AquilaDB.AquilaDbPort,
	)

	walletInitStruct, err := CreateWalletSign(createAquilaDb)
	if err != nil {
		NewErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	// create aquila database
	responseAquilaDb, errResponseAquila := moduleDb.AquilaModule(walletInitStruct).AquilaDbInterface.CreateDatabase(createAquilaDb, createURL)
	if errResponseAquila != nil {
		NewErrorResponse(ctx, http.StatusBadGateway, errResponseAquila.Error())
		return
	}

	customer.AquilaDb = responseAquilaDb.DatabaseName

	// get newest temporary created customer
	customer, errCreate := c.service.CreateTempCustomer(customer)
	if errCreate != nil {
		NewErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"secret_key":    customer.SecretKey,
		"token":         customer.Token,
		"database_name": responseAquilaDb.DatabaseName,
	})
}
