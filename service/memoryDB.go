package service

import (
	"encoding/json"
	"fmt"
	"server/protocol"

	"github.com/redis/go-redis/v9"
)

type RedisUser struct {
	UniqueID     uint64                                 `json:"unique_id"`
	ServerID     int32                                  `json:"server_id"`
	UserID       [protocol.MAX_USER_ID_BYTE_LENGTH]byte `json:"user_id"`
	UserIDLength int8                                   `json:"user_id_length"`
	IsAuth       bool                                   `json:"is_auth"`
}

/*
키가 없으면 1로 생성
가지고와선 1증가한다.
*/
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

/*
 */
func StoreUserInfo(redisUser RedisUser) {

}

func LoadUserInfo(networkUniqueID uint64, serverSessionId int32) {

}

func set(key string, value interface{}) error {
	p, err := json.Marshal(value)
	if err != nil {
		return err
	}

	redisCmd := _redisClient.Set(_ctx, key, p, 0)
	return redisCmd.Err()
}

func get(key string, dest interface{}) error {
	stringCmd := _redisClient.Get(_ctx, key)
	if stringCmd.Err() != nil {
		return stringCmd.Err()
	}

	p := []byte(stringCmd.Val())
	return json.Unmarshal(p, dest)
}
