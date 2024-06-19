package ral

import (
	"fmt"
	"log"
	"time"

	"github.com/wangyanyo/21point/Client/models"
)

func Reconnect(c *models.TcpClient, err error, cnt int) {
	models.Rconn <- true
	c.StopChan <- struct{}{}
	log.Println("断线重连...", err)
	fmt.Printf("第%d次断线重连...", cnt)
	time.Sleep(1 * time.Second)
	fmt.Print("\033[2K\r")
}
