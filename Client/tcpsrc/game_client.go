package tcpsrc

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/wangyanyo/21point/Client/game"
	"github.com/wangyanyo/21point/Client/models"
	"github.com/wangyanyo/21point/Client/ral"
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

	var connection *net.TCPConn
	for {
		connection, err = net.DialTCP("tcp", nil, hawkServer)
		if err == nil {
			break
		}
		log.Printf("connect to hawk server error: [%s]", err.Error())
		log.Printf("connect to hawk server error: [%s]", err.Error())
		time.Sleep(1 * time.Second)
	}

	client := &models.TcpClient{
		Connection: connection,
		HawkServer: hawkServer,
		StopChan:   make(chan struct{}),
		CmdChan:    make(chan *entity.TransfeData),
	}

	ctx, cancel := context.WithCancel(context.Background())

	//启动接收
	go func(ctx context.Context, conn *models.TcpClient) {
		resv := make([]byte, 1024)
		ch := make(chan int)
		for {
			go func(c *models.TcpClient, ch chan int, resv []byte) {
				n, err := conn.Read(resv)
				if err != nil {
					ch <- 0
					return
				}
				ch <- n
			}(conn, ch, resv)

			select {
			case n := <-ch:
				if n > 0 && n < 1025 {
					conn.CmdChan <- entity.TransfeDataDecoder(resv)
				}

			case <-models.ReConnChan:
				conn.Close()
				models.ConnChan <- struct{}{}

			case <-ctx.Done():
				conn.Close()
				fmt.Println("接收协程关闭")
				return

			}

		}
	}(ctx, client)

	go func(ctx context.Context, conn *models.TcpClient) {
		i := 0
		heartBeatTick := time.NewTicker(10 * time.Second)
		for {
			select {
			case <-heartBeatTick.C:
				heartBeat := entity.NewTransfeData(enum.HeartPacket, "", i)
				if _, err := conn.Send(heartBeat); err != nil {
					ral.Connect()
				}
				i++

			case <-conn.StopChan:
				return

			case <-ctx.Done():
				fmt.Println("心跳协程关闭")
				return
			}
		}
	}(ctx, client)

	go game.Home(client)

	for {
		select {
		case <-models.ConnChan:
			models.Connecting = true
			log.Println("Connect")
			client.Connection, err = net.DialTCP("tcp", nil, hawkServer)
			if err != nil {
				log.Printf("connect to hawk server error: [%s] " + err.Error())
			}
			models.Connecting = false
			models.EndConnChan <- struct{}{}

		case <-models.ExitChan:
			cancel()
			time.Sleep(1 * time.Second)
			return

		}
	}
}
