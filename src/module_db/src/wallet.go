package src

type WalletStruct struct {
	SecretKey string
}

func Wallet() WalletStruct {
	return WalletStruct{}
}
