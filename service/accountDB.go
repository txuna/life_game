package service

import (
	"database/sql"
	"fmt"
	"server/errorcode"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
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

/*
데이터베이스 users 테이블 칼럼
*/
type Account struct {
	ID       int64
	UserID   string
	UserNAME string
	UserPW   string
}

var _mysqlClient *sql.DB
var _mysqlConfig DbConfig

/*
Mysql 초기화 부분
쿠버네티스 업로드시 환경변수로 진행
*/
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

/*
user_id를 기반으로 users 테이블에서 user를 가지고옴
*/
func LoadAccount(strUserId string) (Account, error) {
	stmt, err := _mysqlClient.Prepare("select user_id, user_name, user_pw from users where user_id = ?")
	if err != nil {
		return Account{}, err
	}

	defer stmt.Close()

	row := stmt.QueryRow(strUserId)
	var Id int64
	var userId, userName, userPw string
	err = row.Scan(&userId, &userName, &userPw)
	if err != nil && err == sql.ErrNoRows {
		return Account{}, err
	}

	return Account{
		ID:       Id,
		UserID:   userId,
		UserNAME: userName,
		UserPW:   userPw,
	}, nil
}

/*
사용자 로그인
*/
func LoginAccount(userID, userPW []byte) int16 {
	account, err := LoadAccount(string(userID))
	if err != nil {
		return errorcode.ERROR_CODE_MYSQL_ERROR
	}

	err = bcrypt.CompareHashAndPassword([]byte(account.UserPW), userPW)
	if err != nil {
		return errorcode.ERROR_CODE_MYSQL_ERROR
	}
	return errorcode.ERROR_CODE_NONE
}

/*
사용자 추가
*/
func JoinAccount(userID, userPW, userNAME []byte) int16 {
	_, err := LoadAccount(string(userID))
	if err == nil {
		return errorcode.ERROR_CODE_MYSQL_ERROR
	}

	hashPW, err := bcrypt.GenerateFromPassword(userPW, bcrypt.DefaultCost)
	if err != nil {
		return errorcode.ERROR_CODE_MYSQL_ERROR
	}

	stmt, err := _mysqlClient.Prepare("insert into users (user_id, user_name, user_pw) values(?, ?, ?)")
	if err != nil {
		return errorcode.ERROR_CODE_MYSQL_ERROR
	}

	defer stmt.Close()

	result, err := stmt.Exec(string(userID), string(userNAME), string(hashPW))
	if err != nil {
		return errorcode.ERROR_CODE_MYSQL_ERROR
	}

	/* auto increasement된 값 */
	_, _ = result.LastInsertId()
	return errorcode.ERROR_CODE_NONE
}
