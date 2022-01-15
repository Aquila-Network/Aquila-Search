package moduledb

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

/*
	{
		"data": {
			"schema": {
				"description": "this is my database",
				"unique": "r8and0mseEd901",
				"encoder": "strn:msmarco-distilbert-base-tas-b",
				"codelen": 768,
				"metadata": {
					"name": "string",
					"age": "number"
				}
			}
		},
		"signature": "secret"
	}
*/

type MetadataStruct struct {
	Name string `json:"name"`
	Age  string `json:"age"`
}

type SchemaStruct struct {
	Description string         `json:"description"`
	Unique      string         `json:"unique"`
	Encoder     string         `json:"encoder"`
	Codelen     string         `json:"codelen"`
	Metadata    MetadataStruct `json:"metadata"`
}

type DataStruct struct {
	Schema SchemaStruct `json:"schema"`
}

type CreateDbStruct struct {
	Data      DataStruct `json:"data"`
	Signature string     `json:"signature"`
}

type Fruit struct {
	Name  string `json:"name"`
	Color string `json:"color"`
}

func CreateAquilaDatabase() {

	createDb := &CreateDbStruct{
		Data: DataStruct{
			Schema: SchemaStruct{
				Description: "this is my database",
				Unique:      "r8and0mseEd901",
				Encoder:     "strn:msmarco-distilbert-base-tas-b",
				Codelen:     "768",
				Metadata: MetadataStruct{
					Name: "string",
					Age:  "number",
				},
			},
		},
		Signature: "secret",
	}

	data, err := json.Marshal(createDb)
	if err != nil {
		log.Fatal(err)
	}
	reader := bytes.NewReader(data)
	fmt.Println(reader)
	resp, err := http.Post(
		"https://httpbin.org/post",
		"application/json",
		reader,
	)
	if err != nil {
		print(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		print(err)
	}
	fmt.Println(string(body)) // will sent response in the console
}
