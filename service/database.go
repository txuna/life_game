package service

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/redis/go-redis/v9"
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
	Addr             string `json:"addr"`
	Password         string `json:"password"`
	DB               int    `json:"db"`
	SessionUniqueKey string `json:"session_unique_key"`
	UserPrefix       string `json:"user_prefix"`
}

var _redisClient *redis.Client
var _mysqlClient *sql.DB
var _ctx context.Context

var _redisConfig RedisConfig
var _mysqlConfig DbConfig

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

func InitMysql(dbConfig DbConfig) error {
	_mysqlConfig = dbConfig

	user := dbConfig.User
	password := dbConfig.Password
	protocol := dbConfig.Protocol
	host := dbConfig.Host
	port := dbConfig.Port
	database := dbConfig.Database

	var err error
	addr := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", user, password, protocol, host, port, database)
	_mysqlClient, err = sql.Open("mysql", addr)
	if err != nil {
		return err
	}

	return nil
}
