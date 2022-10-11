package cc

import (
	"svr/app/util/rdx"
)

const (
	CC_DEFAULT_TTL = 10
)

func clean(prefix string) {
	res, _ := rdx.Keys(prefix + `*`).Result()
	for _, v := range res {
		rdx.Del(v)
	}
}
