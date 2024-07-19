package service

import (
	"github.com/redis/go-redis/v9"
	"github.com/wangyanyo/21point/Server/dao"
	"github.com/wangyanyo/21point/Server/models"
)

type Server struct {
	UserDao     *dao.UserDao
	CommonCache *redis.Client
}

var defaultServer *Server

func GetServer() *Server {
	return defaultServer
}

func InitServer() {
	defaultServer = &Server{
		UserDao:     &dao.UserDao{},
		CommonCache: models.Cache,
	}
}
