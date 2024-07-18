package service

import "github.com/wangyanyo/21point/Server/dao"

type Server struct {
	UserDao *dao.UserDao
}

var defaultServer *Server

func GetServer() *Server {
	return defaultServer
}

func InitServer() {
	defaultServer = &Server{
		UserDao: &dao.UserDao{},
	}
}
