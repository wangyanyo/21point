package tcpsrc

import (
	"context"
	"log"
	"net"

	"github.com/panjf2000/ants"
	"github.com/wangyanyo/21point/Server/models"
	"github.com/wangyanyo/21point/common/entity"
	"github.com/wangyanyo/21point/common/myerror"
)

type TcpServer struct {
	Listener *net.TCPListener
}

func Run(ctx context.Context) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", "0.0.0.0:8200")
	myerror.PanicErr(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	myerror.PanicErr(err)

	defer listener.Close()

	p, _ := ants.NewPoolWithFunc(1000, func(t interface{}) {
		client := t.(*models.ClientUser)
		resv := make([]byte, 1024)
		n, err := client.Connection.Read(resv)
		if err != nil {
			log.Println("[接收数据失败]", client.Connection.RemoteAddr().String(), client.Connection)
			return
		}

		if n > 0 && n < 1025 {
			Router(ctx, entity.TransfeDataDecoder(resv), client)
		} else {
			log.Println("[数据错误]", client.Connection.RemoteAddr().String(), client.Connection)
			return
		}
	})

	defer p.Release()

	tcpServer := TcpServer{
		Listener: listener,
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
		}

		p.Invoke(clientUser)
	}
}
