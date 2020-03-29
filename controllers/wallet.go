package controllers

import (
	"hd-wallet/common"
	"hd-wallet/display"
	"hd-wallet/repo"
	"net/http"

	"github.com/gin-gonic/gin"
)

//AddNewWallet ...
//Request is master public key
//Response is id of wallet and token for access to wallet
func AddNewWallet(c *gin.Context) {
	request := common.AddNewWalletRequest{}

	err := c.ShouldBind(&request)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Data or data type is invalid",
		})
		return
	}

	account := common.MasterKeys[request.MasterPublicKey]

	wallet, err := account.CreateWallet()

	if wallet != nil {
		common.Wallets[wallet.ID] = wallet
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
		})
		return
	}

	walletDisplay := display.Wallet{
		ID: wallet.ID,
	}

	walletDisplay.AddressesInfo = wallet.Addresses()

	common.MapWalletPubKey[wallet.ID] = request.MasterPublicKey

	c.JSON(200, wallet)
}

//RetrieveAddress ...
//Request is token (or master public key) (optimize later)
//Currency (Bitcoin, Ethereum, ...)
//Response is address info in given blockchain
func RetrieveAddress(c *gin.Context) {
	// Request will be a master public key, wallet id
	// And currency (Bitcoin, Ethereum, ...)

	request := common.RetrieveAddressRequest{}

	err := c.ShouldBind(&request)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Data or data type is invalid",
		})
		return
	}

	wallet := common.Wallets[request.WalletID]
	memonic := common.MasterKeys[common.MapWalletPubKey[request.WalletID]].Mnemonic()

	index := 0

	if wallet != nil {

		if common.Addresses[wallet] == nil {
			common.Addresses[wallet] = make(map[int][]*repo.AddressInfo)
		}

		if common.Addresses[wallet][request.Currency] != nil {
			index = len(common.Addresses[wallet][request.Currency])
		}
	}

	addressInfo, err := wallet.RetrieveAddress(request.Currency, index, memonic)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
		})
		return
	}

	addressDisplay := display.Address{
		ID:       addressInfo.ID,
		Index:    addressInfo.Index(),
		Currency: addressInfo.Currency(),
		Address:  addressInfo.Address(),
	}

	if addressInfo != nil {
		common.Addresses[wallet][request.Currency] =
			append(common.Addresses[wallet][request.Currency], addressInfo)
	}

	c.JSON(200, addressDisplay)

	// Response will be address info in a specific blockchain
}
