package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func SqlIn(raw interface{}) string {
	switch v := raw.(type) {
	case []string:
		buf := []byte{'('}
		replacer := strings.NewReplacer(`'`, `\'`, `\`, `\\`, `"`, `\"`)
		for i := len(v) - 1; i >= 0; i-- {
			buf = append(buf, '\'')
			buf = append(buf, replacer.Replace(strings.TrimSpace(v[i]))...)
			buf = append(buf, '\'')
			if i > 0 {
				buf = append(buf, ',')
			}
		}
		buf = append(buf, ')')
		return string(buf)
	case []int64, []int32, []int16, []int8, []int, []uint64, []uint32, []uint16, []uint8, []uint, []float32, []float64:
		return strings.NewReplacer(` `, `,`, `[`, `(`, `]`, `)`).Replace(fmt.Sprint(raw))
	}
	return `()`
}

// SqlInSlice is same with SqlIn, and faster than SqlIn
func SqlInSlice(raw interface{}) string {
	buf := []byte{'('}
	switch v := raw.(type) {
	case []string:
		replacer := strings.NewReplacer(`'`, `\'`, `\`, `\\`, `"`, `\"`)
		for i := len(v) - 1; i >= 0; i-- {
			buf = append(buf, '\'')
			buf = append(buf, replacer.Replace(strings.TrimSpace(v[i]))...)
			buf = append(buf, '\'')
			if i > 0 {
				buf = append(buf, ',')
			}
		}
	case []int64:
		l := len(v) - 1
		for k, val := range v {
			buf = append(buf, strconv.FormatInt(val, 10)...)
			if k < l {
				buf = append(buf, ',')
			}
		}
	case []int32:
		l := len(v) - 1
		for k, val := range v {
			buf = append(buf, strconv.FormatInt(int64(val), 10)...)
			if k < l {
				buf = append(buf, ',')
			}
		}
	case []int16:
		l := len(v) - 1
		for k, val := range v {
			buf = append(buf, strconv.FormatInt(int64(val), 10)...)
			if k < l {
				buf = append(buf, ',')
			}
		}
	case []int8:
		l := len(v) - 1
		for k, val := range v {
			buf = append(buf, strconv.FormatInt(int64(val), 10)...)
			if k < l {
				buf = append(buf, ',')
			}
		}
	case []int:
		l := len(v) - 1
		for k, val := range v {
			buf = append(buf, strconv.FormatInt(int64(val), 10)...)
			if k < l {
				buf = append(buf, ',')
			}
		}
	case []uint64:
		l := len(v) - 1
		for k, val := range v {
			buf = append(buf, strconv.FormatUint(val, 10)...)
			if k < l {
				buf = append(buf, ',')
			}
		}
	case []uint32:
		l := len(v) - 1
		for k, val := range v {
			buf = append(buf, strconv.FormatUint(uint64(val), 10)...)
			if k < l {
				buf = append(buf, ',')
			}
		}
	case []uint16:
		l := len(v) - 1
		for k, val := range v {
			buf = append(buf, strconv.FormatUint(uint64(val), 10)...)
			if k < l {
				buf = append(buf, ',')
			}
		}
	case []uint8:
		l := len(v) - 1
		for k, val := range v {
			buf = append(buf, strconv.FormatUint(uint64(val), 10)...)
			if k < l {
				buf = append(buf, ',')
			}
		}
	case []uint:
		l := len(v) - 1
		for k, val := range v {
			buf = append(buf, strconv.FormatUint(uint64(val), 10)...)
			if k < l {
				buf = append(buf, ',')
			}
		}
	case []float64, []float32:
		return strings.NewReplacer(` `, `,`, `[`, `(`, `]`, `)`).Replace(fmt.Sprint(raw))
	}
	if len(buf) == 1 {
		return `()`
	}
	buf = append(buf, ')')
	return string(buf)
}

func SqlErrDuplicate(err error) bool {
	if err == nil {
		return false
	}
	return `Error 1062: Duplicate entry` == err.Error()[:27]
}

func SqlErrNoRow(err error) bool {
	if err == nil {
		return false
	}
	return `sql: no rows in result set` == err.Error()
}

func SqlErrNil(err error) bool {
	if err == nil {
		return false
	}
	return `nil pointer passed to StructScan destination` == err.Error()
}
