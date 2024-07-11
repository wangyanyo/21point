package tcpsrc

import (
	"io"
	"log"
	"net"

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

		clientUser := &ClientUser{
			Connection: conn,
		}

		go func(client *ClientUser) {
			resv := make([]byte, 1024)
			ch := make(chan int)
			go func(client *ClientUser, resv []byte, ch chan int) {
				n, err := client.Connection.Read(resv)
				if err != nil {
					if err == io.EOF {
						return
					}
					ch <- 0
					return
				}
				ch <- n
			}(client, resv, ch)

			select {
			case n := <-ch:
				if n > 0 && n < 1025 {
					Router(entity.TransfeDataDecoder(resv))
				}
			}

		}(clientUser)
	}
}
