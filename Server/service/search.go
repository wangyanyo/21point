package service

import (
	"context"
	"log"

	"github.com/wangyanyo/21point/Server/common"
	"github.com/wangyanyo/21point/Server/models"
	"github.com/wangyanyo/21point/common/entity"
)

func (ser *Server) SearchHandler(ctx context.Context, req *entity.TransfeData) *entity.TransfeData {
	if len(req.Token) == 0 {
		return &entity.TransfeData{
			Code: common.TokenIsEmptyError,
			Msg:  common.TokenIsEmptyMsg,
		}
	}
	if len(req.Data.(string)) == 0 {
		return &entity.TransfeData{
			Code: common.ReqDataIsEmptyError,
			Msg:  common.ReqDataIsEmptyMsg,
		}
	}

	resp := ser.search(ctx, req)
	return resp
}

func (ser *Server) search(ctx context.Context, req *entity.TransfeData) (finalResp *entity.TransfeData) {
	var (
		res      entity.TransfeData
		data     entity.UserScoreInfo
		username string
	)

	defer func() {
		if allErr := recover(); allErr != nil {
			finalResp.Code = common.ProcessErr
			log.Println(ctx, allErr)
		}
	}()

	username = req.Data.(string)
	rank, err := ser.CommonCache.ZRevRank(ctx, models.RankListCacheKey, username).Result()
	if err != nil {
		res.Code = common.CallRedisError
		res.Msg = err.Error()
		return &res
	}
	score, err := ser.CommonCache.ZScore(ctx, models.RankListCacheKey, username).Result()
	if err != nil {
		res.Code = common.CallRedisError
		res.Msg = err.Error()
		return &res
	}

	data.Name = username
	data.Rank = int(rank) + 1
	data.Score = int(score)

	res.Code = common.StatusSuccess
	res.Data = data
	return &res
}
