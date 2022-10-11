package e

import (
	"fmt"
	"path"
	"runtime"
)

type E struct {
	Code int
	msg  string
	st   string
}

var ErrDebug bool

func NewErrCode(code int) error {
	if msg, ok := MsgFlags[code]; ok {
		return E{code, msg, stack(3)}
	}
	return E{ERR_UNKNOWN, "unknown", stack(3)}
}

func NewErr(code int, msg string) error {
	return E{code, msg, stack(3)}
}

func (e E) Error() string {
	return e.msg
}

func stack(skip int) (str string) {
	if ErrDebug {
		stk := make([]uintptr, 32)
		l := runtime.Callers(skip, stk[:])
		for i := 0; i < l; i++ {
			f := runtime.FuncForPC(stk[i])
			name := f.Name()
			file, line := f.FileLine(stk[i])
			str += fmt.Sprintf("\n%-30s[%s:%d]", name, path.Base(file), line)
		}
	}
	return
}
