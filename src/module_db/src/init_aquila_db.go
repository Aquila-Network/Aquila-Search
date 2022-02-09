package src

type CreateDbInterface interface {
	Create(createDb *CreateDbRequestStruct, url string) (*CreateAquilaResponsStruct, error)
}

type DocInsertInterface interface {
	SendHTMLForParsingToMercury(mercuryRequest *MercuryRequestStruct, url string) (*MercuryResponseStruct, error)
	SendContentToTxPick(txPickRequest *TxPickRequestStruct, url string) (*TxPickResponseStruct, error)
	SendTextToAquilaHub(a *AquilaHubRequestStruct, url string) (*AquilaHubResponseStruct, error)
	SendVectors(docInsert *DocInsertRequestStruct, url string) (*DocInsertResponseStruct, error)
}

type DocSearchInterface interface {
	Search(searchBody *SearchAquilaDbRequestStruct, url string) (*DocSearchResponseStruct, error)
}

type DocDeleteInterface interface {
	DocDelete(docDelete *DocDeleteRequestStruct, url string) (*DocDeleteResponseStruct, error)
}

type AquilaDb struct {
	CreateDbInterface
	DocInsertInterface
	DocSearchInterface
	DocDeleteInterface
}

func NewAquilaDb(wallet *WalletStruct) *AquilaDb {
	return &AquilaDb{
		CreateDbInterface:  NewCreateAquilaDatabase(wallet),
		DocInsertInterface: NewDocInsertAquilaDatabase(wallet),
		DocSearchInterface: NewSearchAquilaDatabase(wallet),
		DocDeleteInterface: NewDocDeleteAquilaDatabase(wallet),
	}
}
