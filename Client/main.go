package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	logFile, err := os.OpenFile("client.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open log file failed", err)
		return
	}
	log.SetOutput(logFile)
}
