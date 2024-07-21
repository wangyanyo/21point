package service

import (
	"context"
	"log"

	"github.com/wangyanyo/21point/Server/common"
	"github.com/wangyanyo/21point/Server/models"
	"github.com/wangyanyo/21point/common/entity"
)

func (ser *Server) GetSocreHandler(ctx context.Context, req *entity.TransfeData) *entity.TransfeData {
	if len(req.Token) == 0 {
		return &entity.TransfeData{
			Code: common.TokenIsEmptyError,
			Msg:  common.TokenIsEmptyMsg,
		}
	}
	resp := ser.getScore(ctx, req)
	return resp
}

func (ser *Server) getScore(ctx context.Context, req *entity.TransfeData) (finalResp *entity.TransfeData) {
	var res entity.TransfeData

	defer func() {
		if allErr := recover(); allErr != nil {
			finalResp.Code = common.SystemPanicErr
			finalResp.Msg = common.SystemPanicMsg
			log.Println(ctx, allErr)
		}
	}()

	score, err := ser.CommonCache.ZScore(ctx, models.RankListCacheKey, req.Token).Result()
	if err != nil {
		res.Code = common.CallRedisError
		res.Msg = err.Error()
		return &res
	}
	res.Code = common.StatusSuccess
	res.Data = int(score)
	return &res
}
