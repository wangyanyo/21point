package models

import (
	"net"
	"sync"

	"github.com/wangyanyo/21point/common/entity"
)

type TcpClient struct {
	Connection *net.TCPConn
	TcpAddr    *net.TCPAddr
	Token      string
	RoomID     int
	ChatMsg    []*entity.ChatData
	Count      int
	PrintMutex sync.Mutex //光标只有一个，要有顺序的使用光标打印
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
