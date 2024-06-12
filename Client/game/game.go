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

func Game(c *models.TcpClient) error {
	for {
		utils.Cle()
		fmt.Print(view.GameView)
		fmt.Print("你的分数: ")
		if _, err := c.Send(entity.NewTransfeData(enum.GetScorePacket, c.Token, 0)); err != nil {
			myerror.Reconnect(err)
			return err
		}

		myScore := <-c.CmdChan
		log.Println("请求分数", myScore)
		if err := myerror.CheckPacket(myScore, enum.GetScorePacket); err != nil {
			continue
		}
		fmt.Println(myScore.Data.(int))

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
