package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

// DB needs to be implemented
type DB interface {
	Set(inst, category, name string, status CheckStatus, hint string, expirySecs int) error
	Get() (inst, category, name string, status CheckStatus, hint string)
}

type Category struct {
	CategoryName string      `json:"CategoryName,omitempty"`
	CheckName    string      `json:"CheckName,omitempty"`
	CheckValue   CheckStatus `json:"CheckValue,omitempty"`
	CheckHint    string      `json:"CheckHint,omitempty"`
}

// redis implemtention:

type redisBackend struct {
	client *redis.Client
}

func NewRedisBackend(addr, password string, dbnum int) *redisBackend {
	return &redisBackend{redis.NewClient(&redis.Options{
		Addr:     addr,     // "localhost:6379",
		Password: password, // no password set
		DB:       dbnum,    // use default DB
	})}
}

func (b *redisBackend) Set(inst, category, name string, status CheckStatus, hint string, expirySecs int) error {
	key := fmt.Sprintf("%s.%s.%s", inst, category, name)
	err := b.client.HSet(key, "status", status).Err()
	if err != nil {
		panic(err)
	}
	err = b.client.HSet(key, "hint", hint).Err()
	if err != nil {
		panic(err)
	}
	err = b.client.Expire(key, time.Duration(expirySecs)).Err()
	if err != nil {
		panic(err)
	}
	return nil
}

func (b *redisBackend) Get(inst, category, name string) (status CheckStatus, hint string) {
	key := fmt.Sprintf("%s.%s.%s", inst, category, name)
	statusStr, err := b.client.HGet(key, "status").Result()
	if err != nil {
		panic(err)
	}
	status, err = strconv.Atoi(statusStr)
	if err != nil {
		panic(err)
	}
	hint, err := b.client.HGet(key, "hint").Result()
	if err != nil {
		panic(err)
	}
	return status, hint
}

// InMemoryDB poc
type InMemoryDB struct {
	Instance map[string][]Category `json:"Instance,omitempty"`
}

func (b *InMemoryDB) Set(inst, category, name string, status CheckStatus, hint string, expirySecs int) error {
	//
	return nil
}

func NewInMemoryDB() *InMemoryDB {
	mem := InMemoryDB{}
	mem.Instance = make(map[string][]Category)
	return &mem
}
