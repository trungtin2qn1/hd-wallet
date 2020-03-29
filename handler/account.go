package handler

import "hd-wallet/repo"

//AccountHandler ...
type AccountHandler interface {
	PrivateKey() string
	Mnemonic() string
	CreateWallet() (*repo.Wallet, error)
}
