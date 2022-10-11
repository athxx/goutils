package mw

import (
	"github.com/gin-gonic/gin"

	"svr/app/e"
)

// AuthJWT is jwt middleware
func AuthJWT(c *gin.Context) {
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == `` {
		e.OutErr(c, e.ERR_UNAUTHORIZED, "token 不能为空")
		return
	}
}
