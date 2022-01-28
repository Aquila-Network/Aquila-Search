package moduledb

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type DeleteDataStruct struct {
	Ids          []string `json:"ids"`
	DatabaseName string   `json:"database_name"`
}

type DocDeleteRequestStruct struct {
	Data      DeleteDataStruct `json:"data"`
	Signature string           `json:"signature"`
}

type DocDeleteResponseStruct struct {
	Ids     []string `json:"ids"`
	Success bool     `json:"success"`
}

// Delelete Document
// Deleted ids in response
func DocDelete(docDelete *DocDeleteRequestStruct, url string) (*DocDeleteResponseStruct, error) {

	var docDeleteResponse *DocDeleteResponseStruct

	data, err := json.Marshal(docDelete)

	resp, err := http.Post(
		url,
		// "https://httpbin.org/post", // route for debugging
		"application/json",
		bytes.NewBuffer(data),
	)
	if err != nil {
		return docDeleteResponse, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return docDeleteResponse, err
	}

	// fmt.Println(string(body)) // write response in the console

	json.Unmarshal(body, &docDeleteResponse)

	return docDeleteResponse, nil
}
