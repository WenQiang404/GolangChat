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
	return DB.Model(&user).Updates(modules.UserBasic{Name: user.Name, Password: user.Password, Email: user.Email, Phone: user.Phone})
}

func FindUserByName(name string) modules.UserBasic {
	user := modules.UserBasic{}
	DB.Where("name = ?", name).First(&user)
	return user
}

func FindUserByEmail(email string) *gorm.DB {
	user := modules.UserBasic{}
	return DB.Where("email = ?", email).First(&user)
}

func FindUserByPhone(phone string) *gorm.DB {
	user := modules.UserBasic{}
	return DB.Where("phone = ?", phone).First(&user)
}
func FindUserByNameAndPwd(name, password string) modules.UserBasic {
	user := modules.UserBasic{}
	DB.Where("name = ? and pass_word = ?", name, password).First(&user)
	return user
}
