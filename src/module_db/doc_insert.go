package moduledb

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Send row html to mercury to parse
// in response will be json.
// We need take url and content field
func SendHTMLForParsingToMercury(mercuryRequest *MercuryRequestStruct, url string) (*MercuryResponseStruct, error) {

	var mercuryResponse *MercuryResponseStruct

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(mercuryRequest)
	if err != nil {
		return mercuryResponse, err
	}
	req, _ := json.Marshal(mercuryRequest)

	resp, err := http.Post(
		url,
		"application/json",
		bytes.NewBuffer(req),
	)
	if err != nil {
		return mercuryResponse, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return mercuryResponse, err
	}

	// fmt.Println(string(body)) // will write response in the console

	// unmarshal response and return struct
	json.Unmarshal(body, &mercuryResponse)

	return mercuryResponse, nil
}

// Send content to txpick server
// Response will be an array of text
func SendContentToTxPick(txPickRequest *TxPickRequestStruct, url string) (*TxPickResponseStruct, error) {

	var txPickResponse *TxPickResponseStruct

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(txPickRequest)
	if err != nil {
		log.Fatal(err)
	}

	req, err := json.Marshal(txPickRequest)
	if err != nil {
		return txPickResponse, err
	}

	resp, err := http.Post(
		url,
		// "https://httpbin.org/post", // route for debugging request
		"application/json",
		bytes.NewBuffer(req),
	)
	if err != nil {
		return txPickResponse, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return txPickResponse, err
	}

	// fmt.Println(string(body)) // will write response in the console

	json.Unmarshal(body, &txPickResponse)

	return txPickResponse, nil
}

// Send text array to Aquila hub.
// Response will be an array of vectors.
func SendTextToAquilaHub(a *AquilaHubRequestStruct, url string) (*AquilaHubResponseStruct, error) {

	var aquilaHubResponse *AquilaHubResponseStruct

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(a)
	if err != nil {
		return aquilaHubResponse, err
	}

	req, _ := json.Marshal(a)

	resp, err := http.Post(
		url,
		// "https://httpbin.org/post", // for debugging
		"application/json",
		bytes.NewBuffer(req),
	)
	if err != nil {
		return aquilaHubResponse, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return aquilaHubResponse, err
	}

	// fmt.Println(string(body)) // will write response in the console

	json.Unmarshal(body, &aquilaHubResponse)

	return aquilaHubResponse, nil
}

// Send vectors to Aquila DB
// Response will be an array of ids
func SendVectors(docInsert *DocInsertRequestStruct, url string) (*DocInsertResponseStruct, error) {

	var docInsertResponse *DocInsertResponseStruct

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(docInsert)
	if err != nil {
		return docInsertResponse, err
	}

	req, err := json.Marshal(docInsert)
	if err != nil {
		return docInsertResponse, err
	}

	resp, err := http.Post(
		url,
		// "https://httpbin.org/post", // for debugging
		"application/json",
		bytes.NewBuffer(req),
	)
	if err != nil {
		return docInsertResponse, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return docInsertResponse, err
	}

	fmt.Println(string(body)) // will write response in the console
	json.Unmarshal(body, &docInsertResponse)

	return docInsertResponse, nil
}
