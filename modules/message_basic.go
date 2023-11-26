package modules

import "gorm.io/gorm"

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
