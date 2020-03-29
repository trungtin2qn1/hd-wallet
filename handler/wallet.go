package handler

//WalletHandler ...
type WalletHandler interface {
	AddressInfoHandler() map[string]*AddressInfoHandler
}

//AddressInfoHandler ...
type AddressInfoHandler interface {
	Currency() int
	Address() string
}
