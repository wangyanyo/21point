package service

import (
	"context"
	"log"
	"time"

	"github.com/wangyanyo/21point/Server/common"
	"github.com/wangyanyo/21point/Server/models"
	"github.com/wangyanyo/21point/common/entity"
)

func (ser *Server) GameResultHandler(ctx context.Context, req *entity.TransfeData) *entity.TransfeData {
	if len(req.Token) == 0 {
		return &entity.TransfeData{
			Code: common.TokenIsEmptyError,
			Msg:  common.TokenIsEmptyMsg,
		}
	}
	if req.RoomID == 0 {
		return &entity.TransfeData{
			Code: common.RoomIDIsEmptyError,
			Msg:  common.RoomIDIsEmptyMsg,
		}
	}
	resp := ser.gameResult(ctx, req)
	return resp
}

func (ser *Server) gameResult(ctx context.Context, req *entity.TransfeData) (finalResp *entity.TransfeData) {
	var (
		res   entity.TransfeData
		point = req.Data.(int)
	)

	defer func() {
		if allErr := recover(); allErr != nil {
			finalResp.Code = common.SystemPanicErr
			finalResp.Msg = common.SystemPanicMsg
			log.Println(ctx, allErr)
		}
	}()

	jud, msg := ser.CheckRoom(req.RoomID, req.Token)
	if !jud {
		res.Code = common.RoomWrongError
		res.Msg = msg
		return &res
	}

	ser.RoomSet[req.RoomID].SetPoint(req.Token, point)
	for {
		if ser.RoomSet[req.RoomID].JudgeOtherTimeOut(req.Token) {
			otherPlayer := ser.RoomSet[req.RoomID].GetOtherPlayer(req.Token)
			ser.CommonCache.ZIncrBy(ctx, models.RankListCacheKey, 10.0, req.Token)
			ser.CommonCache.ZIncrBy(ctx, models.RankListCacheKey, -10.0, otherPlayer)

			ser.ExitRoom(req.RoomID)
		}

		if ser.RoomSet[req.RoomID].GetOtherPoint(req.Token) > 0 {
			otherPoint := ser.RoomSet[req.RoomID].GetOtherPoint(req.Token)
			flag := utils.CheckGameResult(point, otherPoint)
			if flag == 1 {
				ser.CommonCache.ZIncrBy(ctx, models.RankListCacheKey, 10.0, req.Token)
			} else if flag == -1 {
				ser.CommonCache.ZIncrBy(ctx, models.RankListCacheKey, -10.0, req.Token)
			}
			res.Code = common.StatusSuccess
			res.Data = otherPoint
			return &res
		}
		time.Sleep(500 * time.Millisecond)
	}
}
