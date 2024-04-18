package redis

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type RedisConfig struct {
	Addr             string `json:"addr"`
	Password         string `json:"password"`
	DB               int    `json:"db"`
	SessionUniqueKey string `json:"session_unique_key"`
	UserPrefix       string `json:"user_prefix"`
}

var _redisClient *redis.Client
var _ctx context.Context
var _redisConfig RedisConfig

func InitRedis(redisConfig RedisConfig) error {
	_redisConfig = redisConfig
	_redisClient = redis.NewClient(&redis.Options{
		Addr:     redisConfig.Addr,
		Password: redisConfig.Password,
		DB:       redisConfig.DB,
	})

	_ctx = context.Background()

	_ = _redisClient
	return nil
}

func GetUniqueSessionId() uint64 {
	script := redis.NewScript(`
		local key = KEYS[1]
		local value = redis.call("EXISTS", key)
		if value == 0 then
			redis.call("SET", key, 1)
		else 
			redis.call("INCR", key)
		end

		local ret = redis.call("GET", key)
		return ret
	`)

	id, err := script.Run(_ctx, _redisClient, []string{_redisConfig.SessionUniqueKey}).Uint64()
	if err != nil {
		fmt.Println(err)
		return 0
	}

	return id
}

func UserPrefix() string {
	return _redisConfig.UserPrefix
}

func Set(key string, value interface{}) error {
	p, err := json.Marshal(value)
	if err != nil {
		return err
	}

	redisCmd := _redisClient.Set(_ctx, key, p, 0)
	return redisCmd.Err()
}

func Get(key string, dest interface{}) error {
	stringCmd := _redisClient.Get(_ctx, key)
	if stringCmd.Err() != nil {
		return stringCmd.Err()
	}

	p := []byte(stringCmd.Val())
	return json.Unmarshal(p, dest)
}

func Del(key string) error {
	intCmd := _redisClient.Del(_ctx, key)
	return intCmd.Err()
}
