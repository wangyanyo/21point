package tcpsrc

import (
	"log"
	"net"
	"time"

	"github.com/wangyanyo/21point/Client/game"
	"github.com/wangyanyo/21point/Client/models"
	"github.com/wangyanyo/21point/common/entity"
)

func Run() {

	host := "192.168.245.170:8200"
	tcpAddr, err := net.ResolveTCPAddr("tcp", host)
	if err != nil {
		log.Printf("hawk server [%s] reslove error: [%s]", host, err.Error())
		time.Sleep(1 * time.Second)
		return
	}

	client := &models.TcpClient{
		TcpAddr: tcpAddr,
		CmdChan: make(chan *entity.TransfeData),
	}

	game.Home(client)
}
