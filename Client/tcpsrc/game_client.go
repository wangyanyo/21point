package tcpsrc

import (
	"io"
	"log"
	"net"
	"time"

	"github.com/wangyanyo/21point/Client/game"
	"github.com/wangyanyo/21point/Client/models"
	"github.com/wangyanyo/21point/common/entity"
	"github.com/wangyanyo/21point/common/enum"
)

func Run() {

	host := "192.168.245.170:8200"
	hawkServer, err := net.ResolveTCPAddr("tcp", host)
	if err != nil {
		log.Printf("hawk server [%s] reslove error: [%s]", host, err.Error())
		time.Sleep(1 * time.Second)
		return
	}

	connection, err := net.DialTCP("tcp", nil, hawkServer)
	if err != nil {
		log.Printf("connect to hawk server error: [%s]", err.Error())
		time.Sleep(1 * time.Second)
		///////////////////////////////////////////////////////////////////////////////
	}

	client := &models.TcpClient{
		Connection: connection,
		HawkServer: hawkServer,
		StopChan:   make(chan struct{}),
		CmdChan:    make(chan *entity.TransfeData),
	}

	//启动接收
	go func(conn *models.TcpClient) {
		resv := make([]byte, 1024)
		for {
			n, err := conn.Read(resv)
			if err != nil {
				if err == io.EOF {
					log.Printf(conn.Addr(), " 断开了连接")
					conn.Close()
					return
				}
			}
			if n > 0 && n < 1025 {
				conn.CmdChan <- entity.TransfeDataDecoder(resv)
			}
		}
	}(client)

	go func(conn *models.TcpClient) {
		i := 0
		heartBeatTick := time.Tick(10 * time.Second)
		for {
			select {
			case <-heartBeatTick:
				heartBeat := entity.NewTransfeData(enum.HeartPacket, "", i)
				if _, err := conn.Send(heartBeat); err != nil {
					models.Rconn <- true
					return
				}
				i++

			case <-conn.StopChan:
				return
			}
		}
	}(client)

	go game.Home(client)

}
