package tcpsrc

import (
	"context"
	"io"
	"log"
	"net"
	"time"

	"github.com/wangyanyo/21point/Server/models"
	"github.com/wangyanyo/21point/common/entity"
	"github.com/wangyanyo/21point/common/myerror"
)

type TcpServer struct {
	Listener *net.TCPListener
	TcpAddr  *net.TCPAddr
}

func Run() {
	tcpAddr, err := net.ResolveTCPAddr("tcp", "0.0.0.0:8200")
	myerror.PanicErr(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	myerror.PanicErr(err)

	defer listener.Close()

	tcpServer := TcpServer{
		Listener: listener,
		TcpAddr:  tcpAddr,
	}

	for {
		conn, err := tcpServer.Listener.Accept()
		if err != nil {
			log.Println("[连接失败]", err.Error())
			continue
		}
		log.Println("[连接成功]", conn.RemoteAddr().String(), conn)

		clientUser := &models.ClientUser{
			Connection: conn,
			CloseChan:  make(chan struct{}, 10),
			LastTime:   time.Now(),
			HeartTimer: *time.NewTicker(10 * time.Second),
		}

		go func(client *models.ClientUser) {
			ctx, cancel := context.WithCancel(context.Background())

			//接收信息协程
			go func(ctx context.Context, client *models.ClientUser) {
				resv := make([]byte, 1024)
				for {
					n, err := client.Connection.Read(resv)
					if err != nil {
						if err == io.EOF {
							return
						}
						continue
					}

					if n > 0 && n < 1025 {
						Router(ctx, entity.TransfeDataDecoder(resv), client)
					}
				}
			}(ctx, client)

			//心跳检测协程
			go func(ctx context.Context, client *models.ClientUser) {
				for {
					select {
					case <-client.HeartTimer.C:
						if time.Since(client.LastTime).Seconds() > float64(20*time.Second) {
							client.CloseChan <- struct{}{}
						}

					case <-ctx.Done():
						return
					}
				}

			}(ctx, client)

			//中央控制协程
			go func(client *models.ClientUser) {
				for {
					select {
					case <-client.CloseChan:
						//这里进行释放连接操作，存好新分数或新账户，并进行退出房间等一系列操作
						//要考虑到read读到一半的数据
						//要做好一切收尾工作，因为即使该连接退出，服务仍然在运行
						cancel()
					}
				}
			}(client)
		}(clientUser)
	}
}
