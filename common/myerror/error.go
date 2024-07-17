package myerror

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/wangyanyo/21point/common/entity"
)

func PrintError(err error) {
	log.Println(err)
	fmt.Println(err)
	time.Sleep(1 * time.Second)
}

func CheckPacket(data *entity.TransfeData, req *entity.TransfeData) error {
	log.Println(string(req.Cmd), data)
	if data.Cmd != req.Cmd {
		err := errors.New(string(req.Cmd) + "--PacketError")
		PrintError(err)
		return err
	} else if data.Code != 1 {
		err := errors.New(string(req.Cmd) + "--RequestError")
		PrintError(err)
		return err
	}
	return nil
}

func PanicErr(err error) {
	if err != nil {
		panic("[终止] 出现致命错误: " + err.Error())
	}
}

func DebugErr(err error) {
	if err != nil {
		log.Println("[Debug] err: " + err.Error())
	}
}
