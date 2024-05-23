package myerror

import (
	"fmt"
	"log"
	"time"

	"github.com/wangyanyo/21point/Client/models"
	"github.com/wangyanyo/21point/common/entity"
	"github.com/wangyanyo/21point/common/enum"
)

type MyError struct {
	s string
}

func New(text string) error {
	return &MyError{text}
}

func (e *MyError) Error() string {
	return e.s
}

func PrintError(err error) {
	log.Println(err)
	fmt.Println(err)
	time.Sleep(1 * time.Second)
}

func CheckPacket(data *entity.TransfeData, cmd enum.Command, str ...string) error {
	if data.Cmd != cmd {
		err := New(string(cmd) + "PacketError")
		PrintError(err)
		return err
	} else {
		err := New(string(cmd) + "Error ")
		fmt.Println(str)
		PrintError(err)
		return err
	}
}

func Reconnect(err error) {
	models.Rconn <- true
	log.Println("断线重连...", err)
	fmt.Println("断线重连...")
	time.Sleep(1 * time.Second)
}
