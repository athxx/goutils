package md

import (
	"time"

	"svr/app/util"
)

const (
	USER_PREFIX = "xkid_"
)

func UserRandomName() string {
	return USER_PREFIX + util.Base62Encode(uint64(time.Now().UnixNano()))
}
