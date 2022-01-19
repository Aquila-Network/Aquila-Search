package moduledb

import (
	"aquiladb/src/config"
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

type MetadataStructDocInsert struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type PayloadStruct struct {
	Metadata MetadataStructDocInsert `json:"metadata"`
	Code     []float32               `json:"code"`
}

type DocsStruct struct {
	Payload PayloadStruct `json:"payload"`
}

type DatatDocInsertStruct struct {
	Docs         []DocsStruct `json:"docs"`
	DatabaseName string       `json:"database_name"`
}

type DocInsertStruct struct {
	Data      DatatDocInsertStruct `json:"data"`
	Signature string               `json:"signature"`
}

func DocInsert(jsonDataBytes []byte) []uint8 {
	// docInsert := &DocInsertStruct{
	// 	Data: DatatDocInsertStruct{
	// 		Docs: []DocsStruct{
	// 			{
	// 				Payload: PayloadStruct{
	// 					Metadata: MetadataStructDocInsert{
	// 						Name: "name1",
	// 						Age:  20,
	// 					},
	// 					Code: []float32{0.1, 0.2, 0.3},
	// 				},
	// 			},
	// 			{
	// 				Payload: PayloadStruct{
	// 					Metadata: MetadataStructDocInsert{
	// 						Name: "name2",
	// 						Age:  20,
	// 					},
	// 					Code: []float32{0.4, 0.5, 0.6},
	// 				},
	// 			},
	// 		},
	// 		DatabaseName: "BN4Bik3RbaY5mzJS94u8SvjZd1keyjTWaDNF36TjYzj7",
	// 	},
	// 	Signature: "secret",
	// }

	// data, err := json.Marshal(docInsert)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// another way to pass reader to post
	// reader := bytes.NewReader(data)

	var configEnv = config.GlobalConfig

	createURL := fmt.Sprintf("http://%v:%v/db/doc/insert",
		configEnv.AquilaDB.Host,
		configEnv.AquilaDB.Port,
	)

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

/*
{
   "data": {
       "docs": [
           {
           "payload":
               {
                   "metadata": {
                       "name":"name1",
                       "age": 20
                   },
                   "code": [0.1, 0.2, 0.3]
               }
           },
           {
           "payload":
               {
                   "metadata": {
                       "name":"name2",
                       "age": 30
                   },
                   "code": [0.4, 0.5, 0.6]
               }
           }
       ],
       "database_name": "BN4Bik3RbaY5mzJS94u8SvjZd1keyjTWaDNF36TjYzj7"
   },
   "signature": "secret"
}
*/
