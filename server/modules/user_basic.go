package modules

import (
	"gorm.io/gorm"
)

type UserBasic struct {
	gorm.Model
	Identity      string
	Name          string
	Password      string
	Phone         string `valid:"matches(^1[3-9]{1}\\d{9}$)"`
	Email         string `valid:"email"`
	ClientIp      string
	ClientPort    string
	LoginTime     string
	HeartBeatTime string
	LoginOutTime  string
	IsLogOut      bool
	DeviceInfo    string
	Random        string
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}
