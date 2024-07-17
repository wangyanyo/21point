package game

import (
	"bufio"
	"fmt"
	"os"

	"github.com/wangyanyo/21point/Client/models"
	"github.com/wangyanyo/21point/Client/ral"
	"github.com/wangyanyo/21point/Client/view"
	"github.com/wangyanyo/21point/common/entity"
	"github.com/wangyanyo/21point/common/enum"
	"github.com/wangyanyo/21point/common/utils"
)

func Game(c *models.TcpClient) error {
	for {
		utils.Cle()
		fmt.Print(view.GameView)
		fmt.Print("你的分数: ")
		req := &entity.TransfeData{
			Cmd:    enum.GetScorePacket,
			Token:  c.Token,
			RoomID: c.RoomID,
		}
		scoreInfo, err := ral.Ral(c, req)
		if err != nil {
			continue
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
