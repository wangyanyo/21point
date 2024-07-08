package db

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var MysqlDB *gorm.DB

func InitMysqlDB(user, pass, host, port, name string) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pass, host, port, name) + "?charset=utf8mb4&parseTime=true&loc=Local&timeout=5s"
	log.Println("连接数据库 = ", dsn)
	MysqlDB, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Println("连接数据库失败", err)
		panic(err)
	}
	MysqlDB.LogMode(true)
	MysqlDB.DB().SetMaxIdleConns(10)
	MysqlDB.DB().SetMaxOpenConns(20)
}

func GetMysqlDB() *gorm.DB {
	return MysqlDB
}
