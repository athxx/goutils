package utils

import (
	"regexp"
)

func ValidateChinaPhone(phone string) bool {
	// china phone rule 1[3-9][\d]{9}
	n := StrToInt64(phone)
	return n >= 13000000000 && n <= 19999999999
}

func ValidateEmail(email string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}
