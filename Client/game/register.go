package game

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

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
		models.Rconn <- true
		log.Println("断线重连", err)
		fmt.Println("连接已断开，正在尝试重连...")
		time.Sleep(1 * time.Second)
		return err
	}
	isRegister := <-c.CmdChan
	log.Println("isReister = ", isRegister)
	if isRegister.Cmd == enum.RegisterPacket {
		if isRegister.Data.(bool) {
			c.Token = isRegister.Token
			log.Println("注册成功", username, password)
			fmt.Println("注册成功！")
			time.Sleep(1 * time.Second)
			return nil
		} else {
			log.Println("用户名已存在", username, password)
			fmt.Println("用户名已存在")
			time.Sleep(1 * time.Second)
			return myerror.New("RepeatUsernameError")
		}
	}
	return nil
}
