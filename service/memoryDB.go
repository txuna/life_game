package service

import (
	"fmt"
	"server/redis"
)

type RedisUser struct {
	UniqueID     uint64 `json:"unique_id"`
	SessionID    int32  `json:"server_id"`
	UserID       []byte `json:"user_id"`
	UserIDLength int8   `json:"user_id_length"`
	IsAuth       bool   `json:"is_auth"`
}

/*
 */
func StoreUserInfo(sessionUniqueId uint64, sessionId int32, isAuth bool) error {
	redisUser := RedisUser{
		UniqueID:  sessionUniqueId,
		SessionID: sessionId,
		IsAuth:    isAuth,
	}

	copy(redisUser.UserID[:], []byte{})
	redisUser.UserIDLength = 0
	userPrefix := redis.UserPrefix()
	key := fmt.Sprintf("%s%d", userPrefix, sessionUniqueId)
	return redis.Set(key, redisUser)
}

func LoadUserInfo(networkUniqueID uint64, serverSessionId int32) {

}

func RemoveUserInfo(sessionUniqueId uint64) error {
	userPrefix := redis.UserPrefix()
	key := fmt.Sprintf("%s%d", userPrefix, sessionUniqueId)
	return redis.Del(key)
}
