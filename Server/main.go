package main

import (
	"github.com/wangyanyo/21point/Server/dao"
	"github.com/wangyanyo/21point/Server/models"
	"github.com/wangyanyo/21point/Server/service"
	"github.com/wangyanyo/21point/Server/tcpsrc"
	"github.com/wangyanyo/21point/common/db"
)

func init() {
	db.InitMysqlDB("WangYanYo", "20030302Wy!", "192.168.245.170", "3306", "wangyanyo_1")

	dao.InitTable()

	models.InitRedis()

	service.InitServer()
}

func main() {
	tcpsrc.Run()
}
