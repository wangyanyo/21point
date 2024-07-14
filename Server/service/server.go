package service

import "github.com/wangyanyo/21point/Server/dao"

type Server struct {
	Dao *dao.UserDao
}

var server *Server

func GetServer() *Server {
	return server
}

func InitServer() {
	server = &Server{
		Dao: &dao.UserDao{},
	}
}
