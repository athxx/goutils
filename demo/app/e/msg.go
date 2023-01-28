package e

import (
	"strings"

	"github.com/gin-gonic/gin"

	"svr/app/util/logx"
)

// GetMsg get error information based on Code
func GetMsg(code int) (int, string) {
	if msg, ok := MsgFlags[code]; ok {
		return code / 1000, msg
	}
	return 200, MsgFlags[ERR_BAD_REQUEST]
}

// OutSuc 成功输出, fields 是额外字段, 与code, msg同级
func OutSuc(c *gin.Context, data interface{}, fields ...map[string]interface{}) {
	res := gin.H{"code": 0, "msg": "", "data": data}
	if len(fields) > 0 {
		for k, v := range fields[0] {
			res[k] = v
		}
	}
	c.AbortWithStatusJSON(200, res)
}

func OutJsonBytes(c *gin.Context, data []byte) {
	c.Abort()
	c.Header("Content-Type", "application/json")
	c.Writer.Write([]byte(`{"code":0, "msg":"","data":` + string(data) + `}`))
}

func OutJsonString(c *gin.Context, data string) {
	c.Abort()
	c.Header("Content-Type", "application/json")
	if data == `` {
		data = `[]`
	}
	c.Writer.Write([]byte(`{"code": 0,"msg":"ok","data":` + data + `}`))
}

func OutString(c *gin.Context, data ...string) {
	str := strings.Builder{}
	for k, v := range data {
		if k > 0 {
			str.WriteString(" ")
		}
		str.WriteString(v)
	}
	c.Writer.WriteString(str.String())
}

func OutBytes(c *gin.Context, data []byte) {
	c.Abort()
	c.Writer.Write(data)
}

// OutErr 错误输出
func OutErr(c *gin.Context, code int, err ...interface{}) {
	_, msg := GetMsg(code)
	if len(err) > 0 && err[0] != nil {
		e := err[0]
		switch v := e.(type) {
		case E:
			msg = v.Error()
			code = v.Code
			logx.Error(v.msg + ": " + v.st) // stacking info
		case error:
			msg = v.Error()
			logx.Error(msg)
			break
		case string:
			msg = v
		case int:
			if _, ok := MsgFlags[v]; ok {
				msg = MsgFlags[v]
			}
		}
	}
	c.AbortWithStatusJSON(200, gin.H{"code": code, "msg": msg, "data": struct{}{}})
}

// OutError 会按标准restful输出错误header信息
func OutError(c *gin.Context, code int, err ...interface{}) {
	//statusCode, msg := GetMsg(code)
	_, msg := GetMsg(code)
	if len(err) > 0 && err[0] != nil {
		e := err[0]
		switch v := e.(type) {
		case E:
			//statusCode = v.Code / 1000
			msg = v.Error()
			logx.Error(v.msg + ": " + v.st) // stacking info
		case error:
			msg = v.Error()
			logx.Error(msg)
			break
		case string:
			msg = v
		case int:
			if _, ok := MsgFlags[v]; ok {
				msg = MsgFlags[v]
			}
		}
	}
	//c.AbortWithStatusJSON(statusCode, gin.H{"code": code, "msg": msg, "data": struct{}{}})
	c.AbortWithStatusJSON(200, gin.H{"code": code, "msg": msg, "data": struct{}{}})
}

func OutBool(c *gin.Context, res bool) error {
	var err error
	if res {
		_, err = c.Writer.WriteString(`{"code": 0, "msg": "ok", "data":[]}`)
	} else {
		_, err = c.Writer.WriteString(`{"code": 400, "msg": "failed", "data":[]}`)
	}
	return err
}

// OutRedirect 重定向
func OutRedirect(c *gin.Context, code int, loc string) {
	if code < 301 || code > 308 || code == 304 {
		code = 303
	}
	c.Redirect(code, loc)
	c.Abort()
}
