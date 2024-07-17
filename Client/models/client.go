package models

import (
	"net"

	"github.com/wangyanyo/21point/common/entity"
)

type TcpClient struct {
	Connection *net.TCPConn
	TcpAddr    *net.TCPAddr
	CmdChan    chan *entity.TransfeData
	Token      string
	RoomID     int
}

func (c *TcpClient) Send(b []byte) (int, error) {
	return c.Connection.Write(b)
}

func (c *TcpClient) Read(b []byte) (int, error) {
	return c.Connection.Read(b)
}

func (c *TcpClient) Addr() string {
	return c.Connection.RemoteAddr().String()
}

func (c *TcpClient) Close() {
	c.Connection.Close()
}
