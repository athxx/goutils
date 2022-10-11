package rdx

import (
	"context"
	"errors"
)

var (
	rdb           *redis.Client
	clientCluster *redis.ClusterClient
)

func InitRdb(addr, psw string, db int) error {
	if addr == `` {
		return errors.New("redis connect string cannot be empty")
	}

	rdb = redis.NewClient(&redis.Options{
		Addr:         addr,
		Password:     psw,
		DB:           db,
		DialTimeout:  redisDialTTL,
		ReadTimeout:  redisReadTTL,
		WriteTimeout: redisWriteTTL,
		IdleTimeout:  redisIdleTTL,
		PoolTimeout:  redisPoolTTL,
		PoolSize:     redisPoolSize,
	})
	err := rdb.Ping(context.Background()).Err()
	return err
}

func NewRdb(addr, psw string) *redis.Client {
	return redis.NewClient(&redis.Options{Addr: "addr", Password: psw, DB: 0})
}
