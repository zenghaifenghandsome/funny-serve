package main

import (
	"funny-serve/config"
	"funny-serve/model"
	"funny-serve/router"
)

func main() {
	config.InitConfig()
	model.InitDb()
	//model.InitMongoDb()
	router.InitRouter()
}
