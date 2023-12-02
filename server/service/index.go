package service

import "github.com/gin-gonic/gin"

// GetIndex
// @Tags 首页
// @Success 200 {string} chatRoom
// @Router /index [get]
func GetIndex(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "chatRoom!",
	})
}
