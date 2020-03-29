package controllers

import (
	"hd-wallet/common"
	"hd-wallet/display"
	"hd-wallet/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

//GenerateMasterKeys ...
//Request is null
//Response is pair private and public key
func GenerateMasterKeys(c *gin.Context) {
	request := common.GenerateMasterKeyRequest{}
	err := c.ShouldBind(&request)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Data or data type is invalid",
		})
		return
	}

	account, err := models.GenerateMasterKey(request.Password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
	}

	accountDisplay := display.Account{
		ID:         account.ID,
		Mnemonic:   account.Mnemonic(),
		PrivateKey: account.PrivateKey(),
		PublicKey:  account.PublicKey(),
	}

	// token, err := jwt.IssueToken(account.ID)
	// accountDisplay.Token = token

	c.JSON(200, accountDisplay)
}
