package handler

//AccountHandler ...
type AccountHandler interface {
	PrivateKey() string
	Mnemonic() string
}
