package utils

import (
	"GolangChat/modules"
	"fmt"
)

func GetUserList() []*modules.UserBasic {
	data := make([]*modules.UserBasic, 10)
	DB.Find(&data)

	for _, v := range data {
		fmt.Println(v)
	}
	return data
}
