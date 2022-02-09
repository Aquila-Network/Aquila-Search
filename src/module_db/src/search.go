package src

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type SearchAquilaDatabase struct {
}

func NewSearchAquilaDatabase(wallet *WalletStruct) *SearchAquilaDatabase {
	return &SearchAquilaDatabase{}
}

// Searh method
// Response - vectors
func (s *SearchAquilaDatabase) Search(searchBody *SearchAquilaDbRequestStruct, url string) (*DocSearchResponseStruct, error) {

	var docSearchResponse *DocSearchResponseStruct

	// get request
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(searchBody)
	if err != nil {
		return docSearchResponse, err
	}

	req, err := http.NewRequest(http.MethodGet, url, &buf)
	if err != nil {
		return docSearchResponse, err
	}

	// add header to GET request
	req.Header = map[string][]string{
		"Content-Type": {"application/json"},
	}

	resp, err := http.DefaultClient.Do(req)

	/*
		// post request
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
		return docSearchResponse, err
	}

	// fmt.Println(string(body)) // will write response in the console

	json.Unmarshal(body, &docSearchResponse)

	return docSearchResponse, nil
}
