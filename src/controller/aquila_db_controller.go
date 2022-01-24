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
	Url  string `json:"url"`
	Html string `json:"html"`
}

type DocDelete struct {
	Ids     []string `json:"ids"`
	Success bool     `json:"success"`
}

type MetadataSearchStruct struct {
	Age  int
	Name string
}

type DocSearchData struct {
	Cid      string
	Id       int
	Code     []float32
	Metadata MetadataSearchStruct
}

type DocSearch struct {
	Dist [][]float64
	Docs [][]DocSearchData
}

type AquilaDBController struct {
}

func NewAquilaDBController() *AquilaDBController {
	return &AquilaDBController{}
}

// no need of it because aquila will create automatically
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

	// var docInsert *DocInsert
	jsonDataBytes, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		fmt.Println(err)
	}

	response := moduledb.DocInsert(jsonDataBytes)
	// json.Unmarshal(response, &docInsert)

	ctx.JSON(http.StatusOK, gin.H{
		"response": response,
	})
}

func (a *AquilaDBController) DocDelete(ctx *gin.Context) {

	var docDelete *DocDelete
	// jsonDataBytes, err := ioutil.ReadAll(ctx.Request.Body)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	response := moduledb.DocDelete()
	json.Unmarshal(response, &docDelete)

	ctx.JSON(http.StatusOK, gin.H{
		"message":  "bla",
		"response": &docDelete,
	})
}

func (a *AquilaDBController) DocSearch(ctx *gin.Context) {
	// var docSearch *DocSearch

	moduledb.Search()
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Look to the console.",
		// "response": &docDelete,
	})
}
