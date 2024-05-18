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

func Login(c *models.TcpClient) error {
	utils.Cle()
	fmt.Print(view.LoginView)
	fmt.Print("请输入用户名：")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	username := scanner.Text()
	fmt.Print("请输入密码：")
	scanner.Scan()
	password := scanner.Text()
	log.Println("请求登录", username, password)
	userData := entity.User{
		Name:     username,
		Password: password,
	}
	if _, err := c.Send(entity.NewTransfeData(enum.LoginPacket, "", userData)); err != nil {
		models.Rconn <- true
		log.Println("断线重连...", err)
		fmt.Println("连接已断开，正在尝试重连...")
		time.Sleep(1 * time.Second)
		return err
	}
	isLogin := <-c.CmdChan
	log.Println("isLogin = ", isLogin)
	if isLogin.Cmd == enum.LoginPacket {
		if isLogin.Data.(int) == 0 {
			log.Println("登陆成功")
			fmt.Println("登陆成功！")
			time.Sleep(1 * time.Second)
			c.Token = isLogin.Token
			return nil
		} else if isLogin.Data.(int) == 1 {
			log.Println("用户名或密码错误")
			fmt.Println("用户名或密码错误！")
			time.Sleep(1 * time.Second)
			return &myerror.PasswordWrongError{}
		} else {
			log.Println("用户名不存在")
			fmt.Println("用户名不存在！")
			time.Sleep(1 * time.Second)
			return &myerror.NoUserNameError{}
		}
	}
	return nil
}
