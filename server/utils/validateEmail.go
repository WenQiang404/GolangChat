package utils

import "regexp"

func ValidateEmail(email string) bool {
	// 定义邮箱名的正则表达式
	// 此处使用了一个简单的正则表达式，仅匹配常见的邮箱名格式，实际应根据需求使用更严格的正则表达式
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	// 编译正则表达式
	regex := regexp.MustCompile(emailRegex)

	// 使用正则表达式进行匹配
	return regex.MatchString(email)
}
