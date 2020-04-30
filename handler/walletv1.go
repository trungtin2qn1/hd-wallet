package handler

import (
	"errors"
	"hd-wallet/display"
	"log"
	"strconv"
	"strings"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil/hdkeychain"
)

var chains = []string{"ETH", "BTC"}

//GenerateWallet ...
func GenerateWallet(xpubKey, path string) (*display.WalletV1, error) {

	walletV1 := &display.WalletV1{}

	strs := strings.Split(path, "/")

	if len(strs) != 2 {
		return nil, errors.New("Error in path input")
	}

	childIndex, err := strconv.Atoi(strs[1])
	if err != nil {
		return nil, err
	}

	for _, chain := range chains {
		addressv1 := display.AddressV1{}
		address := ""
		if chain == "BTC" {
			masterkey, err := hdkeychain.NewKeyFromString(xpubKey)
			if err != nil {
				return nil, err
			}

			acct, err := masterkey.Child(uint32(childIndex))
			if err != nil {
				return nil, err
			}

			addr, err := acct.Address(&chaincfg.MainNetParams)
			if err != nil {
				return nil, err
			}

			address = addr.String()

			log.Println("address:", address)

		}
		if chain == "ETH" {

		}

		addressv1.Type = chain
		addressv1.Address = address
		walletV1.Addresses = append(walletV1.Addresses, addressv1)
	}

	return walletV1, nil
}
