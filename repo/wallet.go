package repo

import (
	"github.com/google/uuid"
)

//Wallet ...
type Wallet struct {
	ID            string `json:"id,omitempty"`
	addressesInfo map[string]*AddressInfo
}

//Addresses ...
func (wallet *Wallet) Addresses() map[string]*AddressInfo {
	return wallet.addressesInfo
}

//RetrieveAddress ...
func (wallet *Wallet) RetrieveAddress(currency, index int, mnemonic string) (*AddressInfo, error) {
	addressInfo := &AddressInfo{}
	var err error
	addressInfo, err = addressInfo.CreateAddressInfo(currency, index, mnemonic)

	if err != nil {
		return addressInfo, err
	}

	wallet.addressesInfo[addressInfo.ID] = addressInfo
	return addressInfo, nil
}

//CreateWallet ...
func (wallet *Wallet) CreateWallet() *Wallet {
	wallet.ID = uuid.New().String()
	wallet.addressesInfo = make(map[string]*AddressInfo)
	return wallet
}
