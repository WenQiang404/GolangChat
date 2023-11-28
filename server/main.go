package main

import (
	log2 "server/log"
	"server/router"
	"server/utils"
)

func newLogger() *log2.Logger {
	return log2.NewLogger()
}

var log = newLogger()

func main() {
	_ = utils.InitConfig()
	_ = utils.InitMySQL()
	_ = utils.InitRedis()
	r := router.Router()
	err := r.Run("127.0.0.1:8080")
	if err != nil {
		log.Fatal(err.Error())
	}
}
