package main

import (
	"hd-wallet/common"
	"hd-wallet/router"
)

func main() {
	//Init variables
	common.Init()

	//Init database

	//Init router ...
	router.InitRouter()
}
