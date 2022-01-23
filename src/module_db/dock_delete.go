package moduledb

import (
	"aquiladb/src/config"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

/*
{
    "data": {
        "ids": [
            "3gwTnetiYJfHTBcqGwoxETLsmmdGYVsd5MRBohuTG22C",
            "BXsbHy9B3tU9zaHwU41jATzDBisNEFa67XKvYZhB2fzQ"
        ],
        "database_name": "BN4Bik3RbaY5mzJS94u8SvjZd1keyjTWaDNF36TjYzj7"
    },
    "signature": "secret"
}
*/

type DeleteDataStruct struct {
	Ids          []string `json:"ids"`
	DatabaseName string   `json:"database_name"`
}

type DocDeleteStruct struct {
	Data      DeleteDataStruct `json:"data"`
	Signature string           `json:"signature"`
}

func DocDelete() []uint8 {

	docDelete := &DocDeleteStruct{
		Data: DeleteDataStruct{
			Ids: []string{
				"3gwTnetiYJfHTBcqGwoxETLsmmdGYVsd5MRBohuTG22C",
				"BXsbHy9B3tU9zaHwU41jATzDBisNEFa67XKvYZhB2fzQ",
			},
			DatabaseName: "BN4Bik3RbaY5mzJS94u8SvjZd1keyjTWaDNF36TjYzj7",
		},
		Signature: "secret",
	}

	data, err := json.Marshal(docDelete)
	var configEnv = config.GlobalConfig

	createURL := fmt.Sprintf("http://%v:%v/db/doc/delete",
		configEnv.AquilaDB.Host,
		configEnv.AquilaDB.AquilaDbPort,
	)

	resp, err := http.Post(
		createURL,
		// "https://httpbin.org/post",
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

	fmt.Println("=========================================")
	fmt.Println(string(body)) // write response in the console

	return body
}
