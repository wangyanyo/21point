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

func Game(c *models.TcpClient) error {
	for {
		utils.Cle()
		fmt.Print(view.GameView)
		fmt.Print("你的分数: ")
		if _, err := c.Send(entity.NewTransfeData(enum.GetScorePaket, c.Token, 0)); err != nil {
			models.Rconn <- true
			log.Println("断线重连", err)
			fmt.Println("连接已断开，正在尝试重连...")
			time.Sleep(1 * time.Second)
			return err
		}
		myScore := <-c.CmdChan
		log.Println("请求分数", myScore)
		if myScore.Cmd == enum.GetScorePaket {
			fmt.Println(myScore.Data.(int))
		} else {
			return myerror.New("GetScoreError")
		}
		fmt.Print("请输入: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		text := scanner.Text()
		if text == "0" {

		}
		if text == "1" {
			if err := RankList(c); err != nil {
				continue
			}
			continue
		}
		if text == "2" {
			log.Println("退出游戏界面", c.Token)
			fmt.Println("退出成功！")
			time.Sleep(1 * time.Second)
			return nil
		}
	}

}
