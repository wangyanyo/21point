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
		opt := utils.GetOpt("请输入: ", 1)
		if opt == "0" {
			scanner := bufio.NewScanner(os.Stdin)
			fmt.Print("请输入用户名: ")
			scanner.Scan()
			username := scanner.Scan()
			req := &entity.TransfeData{
				Cmd:    enum.SearchPacket,
				Token:  c.Token,
				RoomID: c.RoomID,
				Data:   username,
			}
			userInfo, err := ral.Ral(c, req)
			if err != nil {
				return err
			}

			targetUser = userInfo.Data.(entity.UserScoreInfo)
			flag = true
		}
		if opt == "1" {
			utils.PrintMessage("退出搜索界面成功！")
			return nil
		}
	}
}
