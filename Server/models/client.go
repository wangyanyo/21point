package models

import (
	"net"
	"time"
)

type ClientUser struct {
	Token      string
	Connection net.Conn
	NowRoom    int
	IsStart    bool
	CloseChan  chan struct{}
	LastTime   time.Time
	HeartTimer time.Ticker
}
