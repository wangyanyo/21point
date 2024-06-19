package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/wangyanyo/21point/common/enum"
)

func Cle() {
	cmd := &exec.Cmd{}
	if enum.SYS_TYPE == "windows" {
		cmd = exec.Command("cmd.exe", "/c", "cls")
	}
	if enum.SYS_TYPE == "linux" {
		cmd = exec.Command("sh", "-c", "clear")
		fmt.Println(cmd)
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func MD5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func PrintMessage(text string) {
	log.Println(text)
	fmt.Println(text)
	time.Sleep(1 * time.Second)
}

func CalcPoint(cards []string) int {
	res, cnt := 0, 0
	for _, s := range cards {
		if s == "A" {
			res++
			cnt++
		} else if s == "10" || s == "J" || s == "Q" || s == "K" {
			res += 10
		} else {
			res += int(s[0] - '0')
		}
	}
	if cnt > 0 && res+10 <= 21 {
		res += 10
	}
	return res
}
