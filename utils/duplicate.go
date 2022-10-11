package utils

func DeduplicateString(ary *[]string) {
	set := make(map[string]struct{}, len(*ary))
	j := 0
	for _, v := range *ary {
		if _, ok := set[v]; !ok {
			set[v] = struct{}{}
			(*ary)[j] = v
			j++
		}
	}
	*ary = (*ary)[:j]
}

func DeduplicateInt(ary *[]int) {
	set := make(map[int]struct{}, len(*ary))
	j := 0
	for _, v := range *ary {
		if _, ok := set[v]; !ok {
			set[v] = struct{}{}
			(*ary)[j] = v
			j++
		}
	}
	*ary = (*ary)[:j]
}

func DeduplicateInt64(ary *[]int64) {
	set := make(map[int64]struct{}, len(*ary))
	j := 0
	for _, v := range *ary {
		if _, ok := set[v]; !ok {
			set[v] = struct{}{}
			(*ary)[j] = v
			j++
		}
	}
	*ary = (*ary)[:j]
}

func DeduplicateUint64(ary *[]uint64) {
	set := make(map[uint64]struct{}, len(*ary))
	j := 0
	for _, v := range *ary {
		if _, ok := set[v]; !ok {
			set[v] = struct{}{}
			(*ary)[j] = v
			j++
		}
	}
	*ary = (*ary)[:j]
}
