package service

import (
	"context"
	"log"

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

	resp := ser.login(ctx, req)
	return resp
}

func (ser *Server) login(ctx context.Context, req *entity.TransfeData) (finalResp *entity.TransfeData) {
	var res entity.TransfeData
	userData := req.Data.(entity.User)

	defer func() {
		if allErr := recover(); allErr != nil {
			finalResp.Code = common.SystemPanicErr
			finalResp.Msg = common.SystemPanicMsg
			log.Println(ctx, allErr)
		}
	}()

	user, err := ser.UserDao.WhereName(userData.Name).Get()
	if err != nil {
		res.Code = common.CallDBError
		res.Msg = err.Error()
		return &res
	}
	if user.Id == 0 {
		res.Code = common.NotFoundUserError
		res.Msg = common.NotFoundUserMsg
		return &res
	}
	if user.Password != userData.Password {
		res.Code = common.PasswordWrongError
		res.Msg = common.PasswordWrongMsg
		return &res
	}

	res.Code = common.StatusSuccess
	res.Data = user.UserName
	return &res
}
