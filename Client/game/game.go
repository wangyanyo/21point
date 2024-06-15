package game

import (
	"bufio"
	"fmt"
	"os"

	"github.com/wangyanyo/21point/Client/models"
	"github.com/wangyanyo/21point/Client/view"
	"github.com/wangyanyo/21point/common/enum"
	"github.com/wangyanyo/21point/common/utils"
)

func Game(c *models.TcpClient) error {
	for {
		utils.Cle()
		fmt.Print(view.GameView)
		fmt.Print("你的分数: ")
		scoreInfo, err := utils.RAL(c, enum.GetScorePacket, c.Token, "")
		if err != nil {
			if err.Error() == "505" {
				return err
			} else {
				continue
			}
		}

		fmt.Println(scoreInfo.Data.(int))

		fmt.Print("请输入: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		text := scanner.Text()
		if text == "0" {
			if err := Match(c); err != nil {
				continue
			}
			continue
		}
		if text == "1" {
			if err := RankList(c); err != nil {
				continue
			}
			continue
		}
		if text == "2" {
			utils.PrintMessage("退出游戏界面成功！")
			return nil
		}
	}

}
