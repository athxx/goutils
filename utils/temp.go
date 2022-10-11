package utils

import (
	"encoding/json"
	"os"
)

// TempSet 系统临时文件夹存取
func TempSet(name string, data interface{}) error {
	p := os.TempDir() + "/" + name
	str, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return os.WriteFile(p, str, os.ModePerm)
}

// TempGet os temp dir get
func TempGet(name string) (interface{}, error) {
	p := os.TempDir() + "/" + name
	s, err := os.ReadFile(p)
	if err != nil {
		return nil, err
	}
	var data interface{}
	if err = json.Unmarshal(s, &data); err != nil {
		return nil, err
	}
	return data, nil
}
