package modules

import (
	log2 "GolangChat/log"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"gopkg.in/fatih/set.v0"
	"gorm.io/gorm"
	"net"
	"net/http"
	"strconv"
	"sync"
)

// 消息表
type MessageBasic struct {
	gorm.Model
	UserId   int64  //发送者
	TargetId int64  //接收者
	Type     int    //消息类型 群聊，私聊
	Media    int    //消息类型 文字 图片
	Content  string //消息内容
	Pic      string
	Url      string
	Desc     string
	Amount   int //其他字数统计

}

func (msg *MessageBasic) Message() string {
	return "message"
}

type Client struct {
	Conn      *websocket.Conn
	DataQueue chan []byte
	GroupSets set.Interface
}

var clientMap map[int64]*Client = make(map[int64]*Client, 0)

func newLogger() *log2.Logger {
	return log2.NewLogger()
}

var log = newLogger()

var rwLocker = sync.RWMutex{}

func Chat(writer http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	Id := query.Get("userId")
	userId, _ := strconv.ParseInt(Id, 10, 64)
	//targetId := query.Get("targetId")
	//context := query.Get("context")
	//token := query.Get("token")
	//msgType := query.Get("type")
	var isvalid = true //check Token（）
	conn, err := (&websocket.Upgrader{
		//token校验
		CheckOrigin: func(r *http.Request) bool {
			return isvalid
		},
	}).Upgrade(writer, request, nil)
	if err != nil {
		log.Error("Failed to set the websocket connect! : " + err.Error())
		return
	}

	//获取connection
	client := &Client{
		Conn:      conn,
		DataQueue: make(chan []byte, 100),
		GroupSets: set.New(set.ThreadSafe),
	}

	//用户关系
	//id绑定client
	rwLocker.Lock()
	clientMap[userId] = client
	rwLocker.Unlock()

	//发送消息
	go sendProc(client)
	//接收消息
	go receiveProc(client)
	sendMsg(userId, []byte("Welcome to the lulu chat room !"))
}

func sendProc(cli *Client) {
	for {
		select {
		case data := <-cli.DataQueue:
			err := cli.Conn.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				log.Error("Failed to write message : " + err.Error())
				return
			}
		}
	}
}
func receiveProc(cli *Client) {
	for {
		_, data, err := cli.Conn.ReadMessage()
		if err != nil {
			log.Error("Failed to read message : " + err.Error())
			return
		}
		broadMsg(data)
		fmt.Println("[ws] >>>>", data)
	}
}

var udpsendChan chan []byte = make(chan []byte, 1024)

func broadMsg(data []byte) {
	udpsendChan <- data
}

func init() {
	go udpSendProc()
	go udpReceiveProc()

}

// finish the udp send channel
func udpSendProc() {
	conn, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(10, 200, 20, 178),
		Port: 3000,
	})
	defer conn.Close()
	if err != nil {
		log.Error("Failed to connect the UDP network : " + err.Error())
		return
	}

	for {
		select {
		case data := <-udpsendChan:
			_, err := conn.Write(data)
			if err != nil {
				log.Error(err.Error())
				return
			}
		}
	}
}

func udpReceiveProc() {
	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4zero,
		Port: 3000,
	})
	if err != nil {
		log.Error("Failed to connect UDP," + err.Error())
	}
	defer conn.Close()
	for {
		var buf [512]byte
		n, err := conn.Read(buf[0:])
		if err != nil {
			log.Error("Failed to read buffer, " + err.Error())
			return
		}
		err = dispatch(buf[0:n])
		if err != nil {
			log.Error(err.Error())
			return
		}
	}
}

func dispatch(data []byte) error {
	msg := MessageBasic{}
	err := json.Unmarshal(data, &msg)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	switch msg.Type {
	case 1:
		sendMsg(msg.TargetId, data) //私聊
		//case 2:
		//	sendGroup() //群发
		//case 3:
		//	sendAll() //广播
	}
	return nil
}

func sendMsg(userId int64, msg []byte) {
	rwLocker.RLock()
	node, ok := clientMap[userId]
	rwLocker.RUnlock()
	if ok {
		node.DataQueue <- msg
	}
}
