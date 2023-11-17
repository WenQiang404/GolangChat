package utils

import (
	"GolangChat/modules"
	"fmt"
	"gorm.io/gorm"
)

func GetUserList() []*modules.UserBasic {
	data := make([]*modules.UserBasic, 10)
	DB.Find(&data)

	for _, v := range data {
		fmt.Println(v)
	}
	return data
}

func CreateUser(user modules.UserBasic) *gorm.DB {

	return DB.Create(&user)
}

func DeleteUser(user modules.UserBasic) *gorm.DB {
	return DB.Delete(&user)
}

func UpdateUser(user modules.UserBasic) *gorm.DB {
	return DB.Model(&user).Updates(modules.UserBasic{Name: user.Name, Password: user.Password})
}
