package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/wangyanyo/21point/Client/models"
	"github.com/wangyanyo/21point/common/entity"
	"github.com/wangyanyo/21point/common/enum"
	"github.com/wangyanyo/21point/common/myerror"
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

func RAL(c *models.TcpClient, cmd enum.Command, token string, data interface{}, errMsg ...string) (*entity.TransfeData, error) {
	if err := SendRequest(c, cmd, token, data); err != nil {
		PrintMessage("505")
		return nil, myerror.New("505")
	}
	res := <-c.CmdChan
	if err := myerror.CheckPacket(res, cmd, errMsg); err != nil {
		return nil, err
	}
	return res, nil
}

func SendRequest(c *models.TcpClient, cmd enum.Command, token string, data interface{}) error {
	retry := 10
	for retry > 0 {
		if _, err := c.Send(entity.NewTransfeData(cmd, token, data)); err != nil {
			myerror.Reconnect(err, 10-retry+1)
			continue
		}
	}
	if retry == 0 {
		return myerror.New("505")
	}
	return nil
}
