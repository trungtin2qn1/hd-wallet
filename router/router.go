package router

import (
	"hd-wallet/controllers"

	"github.com/gin-gonic/gin"
)

//InitRouter ...
func InitRouter() {
	router := gin.Default()

	router.POST("/wallet", controllers.AddNewWallet)
	router.POST("/retrieve-address", controllers.RetrieveAddress)
	router.POST("/key", controllers.GenerateMasterKeys)

	router.Run(":6000")
}
