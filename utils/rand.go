package utils

import (
	crand "crypto/rand"
	"math/big"
	"math/rand"
)

const StdLen = 16

// RandString more fast
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

// RandStr more random
func RandStr(l int, c ...string) string {
	var (
		chars = []byte(`0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz`)
		num   *big.Int
		str   []byte
	)
	if len(c) > 0 {
		chars = []byte(c[0])
	}
	chrLen := int64(len(chars))
	for len(str) < l {
		num, _ = crand.Int(crand.Reader, big.NewInt(chrLen))
		str = append(str, chars[num.Int64()])
	}
	return string(str)
}

func RandNum(i int) string {
	return RandStr(i, `0123456789`)
}

func RandHexHash(l int) string {
	return string(RandStr(l, `0123456789abcdef`))
}
