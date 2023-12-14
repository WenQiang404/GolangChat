package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"time"
)

type MailBoxConf struct {
	Title       string
	Body        string
	ReceiptList []string
	Sender      string
	SPassword   string
	SMTPAddr    string
	SMTPPort    int
}

func EmailCheck(c *gin.Context) {
	email := c.PostForm("email")

	//格式校验
	var emailType = "= /^[^\\s@]+@[^\\s@]+\\.[^\\s@]+$/"
	if email != emailType {
		c.JSON(200, gin.H{
			"code":    -1,
			"Message": "邮箱格式错误",
		})
	} else {

	}
}

func sendEmail(email string) {
	var mailConf MailBoxConf
	mailConf.Title = "验证"
	mailConf.Body = "欢迎注册小豆聊天室，以下是验证码请查收，祝你生活愉快！"
	mailConf.Sender = email
	mailConf.SPassword = ""
	mailConf.SMTPAddr = "smtp.163.com"
	mailConf.SMTPPort = 25

	randNum := rand.New(rand.NewSource(time.Now().UnixNano()))
	vcode := fmt.Sprint("%06v", randNum.Int31n(1000000))

	//发送的内容
	//
}
