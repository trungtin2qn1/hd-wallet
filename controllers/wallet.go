package controllers

import (
	"hd-wallet/common"
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

	adressInfo, err := wallet.RetrieveAddress(request.Currency)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
		})
		return
	}

	c.JSON(200, adressInfo)

	// Response will be address info in a specific blockchain
}
