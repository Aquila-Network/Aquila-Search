package moduledb

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type CreateAquilaResponsStruct struct {
	DatabaseName string `json:"database_name"`
	Success      bool   `json:"success"`
}

type MetadataStructCreateDb struct {
	Name string `json:"name"`
	Age  string `json:"age"`
}

type SchemaStruct struct {
	Description string                 `json:"description"`
	Unique      string                 `json:"unique"`
	Encoder     string                 `json:"encoder"`
	Codelen     int                    `json:"codelen"`
	Metadata    MetadataStructCreateDb `json:"metadata"`
}

type DataStructCreateDb struct {
	Schema SchemaStruct `json:"schema"`
}

type CreateDbRequestStruct struct {
	Data      DataStructCreateDb `json:"data"`
	Signature string             `json:"signature"`
}

func CreateAquilaDatabase(createDb *CreateDbRequestStruct, url string) (*CreateAquilaResponsStruct, error) {

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
