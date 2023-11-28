package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	log2 "server/log"
	"server/utils"
	"time"
)

func newLogger() *log2.Logger {
	return log2.NewLogger()
}

var log = newLogger()

// 防止跨域站点伪请求
var upGrade = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func SendMsg(c *gin.Context) {
	ws, err := upGrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Error(err.Error())
	}
	defer func(ws *websocket.Conn) {
		err = ws.Close()
		if err != nil {
			log.Error(err.Error())
		}
	}(ws)
	MsgHandler(ws, c)
}

func MsgHandler(ws *websocket.Conn, c *gin.Context) {
	msg, err := utils.Subscribe(c, utils.PublishKey)
	if err != nil {
		log.Error(err.Error())
	}
	now := time.Now().Format("2006-01-03 15:01:01")
	m := fmt.Sprintf("[ws][%s]:%s", now, msg)
	err = ws.WriteMessage(1, []byte(m))
	if err != nil {
		log.Error(err.Error())
	}

}
