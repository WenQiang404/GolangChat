package service

import (
	"github.com/gin-gonic/gin"
	"server/modules"
	"server/utils"
)

func Register(c *gin.Context) {
	newUser := modules.UserBasic{}
	name := c.PostForm("name")
	encryptPwd := c.PostForm("password")
	email := c.PostForm("email")

	newUser = utils.FindUserByName(name)
	if newUser.Name != "" {
		c.JSON(100, gin.H{
			"message": "用户名已经注册",
			"code":    "-1",
		})
		return
	} else {
		newUser.Name = name
		newUser.Password = encryptPwd
		newUser.Email = email
		utils.CreateUser(newUser)
		c.JSON(200, gin.H{
			"code":    1,
			"message": "注册成功",
		})
		return
	}
}
