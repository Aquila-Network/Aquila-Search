package moduledb

import (
	"aquiladb/src/config"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Search() []uint8 {
	var configEnv = config.GlobalConfig

	createURL := fmt.Sprintf("http://%v:%v/db/search",
		configEnv.AquilaDB.Host,
		configEnv.AquilaDB.AquilaDbPort,
	)

	fmt.Println("=====================================")
	fmt.Println(createURL)

	// pattern
	// https://www.geeksforgeeks.org/slice-of-slices-in-golang/
	matrix := make([][]float64, 1)
	matrix[0] = make([]float64, 1)
	matrix[0] = []float64{
		-0.01806008443236351, -0.17380790412425995, 0.03992759436368942, 0.43514639139175415,
	}

	searchBody := &SearchAquilaDbStruct{
		Data: DataSearchStruct{
			Matrix:       matrix,
			K:            10,
			DatabaseName: "BN4Bik3RbaY5mzJS94u8SvjZd1keyjTWaDNF36TjYzj7",
		},
	}

	// get
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(searchBody)
	req, err := http.NewRequest(http.MethodGet, createURL, &buf)
	if err != nil {
		panic(err)
	}
	fmt.Println("=====================================")
	fmt.Printf("%+v\n", req)
	fmt.Println("=====================================")
	fmt.Println(req)

	resp, err := http.DefaultClient.Do(req)

	/*
		// post
		var buf bytes.Buffer
		err := json.NewEncoder(&buf).Encode(searchBody)
		if err != nil {
			log.Fatal(err)
		}

		req, _ := json.Marshal(searchBody)

		resp, err := http.Post(
			// createURL,
			"https://httpbin.org/post",
			"application/json",
			bytes.NewBuffer(req),
		)
		if err != nil {
			print(err)
		}
		defer resp.Body.Close()
	*/
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		print(err)
	}

	fmt.Println(string(body)) // will write response in the console

	return body
}
