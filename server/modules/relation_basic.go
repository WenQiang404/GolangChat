package modules

import (
	"gorm.io/gorm"
)

// 人员关系
type RelationBasic struct {
	gorm.Model
	OwerId   uint //谁的关系
	TargetId uint //对应
	Type     int  //类型 0 1 2
	Desc     string
}

func (re *RelationBasic) TableName() string {
	return "Relation"
}
