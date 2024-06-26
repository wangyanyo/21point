package game

import (
	"bufio"
	"fmt"
	"os"

	"github.com/wangyanyo/21point/Client/models"
	"github.com/wangyanyo/21point/Client/view"
	"github.com/wangyanyo/21point/common/utils"
)

func Home(c *models.TcpClient) error {
	for {
		utils.Cle()
		fmt.Print(view.HomeView)
		fmt.Print("请输入：")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		text := scanner.Text()
		if text == "0" {
			if err := Login(c); err != nil {
				continue
			}
			if err := Game(c); err != nil {
				continue
			}
			continue
		}
		if text == "1" {
			if err := Register(c); err != nil {
				continue
			}
			if err := Game(c); err != nil {
				continue
			}
			continue
		}
		if text == "2" {
			models.ExitChan <- struct{}{}
			utils.PrintMessage("退出游戏成功！")
			return nil
		}
	}
}
