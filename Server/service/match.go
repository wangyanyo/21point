package service

import (
	"context"
	"log"
	"time"

	"github.com/wangyanyo/21point/Server/common"
	"github.com/wangyanyo/21point/Server/models"
	"github.com/wangyanyo/21point/common/entity"
	"github.com/wangyanyo/21point/common/utils"
)

func (ser *Server) MatchHandler(ctx context.Context, req *entity.TransfeData) *entity.TransfeData {
	if len(req.Token) == 0 {
		return &entity.TransfeData{
			Code: common.TokenIsEmptyError,
			Msg:  common.TokenIsEmptyMsg,
		}
	}
	resp := ser.match(ctx, req)
	return resp
}

func (ser *Server) match(ctx context.Context, req *entity.TransfeData) (finalResp *entity.TransfeData) {
	var (
		res entity.TransfeData
		err error
	)

	defer func() {
		if allErr := recover(); allErr != nil {
			finalResp.Code = common.SystemPanicErr
			finalResp.Msg = common.SystemPanicMsg
			log.Println(ctx, allErr)
		}
	}()

	val, err := ser.CommonCache.ZScore(ctx, models.RankListCacheKey, req.Token).Result()
	if err != nil {
		res.Code = common.CallRedisError
		res.Msg = err.Error()
		return &res
	}
	myScore := int(val)

	roomID := 0
	for i := 1; i <= 500; i++ {
		ser.MatchRWMutex[i].RLock()
		if ser.MatchSet[i] == nil {
			ser.MatchRWMutex[i].RUnlock()
			continue
		}
		score := *ser.MatchSet[i]
		ser.MatchRWMutex[i].RUnlock()

		if utils.Abs(int(myScore)-score) <= 100 {
			ser.MatchRWMutex[i].Lock()
			if ser.MatchSet[i] == nil {
				ser.MatchRWMutex[i].Unlock()
				continue
			}
			ser.MatchSet[i] = nil
			ser.MatchRWMutex[i].Unlock()

			roomID = i
			ser.RoomSet[roomID].Flag = true
			ser.RoomSet[roomID].Player2 = req.Token
			break
		}
	}

	if roomID == 0 {
		roomID, err := ser.RoomIDQueue.Pop()
		if err != nil {
			res.Code = common.AskQueueError
			res.Msg = err.Error()
			return &res
		}

		ser.RoomSet[roomID] = &models.Room{
			Flag:    false,
			Player1: req.Token,
		}

		ser.MatchRWMutex[roomID].Lock()
		ser.MatchSet[roomID] = &myScore
		ser.MatchRWMutex[roomID].Unlock()

		for {
			ser.MatchRWMutex[roomID].RLock()
			if ser.RoomSet[roomID].Flag {
				ser.MatchRWMutex[roomID].RUnlock()
				break
			}
			ser.MatchOffFlag[req.Token] = false
			ser.MatchRWMutex[roomID].RUnlock()

			if ser.MatchOffFlag[req.Token] {
				ser.MatchRWMutex[roomID].Lock()
				if ser.MatchSet[roomID] != nil {
					ser.MatchSet[roomID] = nil
				}
				ser.MatchRWMutex[roomID].Unlock()
				ser.RoomSet[roomID] = nil
				ser.RoomIDQueue.Push(roomID)
				ser.MatchOffFlag[req.Token] = false

				res.Code = common.MatchOffCode
				res.Msg = common.MatchOffMsg
				return &res
			}

			time.Sleep(500 * time.Millisecond)
		}
	}

	res.Code = common.StatusSuccess
	res.Data = roomID
	return &res
}
