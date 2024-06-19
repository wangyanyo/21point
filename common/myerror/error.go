package myerror

import (
	"fmt"
	"log"
	"time"

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

func CheckPacket(data *entity.TransfeData, cmd enum.Command, errMsg ...[]string) error {
	log.Println(string(cmd), data)
	if data.Cmd != cmd {
		err := New(string(cmd) + "--PacketError")
		PrintError(err)
		return err
	} else if data.Code != 1 {
		err := New(string(cmd) + "--RequestError ")
		fmt.Println(errMsg)
		PrintError(err)
		return err
	}
	return nil
}
