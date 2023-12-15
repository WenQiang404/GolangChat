package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/gomail.v2"
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
	html := fmt.Sprintf(
		`<div>
					欢迎注册聊天室
				</div>
				<div>
					<p>验证码为：%s</p>
					</div>`, vcode)
	m := gomail.NewMessage()
	m.SetHeader(`From`, mailConf.Sender, "小地瓜")
	m.SetHeader(`To`, mailConf.ReceiptList...)
	m.SetHeader(`Subject`, mailConf.Title)
	m.SetBody(`text/html`, html)
	err := gomail.NewDialer(mailConf.SMTPAddr, mailConf.SMTPPort, mailConf.Sender, mailConf.SPassword)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("success")
}
