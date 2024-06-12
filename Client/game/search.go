package game

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/wangyanyo/21point/Client/models"
	"github.com/wangyanyo/21point/Client/view"
	"github.com/wangyanyo/21point/common/entity"
	"github.com/wangyanyo/21point/common/enum"
	"github.com/wangyanyo/21point/common/myerror"
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
			if _, err := c.Send(entity.NewTransfeData(enum.SearchPacket, "", username)); err != nil {
				myerror.Reconnect(err)
				return err
			}
			user := <-c.CmdChan
			log.Println("查找用户", user)
			if user.Cmd == enum.SearchPacket {
				if user.Code == 1 {
					targetUser = user.Data.(entity.UserScoreInfo)
					flag = true
				} else {
					flag = false
					utils.PrintMessage("用户不存在！")
					time.Sleep(1 * time.Second)
					continue
				}
			} else {
				utils.PrintMessage("查找用户错误")
				time.Sleep(1 * time.Second)
				return myerror.New("SearchError")
			}
		}
		if text == "1" {
			utils.PrintMessage("退出搜索界面成功！")
			return nil
		}
	}
}
