package router

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"server/docs"
	"server/service"
	"server/utils"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.Use(utils.Cors()) //解决跨域问题
	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	//static

	//首页
	r.GET("/index", service.GetIndex)

	//User相关
	r.GET("/user/getUser", service.GetUser)
	r.GET("/user/createUser", service.CreateUser)
	r.GET("/user/deleteUser", service.DeleteUser)
	r.POST("/user/updateUser", service.UpdateUser)
	r.POST("/Login", service.Login)

	//Message相关
	r.GET("/user/sendMsg", service.SendMsg)
	r.GET("user/sendUserMsg", service.SendUserMessage)

	return r

}
