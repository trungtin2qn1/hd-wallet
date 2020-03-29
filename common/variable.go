package common

import "hd-wallet/handler"

//MasterKeys : Save public key -> accounts
var MasterKeys map[string]handler.AccountHandler

//Accounts : Save id -> accounts
var Accounts map[string]handler.AccountHandler

//Wallets ...
var Wallets map[handler.AccountHandler]handler.WalletHandler

//Init ....
func Init() {
	MasterKeys = make(map[string]handler.AccountHandler)
	Accounts = make(map[string]handler.AccountHandler)
}
