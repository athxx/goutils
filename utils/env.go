package utils

import (
	"strconv"
	"strings"
	"syscall"
)

func EnvStr(key string, def ...string) string {
	v, has := syscall.Getenv(key)
	if !has {
		if len(def) > 0 {
			return def[0]
		}
		panic(`env variable '` + key + `' not exist`)
	}
	return v
}

func EnvStrs(key string, def ...string) []string {
	v, has := syscall.Getenv(key)
	if !has {
		if len(def) > 0 {
			return def
		}
		panic(`env variable '` + key + `' not exist`)
	}
	return strings.Split(v, ",")
}

func EnvBool(key string, def ...bool) bool {
	v, has := syscall.Getenv(key)
	if !has {
		if len(def) > 0 {
			return def[0]
		}
		panic(`env variable '` + key + `' not exist`)
	}
	switch v {
	case "1", "t", "T", "true", "TRUE", "True":
		return true
	}
	return false
}

func EnvInt(key string, def ...int) int {
	v, has := syscall.Getenv(key)
	if !has {
		if len(def) > 0 {
			return def[0]
		}
		panic(`env variable '` + key + `' not exist`)
	}
	i, _ := strconv.Atoi(v)
	return i
}

func EnvInt64(key string, def ...int64) int64 {
	v, has := syscall.Getenv(key)
	if !has {
		if len(def) > 0 {
			return def[0]
		}
		panic(`env variable '` + key + `' not exist`)
	}
	i, _ := strconv.ParseInt(v, 10, 64)
	return i
}
