package ral

import (
	"errors"
	"log"
	"net"

	"github.com/wangyanyo/21point/Client/models"
	"github.com/wangyanyo/21point/common/entity"
	"github.com/wangyanyo/21point/common/myerror"
	"github.com/wangyanyo/21point/common/utils"
)

var resv = make([]byte, 1024)

func sendRequest(c *models.TcpClient, req *entity.TransfeData) error {
	retry := 2
	for retry > 0 {
		if _, err := c.Send(req.Byte()); err != nil {
			log.Println("第", 10-retry+1, "次请求失败")
			retry--
		}
		break
	}
	if retry == 0 {
		utils.PrintMessage("网络断开连接")
		return errors.New("网络断开连接")
	}
	return nil
}

func read(c *models.TcpClient) (*entity.TransfeData, error) {
	var err error
	c.Connection, err = net.DialTCP("tcp", nil, c.TcpAddr)
	if err != nil {
		return nil, err
	}
	n, err := c.Read(resv)
	if err != nil {
		return nil, err
	}

	if n > 0 && n < 1025 {
		return entity.TransfeDataDecoder(resv), nil
	}
	return nil, errors.New("接收数据出错")
}

func Ral(c *models.TcpClient, req *entity.TransfeData) (*entity.TransfeData, error) {
	if err := sendRequest(c, req); err != nil {
		return nil, err
	}
	res, err := read(c)
	if err != nil {
		return nil, err
	}
	if err := myerror.CheckPacket(res, req); err != nil {
		return nil, err
	}
	return res, nil
}
