package mw

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/gin-gonic/gin"
)

// Breaker
// r.Use(BreakerWrapper)
// r.Use(gin.Recovery())
func Breaker(c *gin.Context) {
	name := c.Request.Method + "-" + c.Request.RequestURI
	hystrix.Do(name, func() error {
		c.Next()
		statusCode := c.Writer.Status()
		if statusCode >= http.StatusInternalServerError {
			return errors.New("status code " + strconv.Itoa(statusCode))
		}
		return nil
	}, func(e error) error {
		if e == hystrix.ErrCircuitOpen {
			c.String(http.StatusAccepted, "请稍后重试") //todo 修改报错方法
		}
		return e
	})
}
