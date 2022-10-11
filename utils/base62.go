package utils

const (
	// 按ASCII码排序
	base62Char = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	base62Len  = 62
)

var base62Arr = map[int32]uint64{'0': 0, '1': 1, '2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9, 'A': 10, 'B': 11, 'C': 12, 'D': 13, 'E': 14, 'F': 15, 'G': 16, 'H': 17, 'I': 18, 'J': 19, 'K': 20, 'L': 21, 'M': 22, 'N': 23, 'O': 24, 'P': 25, 'Q': 26, 'R': 27, 'S': 28, 'T': 29, 'U': 30, 'V': 31, 'W': 32, 'X': 33, 'Y': 34, 'Z': 35, 'a': 36, 'b': 37, 'c': 38, 'd': 39, 'e': 40, 'f': 41, 'g': 42, 'h': 43, 'i': 44, 'j': 45, 'k': 46, 'l': 47, 'm': 48, 'n': 49, 'o': 50, 'p': 51, 'q': 52, 'r': 53, 's': 54, 't': 55, 'u': 56, 'v': 57, 'w': 58, 'x': 59, 'y': 60, 'z': 61}

// Base62Encode encode int to str
func Base62Encode(num uint64) string {
	if num == 0 {
		return "0"
	}
	res := ``
	for num > 0 {
		round := num / base62Len
		remain := num % base62Len
		res = string(base62Char[remain]) + res
		num = round
	}
	return res
}

// Base62Decode decode str to int
func Base62Decode(str string) uint64 {
	l := len(str) - 1
	var res uint64
	for _, v := range str {
		if _, ok := base62Arr[v]; !ok {
			return 0
		}
		res += base62Arr[v] * func(n int) uint64 { // pow
			var pow uint64 = 1
			for n > 0 {
				pow = pow * base62Len
				n -= 1
			}
			return pow
		}(l)
		l -= 1
	}
	return res
}

// Base62CusEncode 自定义base62加密
func Base62CusEncode(num uint64, baseChar string) string {
	if num == 0 || len(baseChar) != base62Len {
		return "0"
	}
	res := ``
	for num > 0 {
		round := num / base62Len
		remain := num % base62Len
		res = string(baseChar[remain]) + res
		num = round
	}
	return res
}

// Base62CusDecode 自定义base62解密
func Base62CusDecode(str string, baseChar string) uint64 {
	if str == `` || len(baseChar) != 62 {
		return 0
	}
	baseArr := map[int32]uint64{}
	for k, v := range baseChar {
		baseArr[v] = uint64(k)
	}
	l := len(str) - 1
	var res uint64
	for _, v := range str {
		if _, ok := baseArr[v]; !ok {
			return 0
		}
		res += baseArr[v] * func(n int) uint64 { // pow
			var pow uint64 = 1
			for n > 0 {
				pow = pow * base62Len
				n -= 1
			}
			return pow
		}(l)
		l -= 1
	}
	return res
}
