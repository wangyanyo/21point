package tcpsrc

import (
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
		log.Println("[连接成功]", )
		go func(conn net.Conn) {
			
		}(conn)
	}
}
