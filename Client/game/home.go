package game

import (
	"fmt"

	"github.com/wangyanyo/21point/Client/models"
	"github.com/wangyanyo/21point/Client/view"
	"github.com/wangyanyo/21point/common/utils"
)

func Home(c *models.TcpClient) error {
	for {
		utils.Cle()
		fmt.Print(view.HomeView)
		opt := utils.GetOpt("请输入: ", 2)
		if opt == "0" {
			if err := Login(c); err != nil {
				continue
			}
			if err := Game(c); err != nil {
				continue
			}
			continue
		}
		if opt == "1" {
			if err := Register(c); err != nil {
				continue
			}
			if err := Game(c); err != nil {
				continue
			}
			continue
		}
		if opt == "2" {
			utils.PrintMessage("退出游戏成功！")
			return nil
		}
	}
}
