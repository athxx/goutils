package utils

import (
	"encoding/base64"
	"fmt"
)

const (
	Base64Std = iota
	Base64Url
	Base64RawStd
	Base64RawUrl
)

func Base64StdEncode(str interface{}) string {
	return Base64Encode(str, Base64Std)
}

func Base64StdDecode(str interface{}) string {
	return Base64Decode(str, Base64Std)
}

func Base64UrlEncode(str interface{}) string {
	return Base64Encode(str, Base64Url)
}

func Base64UrlDecode(str interface{}) string {
	return Base64Decode(str, Base64Url)
}

func Base64RawStdEncode(str interface{}) string {
	return Base64Encode(str, Base64RawStd)
}

func Base64RawStdDecode(str interface{}) string {
	return Base64Decode(str, Base64RawStd)
}

func Base64RawUrlEncode(str interface{}) string {
	return Base64Encode(str, Base64RawUrl)
}

func Base64RawUrlDecode(str interface{}) string {
	return Base64Decode(str, Base64RawUrl)
}

func Base64Encode(str interface{}, encode int) string {
	newEncode := base64Encode(encode)
	if newEncode == nil {
		return ""
	}
	switch v := str.(type) {
	case string:
		return newEncode.EncodeToString([]byte(v))
	case []byte:
		return newEncode.EncodeToString(v)
	}
	return newEncode.EncodeToString([]byte(fmt.Sprint(str)))
}

func Base64Decode(str interface{}, encode int) string {
	var err error
	var b []byte
	newEncode := base64Encode(encode)
	if newEncode == nil {
		return ""
	}
	switch v := str.(type) {
	case string:
		b, err = newEncode.DecodeString(v)
	case []byte:
		b, err = newEncode.DecodeString(string(v))
	default:
		return ""
	}
	if err != nil {
		return ""
	}
	return string(b)
}

func base64Encode(encode int) *base64.Encoding {
	switch encode {
	case Base64Std:
		return base64.StdEncoding
	case Base64Url:
		return base64.URLEncoding
	case Base64RawStd:
		return base64.RawStdEncoding
	case Base64RawUrl:
		return base64.RawURLEncoding
	default:
		return nil
	}
}

const (
	base64Table = "1aA2bB3cC4dD5eE6fF7gG8hH9iI0jJ-kK_lLmMnNoOpPqQrRsStTuUvVwWxXyYzZ"
)

// CustomBase64 custom base64 characters
func CustomBase64(char ...string) *base64.Encoding {
	table := base64Table
	if len(char) > 1 && char[1] != `` && len(char[1]) == 64 {
		table = char[1]
	}
	return base64.NewEncoding(table)
}
