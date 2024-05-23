package game

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/wangyanyo/21point/Client/models"
	"github.com/wangyanyo/21point/Client/view"
	"github.com/wangyanyo/21point/common/entity"
	"github.com/wangyanyo/21point/common/enum"
	"github.com/wangyanyo/21point/common/myerror"
	"github.com/wangyanyo/21point/common/utils"
)

func Register(c *models.TcpClient) error {
	utils.Cle()
	fmt.Print(view.RegisterView)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("请输入用户名：")
	scanner.Scan()
	username := scanner.Text()
	fmt.Print("请输入密码：")
	scanner.Scan()
	password := scanner.Text()
	log.Println("注册", username, password)
	userData := entity.User{
		Name:     username,
		Password: password,
	}
	if _, err := c.Send(entity.NewTransfeData(enum.RegisterPacket, "", userData)); err != nil {
		myerror.Reconnect(err)
		return err
	}
	isRegister := <-c.CmdChan
	log.Println("isReister = ", isRegister)
	if err := myerror.CheckPacket(isRegister, enum.RegisterPacket, "用户名已存在，注册失败"); err != nil {
		return err
	}
	c.Token = isRegister.Token
	utils.PrintMessage("注册成功！")
	return nil
}
