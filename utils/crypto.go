package utils

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
)

func GetMd5(raw []byte) string {
	h := md5.New()
	h.Write(raw)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func GetBase64Md5(raw []byte) string {
	h := md5.New()
	h.Write(raw)
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
