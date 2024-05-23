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
		myerror.Reconnect(err)
		return err
	}
	isLogin := <-c.CmdChan
	log.Println("isLogin = ", isLogin)
	if isLogin.Cmd == enum.LoginPacket {
		if isLogin.Data.(int) == 0 {
			utils.PrintMessage("登录成功！")
			c.Token = isLogin.Token
			return nil
		} else if isLogin.Data.(int) == 1 {
			utils.PrintMessage("用户名或密码错误！")
			return myerror.New("PassWordWrongError")
		} else {
			utils.PrintMessage("用户名不存在！")
			return myerror.New("NoUserNameError")
		}
	}
	return nil
}
