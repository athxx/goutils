package mw

import (
	"bytes"
	"io"

	"github.com/gin-gonic/gin"

	"svr/app/util"
	"svr/app/util/rdx"
)

// 限流器
func Limiter(c *gin.Context) {
	limit := 100 // 限流次数
	ttl := 1     // 限流过期时间
	ip := c.ClientIP()
	// 读取token或者ip
	token := c.GetHeader("Authorization")
	// 判断是否已经超出限额次数
	method := c.Request.Method
	host := c.Request.Host
	uri := c.Request.URL.String()

	buf := make([]byte, 2048)
	num, _ := c.Request.Body.Read(buf)
	body := buf[:num]
	// Write body back
	c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
	Md5 := util.Md5(ip + token + method + host + uri + string(body))
	if rdx.Exists(Md5) > 0 {
		c.AbortWithStatusJSON(429, gin.H{
			"code": 429,
			"msg":  "don't repeat the request",
			"data": struct{}{},
		})
		return
	}
	// 2s后没返回自动释放
	go rdx.SetEx(Md5, "0", ttl)
	key := "LIMITER_" + ip
	reqs, _ := rdx.Get(key).Int()
	if reqs >= limit {
		c.AbortWithStatusJSON(429, gin.H{
			"code": 429,
			"msg":  "too many requests",
			"data": struct{}{},
		})
		return
	}
	if reqs > 0 {
		go rdx.Incr(key)
	} else {
		go rdx.SetEx(key, 1, ttl)
	}
	c.Next()
	go rdx.Del(Md5)
}
