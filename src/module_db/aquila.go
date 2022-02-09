package moduledb

import "aquiladb/src/module_db/src"

func InitAquilaDb() *src.AquilaDb {
	wallet := src.Wallet()
	return src.NewAquilaDb(&wallet)
}
