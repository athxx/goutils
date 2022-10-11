package utils

import (
	"strconv"
)

// ContainsString is 字符串是否包含在字符串切片里
func ContainsString(array []string, val string) (index int) {
	index = -1
	for i := 0; i < len(array); i++ {
		if array[i] == val {
			index = i
			return
		}
	}
	return
}

func SliceStringToUint64(s []string) (i []uint64) {
	for _, v := range s {
		if ii, err := strconv.ParseInt(v, 10, 64); err == nil {
			i = append(i, uint64(ii))
		}
	}
	return
}
