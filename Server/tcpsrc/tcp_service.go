package tcpsrc

import "net"

type TcpServer struct {
	Listener   *net.Listener
	HawkServer *net.TCPAddr
}
