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

func SendRequest(c *models.TcpClient, req *entity.TransfeData) error {
	retry := 2
	for retry > 0 {
		if _, err := c.Send(req.Byte()); err != nil {
			log.Println("第", 10-retry+1, "次请求失败")
			retry--
		}
	}
	if retry == 0 {
		utils.PrintMessage("网络断开连接")
		return errors.New("网络断开连接")
	}
	return nil
}

func Read(c *models.TcpClient) (*entity.TransfeData, error) {
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

func CheckPacket(data *entity.TransfeData, req *entity.TransfeData) error {
	log.Println(string(req.Cmd), data)
	if data.Cmd != req.Cmd {
		err := errors.New(string(req.Cmd) + "--PacketError")
		myerror.PrintError(err)
		return err
	} else if data.Code != 0 {
		err := errors.New(string(req.Cmd) + "--RequestError" + data.Message)
		myerror.PrintError(err)
		return err
	}
	return nil
}

func Ral(c *models.TcpClient, req *entity.TransfeData) (*entity.TransfeData, error) {
	if err := SendRequest(c, req); err != nil {
		myerror.PrintError(err)
		return nil, err
	}
	res, err := Read(c)
	if err != nil {
		myerror.PrintError(err)
		return nil, err
	}
	if err := CheckPacket(res, req); err != nil {
		return nil, err
	}
	return res, nil
}
