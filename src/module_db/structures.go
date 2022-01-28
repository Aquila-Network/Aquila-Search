package moduledb

// =====================================
// Mercury sturctures
// =====================================

type MercuryRequestStruct struct {
	Url  string `json:"url"`
	Html string `json:"html"`
}

// -----------------------------

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
// Doc insert struct
// =====================================

type MetadataStructDocInsert struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type PayloadStruct struct {
	Metadata MetadataStructDocInsert `json:"metadata"`
	Code     []float64               `json:"code"`
}

type DocsStruct struct {
	Payload PayloadStruct `json:"payload"`
}

type DatatDocInsertStruct struct {
	Docs         []DocsStruct `json:"docs"`
	DatabaseName string       `json:"database_name"`
}

type DocInsertRequestStruct struct {
	Data      DatatDocInsertStruct `json:"data"`
	Signature string               `json:"signature"`
}

// -----------------------

type DocInsertResponseStruct struct {
	Ids     []string `json:"ids"`
	Success bool     `json:"success"`
}

// =====================================
// TxPick
// =====================================

type TxPickRequestStruct struct {
	Url  string `json:"url"`
	Html string `json:"html"`
}

// ---------------------------------

// Response txpick

type TxPickResponseStruct struct {
	Result  []string `json:"result"`
	Success bool     `json:"success"`
}

// =====================================
//  Aquila Hub
// =====================================

type AquilaDataRequestStruct struct {
	Text         []string `json:"text"`
	DatabaseName string   `json:"databaseName"`
}

type AquilaHubRequestStruct struct {
	Data AquilaDataRequestStruct `json:"data"`
}

// --------------------------------

// Response Aquila Hub

type AquilaHubResponseStruct struct {
	Vectors [][]float64
	Success bool
}

// =====================================
// Db Search:
// =====================================

type DataSearchStruct struct {
	Matrix       [][]float64 `json:"matrix"`
	K            int         `json:"k"`
	R            int         `json:"r"`
	DatabaseName string      `json:"database_name"`
}

type SearchAquilaDbRequestStruct struct {
	Data DataSearchStruct `json:"data"`
}

// --------------------------------

type MetadataSearchStruct struct {
	Age  int
	Name string
}

type DocSearchData struct {
	Cid      string
	Id       int
	Code     []float64
	Metadata MetadataSearchStruct
}

type DocSearchResponseStruct struct {
	Dist [][]float64
	Docs [][]DocSearchData
}
