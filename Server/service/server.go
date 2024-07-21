package service

import (
	"context"
	"sync"

	"github.com/redis/go-redis/v9"
	"github.com/wangyanyo/21point/Server/common"
	"github.com/wangyanyo/21point/Server/dao"
	"github.com/wangyanyo/21point/Server/models"
)

type Server struct {
	UserDao      *dao.UserDao
	CommonCache  *redis.Client
	RoomIDQueue  *models.Queue
	MatchSet     []*int
	RoomSet      map[int]*models.Room
	MatchRWMutex []sync.RWMutex
	MatchOffFlag map[string]bool
}

var defaultServer *Server

func GetServer() *Server {
	return defaultServer
}

func InitServer(ctx context.Context) {
	defaultServer = &Server{
		UserDao:      &dao.UserDao{},
		CommonCache:  models.Cache,
		RoomIDQueue:  new(models.Queue),
		MatchSet:     make([]*int, 501),
		RoomSet:      make(map[int]*models.Room),
		MatchRWMutex: make([]sync.RWMutex, 501),
		MatchOffFlag: make(map[string]bool),
	}

	defaultServer.RoomIDQueue.Init(500)

	for i := 1; i <= 500; i++ {
		err := defaultServer.RoomIDQueue.Push(i)
		if err != nil {
			panic(err)
		}
	}
}

func (ser *Server) IsHave(ctx context.Context, username string) (bool, error) {
	_, err := ser.CommonCache.ZRank(ctx, models.RankListCacheKey, username).Result()
	if err != nil {
		if err == redis.Nil {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (ser *Server) CheckRoom(roomID int, name string) (bool, string) {
	ser.MatchRWMutex[roomID].RLock()
	if ser.RoomSet[roomID] == nil || ser.RoomSet[roomID].Flag == false || ser.RoomSet[roomID].Exist(name) {
		_, ok := models.RoomExitMsgMap[name]
		if ok {
			msg := models.RoomExitMsgMap[name]
			delete(models.RoomExitMsgMap, name)
			return false, msg
		} else {
			return false, common.RoomIDIsWrongMsg
		}
	}
	return true, ""
}
