package service

import (
	"context"

	"github.com/wangyanyo/21point/Server/common"
	"github.com/wangyanyo/21point/Server/models"
	"github.com/wangyanyo/21point/common/entity"
)

func (ser *Server) ExitRoomHandler(ctx context.Context, req *entity.TransfeData) *entity.TransfeData {
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
	resp := ser.exitRoom(ctx, req)
	return resp
}

func (ser *Server) exitRoom(ctx context.Context, req *entity.TransfeData) (finalResp *entity.TransfeData) {
	var (
		res entity.TransfeData
	)

	flag := req.Data.(int)
	otherPlayer := ser.RoomSet[req.RoomID].GetOtherPlayer(req.Token)
	if flag == 1 {
		models.RoomExitMsgMap[otherPlayer] = common.OtherExitMsg
		res.Code = common.ExitRoomCode
		res.Msg = common.ExitMsg
	} else if flag == 2 {
		ser.CommonCache.ZIncrBy(ctx, models.RankListCacheKey, -10.0, req.Token)
		ser.CommonCache.ZIncrBy(ctx, models.RankListCacheKey, 10.0, otherPlayer)
		models.RoomExitMsgMap[otherPlayer] = common.OtherEscapeMsg
		res.Code = common.ExitRoomCode
		res.Msg = common.EscapeMsg
	}
	ser.DeleteRoom(req.RoomID)

	return &res
}
