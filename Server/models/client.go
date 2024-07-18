package models

import (
	"net"
)

type ClientUser struct {
	Connection net.Conn
}
