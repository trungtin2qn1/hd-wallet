package models

import "github.com/google/uuid"

//Wallet ...
type Wallet struct {
	ID            string `json:"id,omitempty"`
	addressesInfo map[string]*AddressInfo
}

//CreateWallet ...
func (wallet *Wallet) CreateWallet() *Wallet {
	wallet.ID = uuid.New().String()
	wallet.addressesInfo = make(map[string]*AddressInfo)
	return wallet
}

//AddressInfo ...
type AddressInfo struct {
	ID       string `json:"id,omitempty"`
	currency int
	address  string
}

//Currency ...
func (addressInfo *AddressInfo) Currency() int {
	return addressInfo.currency
}

//Address ...
func (addressInfo *AddressInfo) Address() string {
	return addressInfo.address
}

//SetCurrency ...
func (addressInfo *AddressInfo) SetCurrency(currency int) {
	addressInfo.currency = currency
}

//SetAddress ...
func (addressInfo *AddressInfo) SetAddress(address string) {
	addressInfo.address = address
}

//CreateAddressInfo ...
func (addressInfo *AddressInfo) CreateAddressInfo(currency int) *AddressInfo {
	addressInfo.ID = uuid.New().String()
	addressInfo.currency = currency
	return addressInfo
}
