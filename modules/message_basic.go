package modules

import (
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
	FromId   uint   //发送者
	TargetId uint   //接收者
	Type     string //消息类型 群聊，私聊
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

var rwMutex = sync.RWMutex{}

func Chat(writer http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	Id := query.Get("userId")
	userId, _ := strconv.ParseInt(Id, 10, 64)
	targetId := query.Get("targetId")
	context := query.Get("context")
	token := query.Get("token")
	msgType := query.Get("type")
	var isvalid = true //check Token（）
	conn, err := (&websocket.Upgrader{
		//token校验
		CheckOrigin: func(r *http.Request) bool {
			return isvalid
		},
	}).Upgrade(writer, request, nil)
	if err != nil {
		fmt.Println(err)
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
	rwMutex.Lock()
	clientMap[userId] = client
	rwMutex.Unlock()

	//发送消息
	go sendProc(client)
	//接收消息
	go receiveProc(client)
}

func sendProc(cli *Client) {
	for {
		select {
		case data := <-cli.DataQueue:
			err := cli.Conn.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}
func receiveProc(cli *Client) {
	for {
		_, data, err := cli.Conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}
		broadMsg(data)
		fmt.Println("[ws] >>>>", data)
	}
}

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
		fmt.Println(err)
	}

	for {
		select {
		case data := <-udpsendChan:
			_, err := conn.Write(data)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}
