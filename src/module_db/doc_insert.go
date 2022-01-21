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

	// fmt.Println("=====================================")
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

func SendTextToAquilaHub(t *TxPickResponseStruct) {
	var configEnv = config.GlobalConfig

	createURL := fmt.Sprintf("http://%v:%v/compress",
		configEnv.AquilaDB.Host,
		configEnv.AquilaDB.AquilaHubPort,
	)

	aquilaHubRequest := &AquilaHubRequestStruct{
		Data: AquilaDataRequestStruct{
			Text:         t.Result,
			DatabaseName: "BN4Bik3RbaY5mzJS94u8SvjZd1keyjTWaDNF36TjYzj7",
		},
	}

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

	fmt.Println(string(body)) // will write response in the console

	// return body
}

func DocInsert(jsonDataBytes []byte) *MercuryResponseStruct {

	var mercuryResponseStruct *MercuryResponseStruct
	responseMercury := SendHTMLForParsingToMercury(jsonDataBytes)
	json.Unmarshal(responseMercury, &mercuryResponseStruct)

	var txPick *TxPickResponseStruct
	responseTxPick := SendContentToTxPick(mercuryResponseStruct)
	json.Unmarshal(responseTxPick, &txPick)

	SendTextToAquilaHub(txPick)

	return mercuryResponseStruct

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
