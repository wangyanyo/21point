package game

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/wangyanyo/21point/Client/models"
	"github.com/wangyanyo/21point/Client/ral"
	"github.com/wangyanyo/21point/Client/view"
	"github.com/wangyanyo/21point/common/entity"
	"github.com/wangyanyo/21point/common/enum"
	"github.com/wangyanyo/21point/common/utils"
)

func RankList(c *models.TcpClient) error {
	cnt := 1
	for {
		utils.Cle()
		log.Println("查看排行榜", cnt)
		fmt.Print(view.RankListViewHead)
		rankListInfo, err := ral.RAL(c, enum.RankListPactet, "", cnt)
		if err != nil {
			return err
		}

		for i, v := range rankListInfo.Data.([]entity.UserScoreInfo) {
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
			userCountInfo, err := ral.RAL(c, enum.UserCountPacket, "", 0)
			if err != nil {
				if err.Error() == "505" {
					return err
				} else {
					continue
				}
			}

			num := userCountInfo.Data.(int)

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
