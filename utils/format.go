package utils

import (
	"math"
)

// MaskString change string to abc***def
func MaskString(src string, hLen int) string {
	str := []rune(src)
	if hLen == 0 {
		hLen = 4
	}
	hideStr := ``
	for i := 0; i < hLen; i++ {
		hideStr += "*"
	}
	hideLen := len(str) / 2
	showLen := len(str) - hideLen
	if hideLen == 0 || showLen == 0 {
		return hideStr
	}
	subLen := showLen / 2
	if subLen == 0 {
		return string(str[:showLen]) + hideStr
	}
	s := string(str[:subLen])
	s += hideStr
	s += string(str[len(str)-subLen:])
	return s
}

func FloatFormat(f float64, i int) float64 {
	if i > 14 {
		return f
	}
	p := math.Pow10(i)
	return float64(int64((f+0.000000000000009)*p)) / p
}
