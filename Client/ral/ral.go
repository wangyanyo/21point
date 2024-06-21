package ral

import (
	"fmt"
	"log"

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
			log.Println("断线重连")
			fmt.Println("断线重连")
			retry--
			Connect()
			continue
		}
		break
	}
	if retry == 0 {
		utils.PrintMessage("505")
		return myerror.New("505")
	}
	return nil
}

func RAL(c *models.TcpClient, cmd enum.Command, token string, data interface{}, errMsg ...string) (*entity.TransfeData, error) {
	if err := SendRequest(c, cmd, token, data); err != nil {
		return nil, err
	}
	res := <-c.CmdChan
	if err := myerror.CheckPacket(res, cmd, errMsg); err != nil {
		return nil, err
	}
	return res, nil
}
