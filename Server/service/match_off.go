package service

import (
	"context"
	"log"

	"github.com/wangyanyo/21point/common/entity"
)

func (ser *Server) MatchOffHandler(ctx context.Context, req *entity.TransfeData) {
	if len(req.Token) == 0 {
		return
	}
	ser.matchOff(ctx, req)
}

func (ser *Server) matchOff(ctx context.Context, req *entity.TransfeData) {
	defer func() {
		if allErr := recover(); allErr != nil {
			log.Println(ctx, allErr)
		}
	}()

	ser.MatchOffFlag[req.Token] = true
}
