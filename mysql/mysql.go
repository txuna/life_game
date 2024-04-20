package mysql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
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

var _mysqlClient *sql.DB
var _mysqlConfig DbConfig

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
