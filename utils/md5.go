package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5(str string) string {
	return Md5Byte([]byte(str))
}

func Md5Byte(b []byte) string {
	h := md5.New()
	h.Write(b)
	return hex.EncodeToString(h.Sum(nil))
}
