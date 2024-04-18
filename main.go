package main

import (
	"encoding/json"
	"fmt"
	"os"
	"server/mysql"
	"server/network"
	"server/redis"
)

func main() {
	netConfig, err := parseNetConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	dbConfig, err := parseDbConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	redisConfig, err := parseRedisConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = redis.InitRedis(redisConfig)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = mysql.InitMysql(dbConfig)
	if err != nil {
		fmt.Println(err)
		return
	}

	startLifeGameServer(netConfig)
}

func parseNetConfig() (network.NetConfig, error) {
	var netConfig network.NetConfig
	file, err := os.Open("./config/net_config.json")
	if err != nil {
		return netConfig, err
	}

	defer file.Close()

	jsonParser := json.NewDecoder(file)
	jsonParser.Decode(&netConfig)

	return netConfig, err
}

func parseDbConfig() (mysql.DbConfig, error) {
	var dbConfig mysql.DbConfig
	file, err := os.Open("./config/db_config.json")
	if err != nil {
		return dbConfig, nil
	}

	defer file.Close()

	jsonParser := json.NewDecoder(file)
	jsonParser.Decode(&dbConfig)

	return dbConfig, nil
}

func parseRedisConfig() (redis.RedisConfig, error) {
	var redisConfig redis.RedisConfig
	file, err := os.Open("./config/redis_config.json")
	if err != nil {
		return redisConfig, err
	}

	defer file.Close()

	jsonParser := json.NewDecoder(file)
	jsonParser.Decode(&redisConfig)

	return redisConfig, nil
}
