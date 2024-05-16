package utils

import "log"

func PanicErr(err error) {
	if err != nil {
		panic("[终止]出现知名错误" + err.Error())
	}
}

func DebugErr(err error) {
	if err != nil {
		log.Panicln("[Debug] err" + err.Error())
	}
}
