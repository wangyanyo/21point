package service

import (
	"context"
	"log"

	"github.com/wangyanyo/21point/Server/common"
	"github.com/wangyanyo/21point/Server/models"
	"github.com/wangyanyo/21point/common/entity"
)

func (ser *Server) RankListHandler(ctx context.Context, req *entity.TransfeData) *entity.TransfeData {
	if len(req.Token) == 0 {
		return &entity.TransfeData{
			Code: common.TokenIsEmptyError,
			Msg:  common.TokenIsEmptyMsg,
		}
	}
	if req.Data.(int) <= 0 {
		return &entity.TransfeData{
			Code: common.ReqDataIsEmptyError,
			Msg:  common.ReqDataIsEmptyMsg,
		}
	}
	resp := ser.rankList(ctx, req)
	return resp
}

func (ser *Server) rankList(ctx context.Context, req *entity.TransfeData) (finalResp *entity.TransfeData) {
	var (
		res      entity.TransfeData
		err      error
		ranklist []entity.UserScoreInfo
		cnt      = req.Data.(int64)
	)

	defer func() {
		if allErr := recover(); allErr != nil {
			finalResp.Code = common.ProcessErr
			log.Println(ctx, allErr)
		}
	}()

	size, err := ser.CommonCache.ZCard(ctx, models.RankListCacheKey).Result()
	if err != nil {
		res.Code = common.CallRedisError
		res.Msg = err.Error()
		return &res
	}

	vals, err := ser.CommonCache.ZRevRangeWithScores(ctx, models.RankListCacheKey, cnt, min(size, cnt+9)).Result()
	if err != nil {
		res.Code = common.CallRedisError
		res.Msg = err.Error()
		return &res
	}
	for i, val := range vals {
		ranklist = append(ranklist, entity.UserScoreInfo{
			Name:  val.Member.(string),
			Score: int(val.Score),
			Rank:  i + 1,
		})
	}

	res.Code = common.StatusSuccess
	res.Data = ranklist
	return &res
}
