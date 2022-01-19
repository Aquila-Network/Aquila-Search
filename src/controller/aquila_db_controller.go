package controller

import (
	moduledb "aquiladb/src/module_db"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateDBResponse struct {
	DatabaseName string `json:"database_name"`
	Success      bool   `json:"success"`
}

type DocInsert struct {
	Ids     []string `json:"ids"`
	Success bool     `json:"success"`
}

type AquilaDBController struct {
}

func NewAquilaDBController() *AquilaDBController {
	return &AquilaDBController{}
}

func (a *AquilaDBController) CreateAquilaDB(ctx *gin.Context) {
	// var createDB *CreateDBResponse
	// jsonDataBytes, err := ioutil.ReadAll(ctx.Request.Body)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println(jsonDataBytes)

	// response := moduledb.CreateAquilaDatabase(jsonDataBytes)
	// json.Unmarshal(response, &createDB)

	// ctx.JSON(http.StatusOK, gin.H{
	// 	"response": &createDB,
	// })
}

func (a *AquilaDBController) DocInsert(ctx *gin.Context) {
	var docInsert *DocInsert
	jsonDataBytes, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		fmt.Println(err)
	}

	response := moduledb.DocInsert(jsonDataBytes)
	json.Unmarshal(response, &docInsert)

	ctx.JSON(http.StatusOK, gin.H{
		"response": &docInsert,
	})
}
