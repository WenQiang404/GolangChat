package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
)

// 小写
func Md5Encode(data string) string {
	lock := md5.New()
	lock.Write([]byte(data))
	tempStr := lock.Sum(nil)
	return hex.EncodeToString(tempStr)
}

// 大写
func MD5Encode(data string) string {
	return strings.ToUpper(Md5Encode(data))
}

func RandomEncrypt(pwd, random string) string {
	return Md5Encode(pwd + random)
}

func DeEncyypt(pwd, random string, key string) bool {
	md := Md5Encode(pwd + random)
	fmt.Println(md + "            " + key)
	return Md5Encode(pwd+random) == key
}
