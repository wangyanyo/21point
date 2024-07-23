package game

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/wangyanyo/21point/Client/models"
	"github.com/wangyanyo/21point/Client/ral"
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
	req := &entity.TransfeData{
		Cmd:    enum.LoginPacket,
		Token:  c.Token,
		RoomID: c.RoomID,
		Data:   userData,
	}
	loginInfo, err := ral.Ral(c, req)
	if err != nil {
		myerror.PrintError(err)
		return err
	}
	utils.PrintMessage("登录成功！")
	c.Token = loginInfo.Data.(string)
	return nil
}
