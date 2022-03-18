package controller

import (
	"aquiladb/src/config"
	localmoduledb "aquiladb/src/module_db"
	"fmt"
	"io/ioutil"
	"net/http"

	moduleDb "github.com/Aquila-Network/go-aquila"
	moduleDbSrc "github.com/Aquila-Network/go-aquila/src"
	"github.com/gin-gonic/gin"
)

var configEnv = config.GlobalConfig

type AquilaDBController struct {
}

func NewAquilaDBController() *AquilaDBController {
	return &AquilaDBController{}
}

// ===================================

// /aquila/doc_insert
func (a *AquilaDBController) DocInsert(ctx *gin.Context) {

	// SendHTMLForParsingToMercury
	// mercury ===============================================
	// localhost:5009/process
	mercuryURL := fmt.Sprintf("http://%v:%v/process",
		configEnv.AquilaDB.Host,
		configEnv.AquilaDB.MercuryPort,
	)

	mercuryRequest := &localmoduledb.MercuryRequestStruct{
		Url:  "http://test.com",
		Html: "<!DOCTYPE html><html><head><title>Bla</title></head><body><h1>Test Aqula DB</h1><p>At the time, no single team member knew Go, but within a month, everyone was writing in Go and we were building out the endpoints. It was the flexibility, how easy it was to use, and the really cool concept behind Go (how Go handles native concurrency, garbage collection, and of course safety+speed.) that helped engage us during the build. Also, who can beat that cute mascot!</p></body></html>",
	}

	mercuryResponse, _ := localmoduledb.SendHTMLForParsingToMercury(mercuryRequest, mercuryURL)

	// TxPick ===============================================
	// localhost:5008/process
	txPicRequest := &localmoduledb.TxPickRequestStruct{
		Url:  mercuryResponse.Data.Url,
		Html: mercuryResponse.Data.Content,
	}

	txPickURL := fmt.Sprintf("http://%v:%v/process",
		configEnv.AquilaDB.Host,
		configEnv.AquilaDB.TxPickPort,
	)

	txPickResponse, _ := localmoduledb.SendContentToTxPick(txPicRequest, txPickURL)

	// Aquila Hub ===============================================
	// localhost:5002/compress
	aquilaHubUrl := fmt.Sprintf("http://%v:%v/compress",
		configEnv.AquilaDB.Host,
		configEnv.AquilaDB.AquilaHubPort,
	)

	aquilaHubRequest := &localmoduledb.AquilaHubRequestStruct{
		Data: localmoduledb.AquilaDataRequestStruct{
			Text:         txPickResponse.Result,
			DatabaseName: "BN4Bik3RbaY5mzJS94u8SvjZd1keyjTWaDNF36TjYzj7",
		},
	}

	aquilaHubResponse, _ := localmoduledb.SendTextToAquilaHub(aquilaHubRequest, aquilaHubUrl)

	// ===============================================
	// module
	// ===============================================
	// localhost:5001/db/doc/insert
	docInsertURL := fmt.Sprintf("http://%v:%v/db/doc/insert",
		configEnv.AquilaDB.Host,
		configEnv.AquilaDB.AquilaDbPort,
	)

	docInsert := &moduleDbSrc.DatatDocInsertStruct{
		Docs: []moduleDbSrc.DocsStruct{
			{
				Payload: moduleDbSrc.PayloadStruct{
					Metadata: moduleDbSrc.MetadataStructDocInsert{
						Name: "name1",
						Age:  20,
					},
					Code: aquilaHubResponse.Vectors[0], // ????
				},
			},
			{
				Payload: moduleDbSrc.PayloadStruct{
					Metadata: moduleDbSrc.MetadataStructDocInsert{
						Name: "name1",
						Age:  20,
					},
					Code: []float64{0.1, 0.2, 0.3},
				},
			},
		},
		DatabaseName: "BN4Bik3RbaY5mzJS94u8SvjZd1keyjTWaDNF36TjYzj7",
	}

	walletInitStruct, err := CreateWalletSign(docInsert)
	if err != nil {
		NewErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	// create database
	docInsertResponse, err := moduleDb.AquilaModule(walletInitStruct).AquilaDbInterface.InsertDocument(docInsert, docInsertURL)
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadGateway, err.Error())
	}

	ctx.JSON(http.StatusOK, gin.H{
		"response": docInsertResponse,
	})
}

// Doc Delete
func (a *AquilaDBController) DocDelete(ctx *gin.Context) {

	url := fmt.Sprintf("http://%v:%v/db/doc/delete",
		configEnv.AquilaDB.Host,
		configEnv.AquilaDB.AquilaDbPort,
	)

	docDelete := &moduleDbSrc.DeleteDataStruct{
		Ids: []string{
			"3gwTnetiYJfHTBcqGwoxETLsmmdGYVsd5MRBohuTG22C",
			"BXsbHy9B3tU9zaHwU41jATzDBisNEFa67XKvYZhB2fzQ",
		},
		DatabaseName: "BN4Bik3RbaY5mzJS94u8SvjZd1keyjTWaDNF36TjYzj7",
	}

	walletInitStruct, err := CreateWalletSign(docDelete)
	if err != nil {
		NewErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	responseDelete, err := moduleDb.AquilaModule(walletInitStruct).AquilaDbInterface.DeleteDocument(docDelete, url)
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
	searchBody := &moduleDbSrc.DataSearchStruct{
		Matrix:       matrix,
		K:            10,
		R:            0,
		DatabaseName: "BN4Bik3RbaY5mzJS94u8SvjZd1keyjTWaDNF36TjYzj7",
	}

	walletInitStruct, err := CreateWalletSign(searchBody)
	if err != nil {
		NewErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	response, err := moduleDb.AquilaModule(walletInitStruct).AquilaDbInterface.SearchKDocument(searchBody, url)
	if err != nil {
		NewErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"response": response,
	})
}

func CreateWalletSign(requestStructure interface{}) (moduleDbSrc.WalletStruct, error) {

	var wallet moduleDbSrc.WalletStruct
	filePath := configEnv.PrivateUnencryptedPemFileStruct.PathToPrivateUnencryptedPemFile

	priv, err := ioutil.ReadFile(filePath)
	if err != nil {
		return wallet, err
	}
	walletInitStruct := moduleDbSrc.NewWallet(string(priv[:]))
	walletSign, err := walletInitStruct.CreateSignatureWallet(requestStructure)
	if err != nil {
		return wallet, err
	}
	walletInitStruct.SecretKey = walletSign

	return walletInitStruct, nil
}
