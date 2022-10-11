package utils

import (
	crand "crypto/rand"
	"math/big"
	"math/rand"
)

const StdLen = 16

// RandString
func RandString(length int, chars []byte) string {
	if length == 0 {
		return ""
	}
	clen := len(chars)
	if clen == 0 {
		chars = []byte(`0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz`)
		clen = len(chars)
	}

	maxRb := 255 - (256 % clen)
	b := make([]byte, length)
	r := make([]byte, length+(length/4)) // storage for random bytes.
	i := 0
	for {
		if _, err := rand.Read(r); err != nil {
			println(`error reading random bytes: ` + err.Error())
			return ""
		}
		for _, rb := range r {
			c := int(rb)
			if c > maxRb {
				// Skip this number to avoid modulo bias.
				continue
			}
			b[i] = chars[c%clen]
			i++
			if i == length {
				return string(b)
			}
		}
	}
}

// RandStr 随机程度高
func RandStr(l int, c ...string) string {
	var (
		chars = `0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz`
		str   string
		num   *big.Int
	)
	if len(c) > 0 {
		chars = c[0]
	}
	chrLen := int64(len(chars))
	for len(str) < l {
		num, _ = crand.Int(crand.Reader, big.NewInt(chrLen))
		str += string(chars[num.Int64()])
	}
	return str
}

func RandNum(i int) string {
	return RandStr(i, `0123456789`)
}

func RandHexHash(l int) string {
	return RandString(l, []byte(`0123456789abcdef`))
}
