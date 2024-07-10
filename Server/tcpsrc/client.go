package tcpsrc

import "net"

type ClientUser struct {
	Token      string
	Connection net.Conn
	NowRoom    int
	IsStart    bool
}
