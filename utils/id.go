package utils

import (
	"strconv"
	"strings"
	"time"
)

// IdInc ID无限增长, 以base62增加,6位为一格, 形式:  zzzzzz-zzzzzz-zzzzzz
func IdInc(s []string) []string {
	n := len(s)
	if n == 0 {
		return []string{`1`}
	}
	last := Base62Decode(s[n-1])
	// base63 6位编码最大值
	if last >= 56800235583 {
		s[n-1] = `0`
		if n-1 < 0 {
			return []string{`1`, `0`}
		}
		return append(IdInc(s[:n-1]), `0`)
	}
	s[n-1] = Base62Encode(last + 1)
	return s
}

// IdEncode 将一系列ID通过随机数封装起来
func IdEncode(args ...uint64) (dst string) {
	if len(args) == 0 {
		return ""
	}
	now := time.Now().Format(".000000")
	nowNum, _ := strconv.ParseUint(now[1:], 10, 64)
	dst = Base62Encode(nowNum)
	for _, v := range args {
		dst += "-" + Base62Encode(v+nowNum)
	}
	return
}

func IdDecode(dst string) (args []uint64) {
	s := strings.Split(dst, `-`)
	if len(s) <= 1 {
		return nil
	}
	now := Base62Decode(s[0])
	for _, v := range s[1:] {
		args = append(args, Base62Decode(v)-now)
	}
	return
}
