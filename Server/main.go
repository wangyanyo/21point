package main

import (
	"context"

	"github.com/wangyanyo/21point/Server/dao"
	"github.com/wangyanyo/21point/Server/models"
	"github.com/wangyanyo/21point/Server/service"
	"github.com/wangyanyo/21point/Server/tcpsrc"
	"github.com/wangyanyo/21point/common/db"
)

func main() {
	ctx := context.Background()

	db.InitMysqlDB("WangYanYo", "20030302Wy!", "192.168.245.170", "3306", "wangyanyo_1")

	dao.InitTable()

	models.InitRedis()

	service.InitServer(ctx)

	tcpsrc.Run(ctx)
}
