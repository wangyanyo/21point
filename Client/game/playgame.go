package game

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/wangyanyo/21point/Client/models"
	"github.com/wangyanyo/21point/Client/view"
	"github.com/wangyanyo/21point/common/enum"
	"github.com/wangyanyo/21point/common/utils"
)

func waitResult(c *models.TcpClient, point int) error {
	resultInfo, err := utils.RAL(c, enum.GameResultPacket, c.Token, point)
	if err != nil {
		return err
	}
	if resultInfo.Data.(int) == 0 {
		fmt.Printf("\033[91m你输了\033[0m")
	} else if resultInfo.Data.(int) == 1 {
		fmt.Printf("\033[93m平局\033[0m")
	} else {
		fmt.Printf("\033[92m平局\033[0m")
	}
	time.Sleep(1500 * time.Millisecond)
	return nil
}

func exitRoom(c *models.TcpClient) {
	utils.RAL(c, enum.ExitRoomPacket, c.Token, "")
}

func PlayGame(c *models.TcpClient) error {
	myCards := []string{}
	for {
		if len(myCards) == 0 {
			initCardInfo, err := utils.RAL(c, enum.InitCardPacket, c.Token, "")
			if err != nil {
				exitRoom(c)
				return err
			}
			myCards = initCardInfo.Data.([]string)
		}
		utils.Cle()
		fmt.Print(view.PlayGameViewHead)
		fmt.Print("你的牌：")
		for _, card := range myCards {
			fmt.Printf("%s  ", card)
		}
		fmt.Printf("\n总计：%d\n\n", utils.CalcPoint(myCards))

		stopFlag := false
		if point := utils.CalcPoint(myCards); point >= 21 {
			stopFlag = true
		}

		if !stopFlag {
			fmt.Print(view.PlayGameViewTail)
			fmt.Print("请输入：")
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			opt := scanner.Text()
			if opt == "0" {
				cardInfo, err := utils.RAL(c, enum.AskCardsPactet, c.Token, "")
				if err != nil {
					continue
				}
				card := cardInfo.Data.(string)
				myCards = append(myCards, card)

			} else if opt == "1" {
				stopFlag = true

			} else if opt == "2" {

			} else if opt == "3" {

			} else {
				continue
			}
		}

		if stopFlag {
			point := utils.CalcPoint(myCards)
			err := waitResult(c, point)
			if err != nil {
				continue
			}
			myCards = []string{}
		}
	}
}
