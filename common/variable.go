package common

import (
	"hd-wallet/handler"
	"hd-wallet/repo"
)

// accounts -> wallets -> addresses (currencies)

//MasterKeys : Save public key -> accounts
var MasterKeys map[string]handler.AccountHandler

//Accounts : Save id -> accounts
var Accounts map[string]handler.AccountHandler

//Wallets : Save id -> wallet
var Wallets map[string]*repo.Wallet

//Addresses : Save wallet -> currency -> addressInfo
var Addresses map[*repo.Wallet]map[int][]*repo.AddressInfo

//MapWalletPubKey : Save wallet id -> public key
var MapWalletPubKey map[string]string

//Init ....
func Init() {
	MasterKeys = make(map[string]handler.AccountHandler)
	Accounts = make(map[string]handler.AccountHandler)
	Wallets = make(map[string]*repo.Wallet)
	Addresses = make(map[*repo.Wallet]map[int][]*repo.AddressInfo)
	MapWalletPubKey = make(map[string]string)
}
