package utils

import (
	"io"
	"net/http"
	"strings"
)

func IPClient(req *http.Request) string {
	ip := req.Header.Get("X-Forwarded-For")
	if strings.Contains(ip, "127.0.0.1") || ip == `` {
		ip = req.Header.Get("X-Real-ip")
	}
	return ip
}

func IPServer() string {
	resp, _ := http.Get("https://ipw.cn/api/ip/myip")
	body, _ := io.ReadAll(resp.Body)
	resp.Body.Close()
	return string(body)
}
