package game

import (
	"fmt"

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

		opt := utils.GetOpt("请输入: ", 3)
		if opt == "0" {
			if err := Match(c); err != nil {
				continue
			}
			continue
		}
		if opt == "1" {
			if err := RankList(c); err != nil {
				continue
			}
			continue
		}
		if opt == "2" {
			c.Token = ""
			utils.PrintMessage("退出登陆成功！")
			return nil
		}
		if opt == "3" {
			utils.PrintMessage("退出游戏界面成功！")
			return nil
		}
	}

}
