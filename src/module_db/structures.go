package moduledb

// =====================================

// Doc insert struct

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

type DocInsertResponseStruct struct {
	Ids     []string `json:"ids"`
	Success bool     `json:"success"`
}

// =====================================

type MercuryDataStruct struct {
	Title         string `json:"title"`
	Author        string `json:"author"`
	DatePublished string `json:"date_published"`
	Dek           string `json:"dek"`
	LeadImageUrl  string `json:"lead_image_url"`
	Content       string `json:"content"`
	NextPageUrl   string `json:"next_page_url"`
	Url           string `json:"url"`
	Domain        string `json:"domain"`
	Except        string `json:"except"`
	WordCount     int    `json:"word_count"`
	Direction     string `json:"direction"`
	TotalPages    int    `json:"total_pages"`
	RenderedPages int    `json:"rendered_pages"`
}

type MercuryResponseStruct struct {
	Data MercuryDataStruct `json:"data"`
}

// =====================================

// Request txpick

type TxPickRequestStruct struct {
	Url  string `json:"url"`
	Html string `json:"html"`
}

// =====================================

// Response txpick

type TxPickResponseStruct struct {
	Result  []string `json:"result"`
	Success bool     `json:"success"`
}

// =====================================

// Request Aquila Hub

type AquilaDataRequestStruct struct {
	Text         []string `json:"text"`
	DatabaseName string   `json:"databaseName"`
}

type AquilaHubRequestStruct struct {
	Data AquilaDataRequestStruct `json:"data"`
}

// Response Aquila Hub

type AquilaHubResponseStruct struct {
	Vectors []float32
	Success bool
}

// =====================================

// Db Search:

type DataSearchStruct struct {
	Matrix       [][]float64 `json:"matrix"`
	K            int         `json:"k"`
	R            int         `json:"r"`
	DatabaseName string      `json:"database_name"`
}

type SearchAquilaDbStruct struct {
	Data DataSearchStruct `json:"data"`
}
