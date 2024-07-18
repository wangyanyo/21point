package service

import (
	"context"
	"log"

	"github.com/wangyanyo/21point/Server/common"
	"github.com/wangyanyo/21point/Server/models"
	"github.com/wangyanyo/21point/common/entity"
)

func (ser *Server) RegisterHandler(ctx context.Context, req *entity.TransfeData) *entity.TransfeData {
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

func (ser *Server) register(ctx context.Context, req *entity.TransfeData) (finalResp *entity.TransfeData) {
	var res entity.TransfeData
	userData := req.Data.(entity.User)

	defer func() {
		if allErr := recover(); allErr != nil {
			finalResp.Code = common.ProcessErr
			log.Println(ctx, allErr)
		}
	}()

	flag, err := ser.UserDao.IsHave(userData.Name)
	if err != nil {
		res.Code = common.CallDBError
		res.Msg = err.Error()
		return &res
	}
	if flag {
		res.Code = common.UserExistedError
		res.Msg = common.UserExistedMsg
		return &res
	}

	user := models.User{
		UserName: userData.Name,
		Password: userData.Password,
	}
	if err := ser.UserDao.Create(user); err != nil {
		res.Code = common.CallDBError
		res.Msg = err.Error()
		return &res
	}

	res.Code = common.StatusSuccess
	res.Data = user.UserName
	return &res
}
