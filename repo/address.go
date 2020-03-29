package repo

import (
	"errors"
	"hd-wallet/constant"

	"github.com/foxnut/go-hdwallet"
	"github.com/google/uuid"
)

//AddressInfo ...
type AddressInfo struct {
	ID       string `json:"id,omitempty"`
	index    int
	currency int
	address  string
}

//SetIndex ...
func (addressInfo *AddressInfo) SetIndex(index int) {
	addressInfo.index = index
}

//Index ...
func (addressInfo *AddressInfo) Index() int {
	return addressInfo.index
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
func (addressInfo *AddressInfo) CreateAddressInfo(currency, index int,
	mnemonic string) (*AddressInfo, error) {
	addressInfo.ID = uuid.New().String()
	addressInfo.currency = currency
	var err error

	master, err := hdwallet.NewKey(
		hdwallet.Mnemonic(mnemonic),
	)

	if err != nil {
		return nil, err
	}

	i := uint32(index)

	switch currency {
	case constant.BTC:
		wallet, _ := master.GetWallet(hdwallet.CoinType(hdwallet.BTC), hdwallet.AddressIndex(i))

		address, _ := wallet.GetAddress()
		addressInfo.address = address

	case constant.BCH:
		wallet, _ := master.GetWallet(hdwallet.CoinType(hdwallet.BCH), hdwallet.AddressIndex(i))

		address, _ := wallet.GetAddress()
		addressInfo.address = address

	case constant.DOGE:
		wallet, _ := master.GetWallet(hdwallet.CoinType(hdwallet.DOGE), hdwallet.AddressIndex(i))

		address, _ := wallet.GetAddress()
		addressInfo.address = address

	case constant.ETC:
		wallet, _ := master.GetWallet(hdwallet.CoinType(hdwallet.ETC), hdwallet.AddressIndex(i))

		address, _ := wallet.GetAddress()
		addressInfo.address = address

	case constant.ETH:
		wallet, _ := master.GetWallet(hdwallet.CoinType(hdwallet.ETH), hdwallet.AddressIndex(i))

		address, _ := wallet.GetAddress()
		addressInfo.address = address

	case constant.LTC:
		wallet, _ := master.GetWallet(hdwallet.CoinType(hdwallet.LTC), hdwallet.AddressIndex(i))

		address, _ := wallet.GetAddress()
		addressInfo.address = address

	default:
		err = errors.New("Invalid currency")
	}

	if err == nil {
		addressInfo.index = index
	}

	return addressInfo, err
}
