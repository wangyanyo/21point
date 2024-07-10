package tcpsrc

import (
	"io"
	"log"
	"net"

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
			for {
				n, err := client.Connection.Read(resv)
				log.Println(n, err)
				if err != nil {
					if err == io.EOF {
						//退出房间
						return
					}
				}
				if n > 0 && n < 1025 {
					Router(client, resv[:n])
				}
			}

		}(clientUser)
	}
}
