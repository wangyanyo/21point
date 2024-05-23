package game

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/wangyanyo/21point/Client/models"
	"github.com/wangyanyo/21point/Client/view"
	"github.com/wangyanyo/21point/common/entity"
	"github.com/wangyanyo/21point/common/enum"
	"github.com/wangyanyo/21point/common/myerror"
	"github.com/wangyanyo/21point/common/utils"
)

func RankList(c *models.TcpClient) error {
	cnt := 1
	for {
		utils.Cle()
		log.Println("查看排行榜", cnt)
		fmt.Print(view.RankListViewHead)
		if _, err := c.Send(entity.NewTransfeData(enum.RankListPactet, "", cnt)); err != nil {
			myerror.Reconnect(err)
			return err
		}
		rankList := <-c.CmdChan
		if err := myerror.CheckPacket(rankList, enum.RankListPactet); err != nil {
			return err
		}

		for i, v := range rankList.Data.([]entity.UserScoreInfo) {
			fmt.Println(strconv.Itoa(cnt+i) + "\t\t" + v.Name + "\t\t" + strconv.Itoa(v.Score))
		}
		fmt.Print(view.RankListViewTail)
		fmt.Print("请输入: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		text := scanner.Text()
		if text == "0" {
			if cnt == 1 {
				utils.PrintMessage("这是第一页！")
				continue
			}
			cnt = max(1, cnt-10)
			continue
		}
		if text == "1" {
			if _, err := c.Send(entity.NewTransfeData(enum.UserCountPacket, "", 0)); err != nil {
				myerror.Reconnect(err)
				return err
			}
			userCount := <-c.CmdChan
			if err := myerror.CheckPacket(userCount, enum.UserCountPacket); err != nil {
				continue
			}
			num := userCount.Data.(int)

			if cnt+10 > num {
				utils.PrintMessage("这是最后一页！")
				continue
			}
			cnt = min(num, cnt+10)
			continue
		}
		if text == "2" {
			if err := Search(c); err != nil {
				continue
			}
			continue
		}
		if text == "3" {
			utils.PrintMessage("退出排行榜成功！")
			return nil
		}
	}
}
