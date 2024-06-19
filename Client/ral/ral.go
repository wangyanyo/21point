package ral

import (
	"github.com/wangyanyo/21point/Client/models"
	"github.com/wangyanyo/21point/common/entity"
	"github.com/wangyanyo/21point/common/enum"
	"github.com/wangyanyo/21point/common/myerror"
	"github.com/wangyanyo/21point/common/utils"
)

func SendRequest(c *models.TcpClient, cmd enum.Command, token string, data interface{}) error {
	retry := 10
	for retry > 0 {
		if _, err := c.Send(entity.NewTransfeData(cmd, token, data)); err != nil {
			Reconnect(c, err, 10-retry+1)
			continue
		}
	}
	if retry == 0 {
		return myerror.New("505")
	}
	return nil
}

func RAL(c *models.TcpClient, cmd enum.Command, token string, data interface{}, errMsg ...string) (*entity.TransfeData, error) {
	if err := SendRequest(c, cmd, token, data); err != nil {
		utils.PrintMessage("505")
		return nil, myerror.New("505")
	}
	res := <-c.CmdChan
	if err := myerror.CheckPacket(res, cmd, errMsg); err != nil {
		return nil, err
	}
	return res, nil
}
