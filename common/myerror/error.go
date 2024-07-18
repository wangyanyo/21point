package myerror

import (
	"fmt"
	"log"
	"time"
)

func PrintError(err error) {
	log.Println(err)
	fmt.Println(err)
	time.Sleep(1 * time.Second)
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
