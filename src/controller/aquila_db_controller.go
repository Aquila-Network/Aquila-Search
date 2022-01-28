package controller

import (
	"aquiladb/src/config"
	moduledb "aquiladb/src/module_db"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AquilaDBController struct {
}

func NewAquilaDBController() *AquilaDBController {
	return &AquilaDBController{}
}

// ===================================

// /aquila/doc_insert
func (a *AquilaDBController) DocInsert(ctx *gin.Context) {

	// SendHTMLForParsingToMercury
	var configEnv = config.GlobalConfig // find out about it and remove !!!

	// mercury ===============================================
	mercuryURL := fmt.Sprintf("http://%v:%v/process",
		configEnv.AquilaDB.Host,
		configEnv.AquilaDB.MercuryPort,
	)

	mercuryRequest := &moduledb.MercuryRequestStruct{
		Url:  "http://test.com",
		Html: "<!DOCTYPE html><html><head><title>Bla</title></head><body><h1>Test Aqula DB</h1><p>At the time, no single team member knew Go, but within a month, everyone was writing in Go and we were building out the endpoints. It was the flexibility, how easy it was to use, and the really cool concept behind Go (how Go handles native concurrency, garbage collection, and of course safety+speed.) that helped engage us during the build. Also, who can beat that cute mascot!</p></body></html>",
	}

	mercuryResponse, _ := moduledb.SendHTMLForParsingToMercury(mercuryRequest, mercuryURL)

	// TxPick ===============================================
	txPicRequest := &moduledb.TxPickRequestStruct{
		Url:  mercuryResponse.Data.Url,
		Html: mercuryResponse.Data.Content,
	}

	txPickURL := fmt.Sprintf("http://%v:%v/process",
		configEnv.AquilaDB.Host,
		configEnv.AquilaDB.TxPickPort,
	)

	txPickResponse, _ := moduledb.SendContentToTxPick(txPicRequest, txPickURL)

	// Aquila Hub ===============================================
	aquilaHubUrl := fmt.Sprintf("http://%v:%v/compress",
		configEnv.AquilaDB.Host,
		configEnv.AquilaDB.AquilaHubPort,
	)

	aquilaHubRequest := &moduledb.AquilaHubRequestStruct{
		Data: moduledb.AquilaDataRequestStruct{
			Text:         txPickResponse.Result,
			DatabaseName: "BN4Bik3RbaY5mzJS94u8SvjZd1keyjTWaDNF36TjYzj7",
		},
	}

	aquilaHubResponse, _ := moduledb.SendTextToAquilaHub(aquilaHubRequest, aquilaHubUrl)

	// Aquila Hub ===============================================
	docInsertURL := fmt.Sprintf("http://%v:%v/db/doc/insert",
		configEnv.AquilaDB.Host,
		configEnv.AquilaDB.AquilaDbPort,
	)

	fmt.Println(docInsertURL)

	docInsert := &moduledb.DocInsertRequestStruct{
		Data: moduledb.DatatDocInsertStruct{
			Docs: []moduledb.DocsStruct{
				{
					Payload: moduledb.PayloadStruct{
						Metadata: moduledb.MetadataStructDocInsert{
							Name: "name1",
							Age:  20,
						},
						Code: aquilaHubResponse.Vectors[0], // ????
					},
				},
				{
					Payload: moduledb.PayloadStruct{
						Metadata: moduledb.MetadataStructDocInsert{
							Name: "name1",
							Age:  20,
						},
						Code: []float64{0.1, 0.2, 0.3},
					},
				},
			},
			DatabaseName: "BN4Bik3RbaY5mzJS94u8SvjZd1keyjTWaDNF36TjYzj7",
		},
		Signature: "secret",
	}

	docInsertResponse, _ := moduledb.SendVectors(docInsert, docInsertURL)

	ctx.JSON(http.StatusOK, gin.H{
		"response": docInsertResponse,
	})
}

// Doc Delete
func (a *AquilaDBController) DocDelete(ctx *gin.Context) {

	var configEnv = config.GlobalConfig

	url := fmt.Sprintf("http://%v:%v/db/doc/delete",
		configEnv.AquilaDB.Host,
		configEnv.AquilaDB.AquilaDbPort,
	)

	docDelete := &moduledb.DocDeleteRequestStruct{
		Data: moduledb.DeleteDataStruct{
			Ids: []string{
				"3gwTnetiYJfHTBcqGwoxETLsmmdGYVsd5MRBohuTG22C",
				"BXsbHy9B3tU9zaHwU41jATzDBisNEFa67XKvYZhB2fzQ",
			},
			DatabaseName: "BN4Bik3RbaY5mzJS94u8SvjZd1keyjTWaDNF36TjYzj7",
		},
		Signature: "secret",
	}

	responseDelete, err := moduledb.DocDelete(docDelete, url)
	if err != nil {
		NewErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":  "Delete Document",
		"response": &responseDelete,
	})
}

// Dock search
func (a *AquilaDBController) DocSearch(ctx *gin.Context) {

	var configEnv = config.GlobalConfig

	url := fmt.Sprintf("http://%v:%v/db/search",
		configEnv.AquilaDB.Host,
		configEnv.AquilaDB.AquilaDbPort,
	)

	// pattern
	// https://www.geeksforgeeks.org/slice-of-slices-in-golang/
	matrix := make([][]float64, 1)
	matrix[0] = make([]float64, 1)
	matrix[0] = []float64{
		-0.01806008443236351, -0.17380790412425995, 0.03992759436368942, 0.43514639139175415,
	}
	searchBody := &moduledb.SearchAquilaDbRequestStruct{
		Data: moduledb.DataSearchStruct{
			Matrix:       matrix,
			K:            10,
			R:            0,
			DatabaseName: "BN4Bik3RbaY5mzJS94u8SvjZd1keyjTWaDNF36TjYzj7",
		},
	}

	response, err := moduledb.Search(searchBody, url)
	if err != nil {
		NewErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"response": response,
	})
}
