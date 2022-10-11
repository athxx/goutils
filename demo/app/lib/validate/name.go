package validate

import (
	"strings"

	"svr/app/md"
	"svr/app/util"
)

func Username(name string) string {
	// 转换成UTF8字符长度
	name = strings.TrimSpace(name)
	if util.EmojiCheck(name) {
		return ""
	}
	l := len(name)
	if l < 4 || l > 50 || util.EmojiCheck(name) {
		return ""
	}
	uname := []rune(name)
	for i := 0; i < len(uname); i++ {
		switch uname[i] {
		case '\t', '\n', '\v', '\f', '\r', 0x85, 0xA0:
			return ""
		}
	}
	if len(name) >= len(md.USER_PREFIX) && name[:len(md.USER_PREFIX)] == md.USER_PREFIX {
		return ""
	}
	return name
}

func Nickname(name string) string {
	// 转换成UTF8字符长度
	name = strings.TrimSpace(name)
	l := len(name)
	if l < 4 || l > 50 {
		return ""
	}
	uname := []rune(name)
	for i := 0; i < len(uname); i++ {
		switch uname[i] {
		case '\t', '\n', '\v', '\f', '\r', 0x85, 0xA0:
			return ""
		}
	}
	return name
}
