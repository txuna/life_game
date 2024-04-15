package main

import (
	"encoding/json"
	"fmt"
	"os"
	"server/network"
)

type DbConfig struct {
	Type     string `json:"type"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
	Protocol string `json:"protocol"`
}

type RedisConfig struct {
}

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

	_ = redisConfig
	_ = dbConfig

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

func parseDbConfig() (DbConfig, error) {
	var dbConfig DbConfig
	file, err := os.Open("./config/db_config.json")
	if err != nil {
		return dbConfig, nil
	}

	defer file.Close()

	jsonParser := json.NewDecoder(file)
	jsonParser.Decode(&dbConfig)

	return dbConfig, nil
}

func parseRedisConfig() (RedisConfig, error) {
	var redisConfig RedisConfig
	file, err := os.Open("./config/redis_config.json")
	if err != nil {
		return redisConfig, err
	}

	defer file.Close()

	jsonParser := json.NewDecoder(file)
	jsonParser.Decode(&redisConfig)

	return redisConfig, nil
}
