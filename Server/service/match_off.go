package service

import (
	"context"
	"log"

	"github.com/wangyanyo/21point/Server/common"
	"github.com/wangyanyo/21point/common/entity"
)

func (ser *Server) MatchOffHandler(ctx context.Context, req *entity.TransfeData) *entity.TransfeData {
	if len(req.Token) == 0 {
		return &entity.TransfeData{
			Code: common.TokenIsEmptyError,
			Msg:  common.TokenIsEmptyMsg,
		}
	}
	resp := ser.matchOff(ctx, req)
	return resp
}

func (ser *Server) matchOff(ctx context.Context, req *entity.TransfeData) (finalResp *entity.TransfeData) {
	var res entity.TransfeData

	defer func() {
		if allErr := recover(); allErr != nil {
			finalResp.Code = common.SystemPanicErr
			finalResp.Msg = common.SystemPanicMsg
			log.Println(ctx, allErr)
		}
	}()

	ser.MatchOffFlag[req.Token] = true
	res.Code = common.StatusSuccess
	return &res
}
