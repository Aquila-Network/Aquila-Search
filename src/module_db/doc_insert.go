package moduledb

import (
	"aquiladb/src/config"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func SendHTMLForParsingToMercury(jsonDataBytes []byte) []uint8 {
	var configEnv = config.GlobalConfig

	createURL := fmt.Sprintf("http://%v:%v/process",
		configEnv.AquilaDB.Host,
		configEnv.AquilaDB.MercuryPort,
	)

	// fmt.Println("Mercury =====================================")
	// fmt.Println(createURL)

	resp, err := http.Post(
		createURL,
		// "https://httpbin.org/post",
		"application/json",
		bytes.NewBuffer(jsonDataBytes),
	)
	if err != nil {
		print(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		print(err)
	}

	fmt.Println(string(body)) // will write response in the console

	return body
}

func SendContentToTxPick(mercury *MercuryResponseStruct) []uint8 {
	var configEnv = config.GlobalConfig

	createURL := fmt.Sprintf("http://%v:%v/process",
		configEnv.AquilaDB.Host,
		configEnv.AquilaDB.TxPickPort,
	)

	txPickRequest := &TxPickRequestStruct{
		Url:  mercury.Data.Url,
		Html: mercury.Data.Content,
	}

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(txPickRequest)
	if err != nil {
		log.Fatal(err)
	}

	req, _ := json.Marshal(txPickRequest)

	resp, err := http.Post(
		createURL,
		// "https://httpbin.org/post",
		"application/json",
		bytes.NewBuffer(req),
	)
	if err != nil {
		print(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		print(err)
	}

	fmt.Println(string(body)) // will write response in the console

	return body
}

func SendTextToAquilaHub(t *TxPickResponseStruct) []uint8 {
	var configEnv = config.GlobalConfig

	createURL := fmt.Sprintf("http://%v:%v/compress",
		configEnv.AquilaDB.Host,
		configEnv.AquilaDB.AquilaHubPort,
	)

	fmt.Println("AquilaHub =====================================")
	fmt.Println(createURL)

	aquilaHubRequest := &AquilaHubRequestStruct{
		Data: AquilaDataRequestStruct{
			Text:         t.Result,
			DatabaseName: "BN4Bik3RbaY5mzJS94u8SvjZd1keyjTWaDNF36TjYzj7",
		},
	}

	fmt.Println("Aquila hub request =====================================")
	fmt.Println(aquilaHubRequest)

	fmt.Println(aquilaHubRequest)

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(aquilaHubRequest)
	if err != nil {
		log.Fatal(err)
	}

	req, _ := json.Marshal(aquilaHubRequest)

	resp, err := http.Post(
		createURL,
		// "https://httpbin.org/post",
		"application/json",
		bytes.NewBuffer(req),
	)
	if err != nil {
		print(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		print(err)
	}

	fmt.Println("===============================")
	fmt.Println(string(body)) // will write response in the console

	return body
}

func SendVectors(vectors *AquilaHubResponseStruct) []uint8 {
	var configEnv = config.GlobalConfig

	createURL := fmt.Sprintf("http://%v:%v/db/doc/insert",
		configEnv.AquilaDB.Host,
		configEnv.AquilaDB.AquilaDbPort,
	)

	docInsert := &DocInsertStruct{
		Data: DatatDocInsertStruct{
			Docs: []DocsStruct{
				{
					Payload: PayloadStruct{
						Metadata: MetadataStructDocInsert{
							Name: "name1",
							Age:  20,
						},
						Code: vectors.Vectors,
					},
				},
			},
			DatabaseName: "BN4Bik3RbaY5mzJS94u8SvjZd1keyjTWaDNF36TjYzj7",
		},
		Signature: "secret",
	}

	fmt.Println(docInsert)

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(docInsert)
	if err != nil {
		log.Fatal(err)
	}

	req, _ := json.Marshal(docInsert)

	resp, err := http.Post(
		createURL,
		// "https://httpbin.org/post",
		"application/json",
		bytes.NewBuffer(req),
	)
	if err != nil {
		print(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		print(err)
	}

	fmt.Println(string(body)) // will write response in the console

	return body
}

func DocInsert(jsonDataBytes []byte) *DocInsertResponseStruct {

	var mercuryResponseStruct *MercuryResponseStruct
	responseMercury := SendHTMLForParsingToMercury(jsonDataBytes)
	json.Unmarshal(responseMercury, &mercuryResponseStruct)

	var txPick *TxPickResponseStruct
	responseTxPick := SendContentToTxPick(mercuryResponseStruct)
	json.Unmarshal(responseTxPick, &txPick)

	var aquilaHubResponseStruct *AquilaHubResponseStruct
	responseAquilaHub := SendTextToAquilaHub(txPick)
	json.Unmarshal(responseAquilaHub, &aquilaHubResponseStruct)

	var docInsertResponse *DocInsertResponseStruct
	aquilaDbResponse := SendVectors(aquilaHubResponseStruct)
	json.Unmarshal(aquilaDbResponse, &docInsertResponse)

	return docInsertResponse

	// var configEnv = config.GlobalConfig

	// createURL := fmt.Sprintf("http://%v:%v/db/doc/insert",
	// 	configEnv.AquilaDB.Host,
	// 	configEnv.AquilaDB.Port,
	// )

	// resp, err := http.Post(
	// 	createURL,
	// 	// "https://httpbin.org/post",
	// 	"application/json",
	// 	bytes.NewBuffer(jsonDataBytes),
	// )
	// if err != nil {
	// 	print(err)
	// }
	// defer resp.Body.Close()
	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	print(err)
	// }

	// fmt.Println(string(body)) // will write response in the console

	// return body
}
