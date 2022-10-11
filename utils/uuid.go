package utils

import (
	"strconv"

	"github.com/google/uuid"
)

func UUID() string {
	id := uuid.NewString()
	id = id[0:8] + id[9:13] + id[14:18] + id[19:23] + id[24:]

	var buf []byte
	for i := 0; i < 4; i++ {
		tmp, _ := strconv.ParseUint(id[i*8:(i+1)*8], 16, 32)
		buf = append(buf, Base62Encode(tmp)...)
	}
	l := len(buf)
	if l < 24 {
		l = 24 - l
		for l > 0 {
			buf = append([]byte{'0'}, buf...)
			l--
		}
	}
	return string(buf)
}
