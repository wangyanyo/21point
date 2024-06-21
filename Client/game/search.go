package game

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/wangyanyo/21point/Client/models"
	"github.com/wangyanyo/21point/Client/ral"
	"github.com/wangyanyo/21point/Client/view"
	"github.com/wangyanyo/21point/common/entity"
	"github.com/wangyanyo/21point/common/enum"
	"github.com/wangyanyo/21point/common/utils"
)

func Search(c *models.TcpClient) error {
	var targetUser entity.UserScoreInfo
	var flag bool = false
	for {
		utils.Cle()
		fmt.Print(view.SearchViewHead)
		if flag {
			fmt.Print(strconv.Itoa(targetUser.Rank) + "\t\t" + targetUser.Name + "\t\t" + strconv.Itoa(targetUser.Score) + "\n\n")
		} else {
			fmt.Print("\n\n")
		}
		fmt.Print(view.SearchViewTail)
		fmt.Print("请输入: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		text := scanner.Text()
		if text == "0" {
			fmt.Print("请输入用户名: ")
			scanner.Scan()
			username := scanner.Scan()
			userInfo, err := ral.RAL(c, enum.SearchPacket, "", username)
			if err != nil {
				return err
			}

			targetUser = userInfo.Data.(entity.UserScoreInfo)
			flag = true
		}
		if text == "1" {
			utils.PrintMessage("退出搜索界面成功！")
			return nil
		}
	}
}
