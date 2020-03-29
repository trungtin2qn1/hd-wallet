package repo

import (
	"github.com/google/uuid"
)

//Wallet ...
type Wallet struct {
	ID            string `json:"id,omitempty"`
	addressesInfo map[string]*AddressInfo
}

//RetrieveAddress ...
func (wallet *Wallet) RetrieveAddress(currency int) (*AddressInfo, error) {
	addressInfo := &AddressInfo{}
	addressInfo = addressInfo.CreateAddressInfo(currency)
	wallet.addressesInfo[addressInfo.ID] = addressInfo
	return addressInfo, nil
}

//CreateWallet ...
func (wallet *Wallet) CreateWallet() *Wallet {
	wallet.ID = uuid.New().String()
	wallet.addressesInfo = make(map[string]*AddressInfo)
	return wallet
}

//Whatever ...
func (wallet *Wallet) Whatever() map[string]string {
	m := make(map[string]string)
	return m
}
