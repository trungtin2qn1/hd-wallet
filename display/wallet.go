package display

import "hd-wallet/repo"

//Wallet ...
type Wallet struct {
	ID            string                       `json:"id,omitempty"`
	AddressesInfo map[string]*repo.AddressInfo `json:"addresses,omitempty"`
}

//WalletV1 ...
type WalletV1 struct {
	Addresses []AddressV1 `json:"addresses,omitempty"`
}

//AddressV1 ...
type AddressV1 struct {
	Type    string `json:"type,omitempty"`
	Address string `json:"address,omitempty"`
}
