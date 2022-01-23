package moduledb

import (
	"aquiladb/src/config"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ResponseAquilaDbStruct struct {
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

type CreateDbStruct struct {
	Data      DataStructCreateDb `json:"data"`
	Signature string             `json:"signature"`
}

func CreateAquilaDatabase(createDb *CreateDbStruct) *ResponseAquilaDbStruct {

	data, err := json.Marshal(createDb)
	var configEnv = config.GlobalConfig

	createURL := fmt.Sprintf("http://%v:%v/db/create",
		configEnv.AquilaDB.Host,
		configEnv.AquilaDB.AquilaDbPort,
	)

	resp, err := http.Post(
		createURL,
		"application/json",
		bytes.NewBuffer(data),
	)
	if err != nil {
		print(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		print(err)
	}

	var responseAquilaDb *ResponseAquilaDbStruct
	json.Unmarshal(body, &responseAquilaDb)

	fmt.Println(string(body)) // write response in the console

	return responseAquilaDb
}
