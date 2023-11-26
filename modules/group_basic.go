package modules

//群信息
import "gorm.io/gorm"

type GroupBasic struct {
	gorm.Model
	Name   string //群聊名称
	OwerId string //群主
	Type   string
	Icon   string
	Desc   string
}

func (table *GroupBasic) TableName() string {
	return "group_basic"
}
