package main

import (
	log2 "GolangChat/log"
	router2 "GolangChat/router"
	"GolangChat/utils"
)

func newLogger() *log2.Logger {
	return log2.NewLogger()
}

var log = newLogger()

func main() {
	utils.InitConfig()
	utils.InitMySQL()
	utils.InitRedis()
	router := router2.Router()
	err := router.Run("127.0.0.1:8080")
	if err != nil {
		log.Fatal(err.Error())
	}
}
