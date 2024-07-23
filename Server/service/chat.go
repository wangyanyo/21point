package service

import (
	"context"
	"log"

	"github.com/wangyanyo/21point/Server/common"
	"github.com/wangyanyo/21point/common/entity"
)

func (ser *Server) ChatHandler(ctx context.Context, req *entity.TransfeData) *entity.TransfeData {
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
	resp := ser.chat(ctx, req)
	return resp
}

func (ser *Server) chat(ctx context.Context, req *entity.TransfeData) (finalResp *entity.TransfeData) {
	var (
		res entity.TransfeData
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

	data := req.Data.(string)
	if len(msg) > 0 {
		ser.RoomSet[req.RoomID].AddMsg(req.Token, data)
	}
	res.Code = common.StatusSuccess
	return &res
}
