package controllers

import (
	"hd-wallet/common"
	"hd-wallet/models"
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

	wallet := &models.Wallet{}

	wallet = wallet.CreateWallet()

	c.JSON(200, wallet)

	// Response will be wallet info
	// id of the wallet
	// type of wallet
	// private and public key of wallet
	// token (maybe)

}

//RetrieveAddress ...
//Request is token (or master public key) (optimize later)
//Currency (Bitcoin, Ethereum, ...)
//Response is address info in given blockchain
func RetrieveAddress(c *gin.Context) {
	// Request will be a master public key, wallet id
	// And currency (Bitcoin, Ethereum, ...)

	// Response will be address info in a specific blockchain
}
