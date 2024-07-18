package service

import (
	"context"

	"github.com/wangyanyo/21point/Server/common"
	"github.com/wangyanyo/21point/common/entity"
)

func (ser *Server) LoginHandler(ctx context.Context, req *entity.TransfeData) *entity.TransfeData {
	userData := req.Data.(entity.User)
	if len(userData.Name) == 0 {
		return &entity.TransfeData{
			Code: common.UserNameEmptyError,
			Msg:  common.UserNameEmptyMsg,
		}
	}
	if len(userData.Password) == 0 {
		return &entity.TransfeData{
			Code: common.PasswordEmptyError,
			Msg:  common.PasswordEmptyMsg,
		}
	}

	resp := ser.register(ctx, req)
	return resp
}
