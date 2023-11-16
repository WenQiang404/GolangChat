package modules

import (
	"gorm.io/gorm"
	"time"
)

type UserBasic struct {
	gorm.Model
	Identity      string
	Name          string
	Password      string
	Phone         string
	Email         string
	ClientIp      string
	ClientPort    string
	LoginTime     time.Time
	HeartBeatTime time.Time
	LoginOutTime  time.Time `gorm:column:login_out_time`
	IsLogOut      bool
	DeviceInfo    string
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}
