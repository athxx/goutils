package utils

import (
	"encoding/json"
)

func Serialize(data interface{}) []byte {
	res, err := json.Marshal(data)
	if err != nil {
		return []byte{}
	}
	return res
}

func Unserialize(b []byte, dst interface{}) {
	if err := json.Unmarshal(b, dst); err != nil {
		dst = nil
	}
}

func SerializeStr(data interface{}, arg ...interface{}) string {
	return string(Serialize(data))
}
