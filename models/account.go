package models

import "hd-wallet/repo"

//Account ...
type Account struct {
	ID         string `json:"id,omitempty"`
	privateKey string
	publicKey  string
	mnemonic   string
}

//PublicKey ...
func (account *Account) PublicKey() string {
	return account.publicKey
}

//PrivateKey ...
func (account *Account) PrivateKey() string {
	return account.privateKey
}

//Mnemonic ...
func (account *Account) Mnemonic() string {
	return account.mnemonic
}

//SetPublicKey ...
func (account *Account) SetPublicKey(publicKey string) {
	account.publicKey = publicKey
}

//SetPrivateKey ...
func (account *Account) SetPrivateKey(privateKey string) {
	account.privateKey = privateKey
}

//SetMnemonic ...
func (account *Account) SetMnemonic(mnemonic string) {
	account.mnemonic = mnemonic
}

//CreateWallet ...
func (account *Account) CreateWallet() (*repo.Wallet, error) {
	wallet := &repo.Wallet{}
	wallet = wallet.CreateWallet()

	return wallet, nil
}
