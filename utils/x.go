package utils

// SplitZero 拆分一个0阻隔, split zero from string
func SplitZero(s string) (string, string) {
	i := false
	div := '0'
	idx := -1
	for k, v := range s {
		if v == div && !i {
			i = true
			continue
		}
		if v != div && i {
			idx = k
			break
		}
	}
	if idx == -1 {
		return s, ""
	}
	return s[:idx-1], s[idx:]
}
