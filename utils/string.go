package utils

import (
	"strings"
)

func StringReplace(s string, old, new []string) string {
	if len(old) != len(new) {
		return ""
	}
	var replacer []string
	for k, v := range old {
		replacer = append(replacer, []string{v, new[k]}...)
	}
	res := strings.NewReplacer(replacer...)
	return res.Replace(s)
}

func SlashAdd(s string) string {
	for i := 0; i < len(s); i++ {
		if s[i] == '\\' || s[i] == '"' || s[i] == '\'' {
			s = s[:i] + `\` + s[i:]
			i++
		}
	}
	return s
}

func SlashDel(s string) string {
	for i := 0; i < len(s); i++ {
		if s[i] == '\\' {
			s = s[:i] + s[i+1:]
		}
	}
	return s
}

// FilterAlphaNum remove not alpha, number and _ characters
func FilterAlphaNum(s string) string {
	var ss []byte
	for i := 0; i < len(s); i++ {
		if (s[i] >= '0' && s[i] <= '9') || (s[i] >= 'A' && s[i] <= 'Z') || (s[i] >= 'a' && s[i] <= 'z') || s[i] == '_' {
			ss = append(ss, s[i])
		}
	}
	return string(ss)
}

// CamelToSnake camel to snake
func CamelToSnake(s string) string {
	num := len(s)
	buf := make([]byte, 0, len(s)*2)
	var c uint8
	for i := 0; i < num; i++ {
		c = s[i]
		if c >= 'A' && c <= 'Z' {
			if i > 0 {
				buf = append(buf, '_')
			}
			buf = append(buf, c+32)
		} else {
			buf = append(buf, c)
		}
	}
	return string(buf)
}

// SnakeToCamel snake to camel
func SnakeToCamel(s string) string {
	l := len(s)
	buf := make([]byte, 0, l)
	for i := 0; i < l; i++ {
		c := s[i]
		if i == 0 && c >= 'a' && c <= 'z' {
			buf = append(buf, c-32)
			continue
		}
		if c == '_' && i+1 < l && s[i+1] >= 'a' && s[i+1] <= 'z' {
			buf = append(buf, s[i+1]-32)
			i++
		} else {
			buf = append(buf, c)
		}
	}
	return string(buf)
}
