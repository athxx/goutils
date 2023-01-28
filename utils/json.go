package utils

import (
	"bytes"
	"encoding/json"
	"io"
)

func JsonMarshal(interface{}) {

}

// 不转义科学计数法
func JsonDecode(data []byte, v interface{}) error {
	d := json.NewDecoder(bytes.NewReader(data))
	d.UseNumber()
	return d.Decode(v)
}

func JsonMinify(data []byte) []byte {
	var buff = new(bytes.Buffer)
	if err := json.Compact(buff, data); err != nil {
		return data
	}
	if b, err := io.ReadAll(buff); err == nil {
		return b
	}
	return data
}
