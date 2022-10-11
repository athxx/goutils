package cache

import (
	"fmt"
	"time"
)

var c Cache

type Cache interface {
	Get(key string) interface{}                                   // get cached value by key.
	GetMulti(keys []string) []interface{}                         // GetMulti is a batch version of Get.
	Put(key string, val interface{}, timeout time.Duration) error // set cached value with key and expire time.
	Delete(key string) error                                      // delete cached value by key.
	Incr(key string) error                                        // increase cached int value by key, as a counter.
	Decr(key string) error                                        // decrease cached int value by key, as a counter.
	IsExist(key string) bool                                      // check if cached value exists or not.
	ClearAll() error                                              // clear all cache.
	StartAndGC(config string) error                               // start gc routine based on config string settings.
}

// Instance is a function create a new Cache Instance
type Instance func() Cache

var adapters = make(map[string]Instance)

// Register makes a cache adapter available by the adapter name.
// If Register is called twice with the same name or if driver is nil,
// it panics.
func Register(name string, adapter Instance) {
	if adapter == nil {
		panic("cache: Register adapter is nil")
	}
	if _, ok := adapters[name]; ok {
		panic("cache: Register called twice for adapter " + name)
	}
	adapters[name] = adapter
}

// NewCache Create a new cache driver by adapter name and config string.
// config need to be correct JSON as string: {"interval":360}.
// it will start gc automatically.
func NewCache(adapterName, config string) (adapter Cache, err error) {
	instanceFunc, ok := adapters[adapterName]
	if !ok {
		err = fmt.Errorf("cache: unknown adapter name %q (forgot to import?)", adapterName)
		return
	}
	adapter = instanceFunc()
	err = adapter.StartAndGC(config)
	if err != nil {
		adapter = nil
	}
	return
}

func InitCache(adapterName, config string) (err error) {
	instanceFunc, ok := adapters[adapterName]
	if !ok {
		return fmt.Errorf("cache: unknown adapter name %q (forgot to import?)", adapterName)
	}
	c = instanceFunc()
	err = c.StartAndGC(config)
	if err != nil {
		c = nil
	}
	return
}

func Get(key string) interface{} {
	return c.Get(key)
}

func GetMulti(keys []string) []interface{} {
	return c.GetMulti(keys)
}
func Put(key string, val interface{}, ttl time.Duration) error {
	return c.Put(key, val, ttl)
}
func Delete(key string) error {
	return c.Delete(key)
}
func Incr(key string) error {
	return c.Incr(key)
}
func Decr(key string) error {
	return c.Decr(key)
}
func IsExist(key string) bool {
	return c.IsExist(key)
}
func ClearAll() error {
	return c.ClearAll()
}
func StartAndGC(cfg string) error {
	return c.StartAndGC(cfg)
}
