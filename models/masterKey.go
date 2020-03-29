package models

import (
	"hd-wallet/common"
	"log"

	"github.com/google/uuid"
	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
)

//GenerateMasterKey ...
//generate master key
func GenerateMasterKey(password string) (*Account, error) {
	// Generate a mnemonic for memorization or user-friendly seeds
	entropy, err := bip39.NewEntropy(128)
	if err != nil {
		return nil, err
	}
	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		return nil, err
	}

	seeds := bip39.NewSeed(mnemonic, password)

	// Create master private key from seed
	computerVoiceMasterKey, err := bip32.NewMasterKey(seeds)

	if err != nil {
		panic(err)
	}

	log.Println(computerVoiceMasterKey.PublicKey().String())
	log.Println(string(computerVoiceMasterKey.String()))

	account := &Account{
		ID: uuid.New().String(),
	}

	account.SetPrivateKey(string(computerVoiceMasterKey.String()))
	account.SetPublicKey(computerVoiceMasterKey.PublicKey().String())
	account.SetMnemonic(mnemonic)

	common.MasterKeys[account.PublicKey()] = account
	common.Accounts[account.ID] = account

	return account, nil
}
