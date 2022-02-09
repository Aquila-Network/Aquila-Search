package src

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type CreateAquilaDatabase struct {
}

func NewCreateAquilaDatabase(wallet *WalletStruct) *CreateAquilaDatabase {
	return &CreateAquilaDatabase{}
}

func (c *CreateAquilaDatabase) Create(createDb *CreateDbRequestStruct, url string) (*CreateAquilaResponsStruct, error) {

	var responseAquilaDb *CreateAquilaResponsStruct
	data, err := json.Marshal(createDb)

	resp, err := http.Post(
		url,
		"application/json",
		bytes.NewBuffer(data),
	)
	if err != nil {
		return responseAquilaDb, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return responseAquilaDb, err
	}

	json.Unmarshal(body, &responseAquilaDb)
	fmt.Println(string(body)) // write response in the console

	return responseAquilaDb, nil
}
