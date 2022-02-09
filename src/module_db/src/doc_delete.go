package src

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type DocDeleteAquilaDatabase struct {
}

func NewDocDeleteAquilaDatabase(wallet *WalletStruct) *DocDeleteAquilaDatabase {
	return &DocDeleteAquilaDatabase{}
}

// Delelete Document
// Deleted ids in response
func (d *DocDeleteAquilaDatabase) DocDelete(docDelete *DocDeleteRequestStruct, url string) (*DocDeleteResponseStruct, error) {

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
