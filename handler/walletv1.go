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

	if len(strs) == 1 {
		return nil, errors.New("Error in path input")
	}

	for _, chain := range chains {
		addressv1 := display.AddressV1{}
		address := ""

		masterkey, err := hdkeychain.NewKeyFromString(xpubKey)
		if err != nil {
			return nil, err
		}

		acct := &hdkeychain.ExtendedKey{}

		if chain == "BTC" {

			// acct := &hdkeychain.ExtendedKey{}

			for i := 1; i < len(strs); i++ {
				childIndex, err := strconv.Atoi(strs[i])
				if err != nil {
					return nil, err
				}

				acct, err = masterkey.Child(uint32(childIndex))
				if err != nil {
					return nil, err
				}

				masterkey = acct
			}
		}

		if chain == "ETH" {
			acc44H, err := masterkey.Child(hdkeychain.HardenedKeyStart + 44)
			if err != nil {
				continue
			}

			acc44H60H, err := acc44H.Child(hdkeychain.HardenedKeyStart + 60)
			if err != nil {
				continue
			}

			acc44H60H0H, err := acc44H60H.Child(hdkeychain.HardenedKeyStart + 0)
			if err != nil {
				continue
			}

			acc44H60H0H0, err := acc44H60H0H.Child(0)
			if err != nil {
				continue
			}

			for i := 1; i < len(strs); i++ {
				childIndex, err := strconv.Atoi(strs[i])
				if err != nil {
					return nil, err
				}

				acct, err = acc44H60H0H0.Child(uint32(childIndex))
				if err != nil {
					return nil, err
				}

				masterkey = acct
			}

		}

		addr, err := acct.Address(&chaincfg.MainNetParams)
		if err != nil {
			return nil, err
		}

		address = addr.String()
		log.Println("address:", address)

		addressv1.Type = chain
		addressv1.Address = address
		walletV1.Addresses = append(walletV1.Addresses, addressv1)
	}

	return walletV1, nil
}
