package service

import (
	"GolangChat/modules"
	"GolangChat/utils"
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"math/rand"
	"strconv"
)

// GetUser
// @Tags 获取用户
// @Success 200 {string} json{"code", "message"}
// @Router /user/getUser [get]
func GetUser(c *gin.Context) {
	data := utils.GetUserList()
	c.JSON(200, gin.H{
		"message": data,
	})
}

// CreateUser
// @Summary 新增用户
// @Tags 用户模块
// @param name query string false "用户名"
// @param password query string false "密码"
// @param repassword query string false "确认密码"
// @Success 200 {string} json{"code", "message"}
// @Router /user/createUser [get]
func CreateUser(c *gin.Context) {
	user := modules.UserBasic{}
	name := c.Query("name")
	//fmt.Println("name: " + name)
	user.Name = c.Query("name")
	password := c.Query("password")
	repassword := c.Query("repassword")

	if password != repassword {
		c.JSON(-1, gin.H{
			"message": "两次密码不一致",
		})
		return
	}
	if name == "" {
		c.JSON(-1, gin.H{
			"message": "请输入用户名",
		})
		return
	}
	if password == "" {
		c.JSON(-1, gin.H{
			"message": "请输入密码",
		})
		return
	}
	data := utils.FindUserByName(name)
	if data.Name != "" {
		c.JSON(200, gin.H{
			"message": "用户名已经注册",
		})
		return
	}

	random := fmt.Sprintf("%06d", rand.Int31())
	user.Password = utils.RandomEncrypt(password, random)
	fmt.Println(random)
	user.Random = random
	fmt.Println(user.Random)
	utils.CreateUser(user)
	c.JSON(200, gin.H{
		"message": "添加用户成功",
	})
}

// DeleteUser
// @Summary 删除用户
// @Tags 用户模块
// @param id query string false "ID"
// @Success 200 {string} json{"code", "message"}
// @Router /user/deleteUser [get]
func DeleteUser(c *gin.Context) {
	user := modules.UserBasic{}
	id, _ := strconv.Atoi(c.Query("id"))
	user.ID = uint(id)
	utils.DeleteUser(user)

	c.JSON(200, gin.H{
		"message": "删除用户成功",
	})
}

// UpdateUser
// @Summary 修改用户
// @Tags 用户模块
// @param id formData string false "ID"
// @param name formData string false "name"
// @param password formData string false "password"
// @param email formData string false "email"
// @param phone formData string false "phone"
// @Success 200 {string} json{"code", "message"}
// @Router /user/updateUser [post]
func UpdateUser(c *gin.Context) {
	user := modules.UserBasic{}
	id, _ := strconv.Atoi(c.PostForm("id"))
	user.ID = uint(id)
	user.Name = c.PostForm("name")
	user.Password = c.PostForm("password")
	user.Email = c.PostForm("email")
	user.Phone = c.PostForm("phone")
	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		c.JSON(200, gin.H{
			"message": "修改内容不匹配",
		})
	}
	utils.UpdateUser(user)

	c.JSON(200, gin.H{
		"message": "修改成功",
	})
}

// Login
// @Summary 登录
// @Tags 用户模块
// @param name formData string false "name"
// @param password formData string false "password"
// @Success 200 {string} json{"code", "message"}
// @Router /Login [get]
func Login(c *gin.Context) {
	loginName := c.Query("name")
	LoginPassword := c.Query("password")
	currentUser := utils.FindUserByName(loginName)

	if currentUser.Name == "" {
		c.JSON(1991, gin.H{
			"message": "用户名不存在",
		})
		return
	}

	flag := utils.DeEncyypt(LoginPassword, currentUser.Random, currentUser.Password)
	if !flag {
		c.JSON(199, gin.H{
			"message": "密码错误",
		})
		return
	}
	pwd := utils.RandomEncrypt(LoginPassword, currentUser.Random)
	data := utils.FindUserByNameAndPwd(loginName, pwd)

	c.JSON(200, gin.H{
		"message": data,
	})

}
