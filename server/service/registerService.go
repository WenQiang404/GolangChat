package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/smtp"
	"server/utils"
	"strings"
	"time"

	"math/rand"
)

func Register(c *gin.Context) {
	email := c.PostForm("email")
	//saltPwd := c.PostForm("password")

	if !utils.ValidateEmail(email) {
		log.Error("邮箱格式错误")
		c.JSON(200, gin.H{
			"code":    -1,
			"Message": "邮箱格式错误",
		})
	} else {
		//发送邮件
		code := sendEmail(email)
		c.JSON(200, gin.H{
			"code": code,
		})
		fmt.Println("已经成功向客户端发送邮件！")
	}
}

func SendToMail(user, sendUserName, password, host, to, subject, body, mailtype string) error {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	var content_type string
	if mailtype == "html" {
		content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	msg := []byte("To: " + to + "\r\nFrom: " + sendUserName + "<" + user + ">" + "\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, send_to, msg)
	return err
}

func sendEmail(addr string) string {
	user := "wenqiang9946@163.com"
	password := "IZUJCNJYUGFQWEKK"
	host := "smtp.163.com:25"
	to := addr

	subject := "使用Golang发送邮件"
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	vcode := fmt.Sprintf("%06v", rnd.Int31n(1000000))

	body := fmt.Sprintf(
		`
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="iso-8859-15">
			<title>MMOGA POWER</title>
		</head>
		<body>
			<div>
							欢迎注册奶茶店烤红薯聊天室</div>
						<div>
							<p>验证码为：%s</p>
							</div>
		</body>
		</html>`, vcode)

	sendUserName := "大地瓜在奶茶店烤鱼" //发送邮件的人名称
	fmt.Println("send email")
	err := SendToMail(user, sendUserName, password, host, to, subject, body, "html")
	if err != nil {
		log.Error("邮件发送错误：" + err.Error())
	} else {
		log.Info("目标用户：" + addr + " 邮件发送成功！")
	}
	return vcode
}
