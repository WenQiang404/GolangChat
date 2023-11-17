package router

import (
	"GolangChat/docs"
	"GolangChat/service"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.GET("/index", service.GetIndex)

	//User相关
	r.GET("/user/getUser", service.GetUser)
	r.GET("/user/createUser", service.CreateUser)
	r.GET("user/deleteUser", service.DeleteUser)
	r.GET("user/updateUser", service.UpdateUser)
	return r
}
