package hdl

import (
	"net/http"

	"svr/app/e"

	"github.com/gin-gonic/gin"
)

const AdmToken = "sid"

// SignIn 后台管理员登陆
func SignIn(c *gin.Context) {
	var args struct {
		Nickname string `json:"nickname"`
		Password string `json:"password"`
		OTP      string `json:"otp"` // one time password 基于TOTP算法
	}
	if err := c.ShouldBind(&args); err != nil || args.Nickname == `` || args.Password == `` || args.OTP == `` {
		e.OutError(c, e.ERR_INVALID_ARGS, err)
		return
	}

	// 判断用户登陆, 如果用户名已经错过了5次
	//cc.AdmSignInGet(args.Nickname)

	e.OutSuc(c, nil)
}

func SignOut(c *gin.Context) {
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     AdmToken,
		Value:    "",
		MaxAge:   -1,
		Path:     "/",
		Secure:   false,
		HttpOnly: true,
	})
	e.OutSuc(c, nil)
}
