package display

import "hd-wallet/repo"

//Wallet ...
type Wallet struct {
	ID            string                       `json:"id,omitempty"`
	AddressesInfo map[string]*repo.AddressInfo `json:"addresses,omitempty"`
}
