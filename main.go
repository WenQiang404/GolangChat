package main

import (
	router2 "GolangChat/router"
	"GolangChat/utils"
	"log"
)

func main() {
	utils.InitConfig()
	utils.InitMySQL()
	router := router2.Router()
	err := router.Run("127.0.0.1:8080")
	if err != nil {
		log.Fatal(err.Error())
	}
}
