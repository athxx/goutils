package utils

import (
	"math/rand"
	"time"
)

func ShuffleString(s *string) {
	if len(*s) > 1 {
		b := []byte(*s)
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(b), func(x, y int) {
			b[x], b[y] = b[y], b[x]
		})
		*s = string(b)
	}
}

func ShuffleSliceBytes(b []byte) {
	if len(b) > 1 {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(b), func(x, y int) {
			b[x], b[y] = b[y], b[x]
		})
	}
}

func ShuffleSliceInt(i []int) {
	if len(i) > 1 {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(i), func(x, y int) {
			i[x], i[y] = i[y], i[x]
		})
	}
}

func ShuffleSliceInterface(i []interface{}) {
	if len(i) > 1 {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(i), func(x, y int) {
			i[x], i[y] = i[y], i[x]
		})
	}
}
