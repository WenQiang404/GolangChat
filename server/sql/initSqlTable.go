package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	modules2 "server/modules"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:wenqiang@tcp(101.133.169.145:3306)/golangChat?charset=utf8mb3&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 迁移 schema
	//db.AutoMigrate(&modules.UserBasic{})		//user表
	//db.AutoMigrate(&modules.MessageBasic{})		//message表
	//db.AutoMigrate(&modules2.RelationBasic{}) //relation表
	//db.AutoMigrate(&modules2.GroupBasic{})    //group表
	// Create
	user := &modules2.UserBasic{}
	user.Name = "cwq"
	db.Create(user)

	// Read
	// 根据整型主键查找
	fmt.Println(db.First(user, 1)) //db.First(user, "code = ?", "D42") // 查找 code 字段值为 D42 的记录

	// Update - 将 product 的 price 更新为 200
	db.Model(user).Update("PassWord", "1234")
	// Update - 更新多个字段
	//db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // 仅更新非零值字段
	//db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - 删除 product
	//db.Delete(&product, 1)
}
