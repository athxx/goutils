package utils

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"unsafe"
)

// string与slice互转,零copy省内存

// zero copy to change slice to string
func Slice2String(b []byte) (s string) {
	pBytes := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	pString := (*reflect.StringHeader)(unsafe.Pointer(&s))
	pString.Data = pBytes.Data
	pString.Len = pBytes.Len
	return
}

// no copy to change string to slice
func StringToSlice(s string) (b []byte) {
	pBytes := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	pString := (*reflect.StringHeader)(unsafe.Pointer(&s))
	pBytes.Data = pString.Data
	pBytes.Len = pString.Len
	pBytes.Cap = pString.Len
	return
}

func StringToBytes(s string) []byte {
	return (*[0x7fff0000]byte)(unsafe.Pointer(
		(*reflect.StringHeader)(unsafe.Pointer(&s)).Data),
	)[:len(s):len(s)]
}

func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// 任意slice合并
func SliceJoin(sep string, elems ...interface{}) string {
	l := len(elems)
	if l == 0 {
		return ""
	}
	if l == 1 {
		s := fmt.Sprint(elems[0])
		sLen := len(s) - 1
		if s[0] == '[' && s[sLen] == ']' {
			return strings.Replace(s[1:sLen], " ", sep, -1)
		}
		return s
	}
	sep = strings.Replace(fmt.Sprint(elems), " ", sep, -1)
	return sep[1 : len(sep)-1]
}

// 批量字符串替换
func StringsBatchReplace(s string, old, new []string) (string, error) {
	if len(old) == 0 || len(old) != len(new) {
		return s, errors.New("invalid arguments")
	}
	for k, v := range old {
		s = strings.Replace(s, v, new[k], -1)
	}
	return s, nil
}
