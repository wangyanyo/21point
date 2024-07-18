package game

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/wangyanyo/21point/Client/models"
	"github.com/wangyanyo/21point/Client/ral"
	"github.com/wangyanyo/21point/Client/view"
	"github.com/wangyanyo/21point/common/entity"
	"github.com/wangyanyo/21point/common/enum"
	"github.com/wangyanyo/21point/common/utils"
)

func waitResult(c *models.TcpClient, point int) error {
	req := &entity.TransfeData{
		Cmd:    enum.GameResultPacket,
		Token:  c.Token,
		RoomID: c.RoomID,
		Data:   point,
	}
	resultInfo, err := ral.Ral(c, req)
	if err != nil {
		return err
	}
	if resultInfo.Data.(int) == 0 {
		fmt.Printf("\033[91m你输了, Score-10\033[0m")
	} else if resultInfo.Data.(int) == 1 {
		fmt.Printf("\033[93m平局\033[0m")
	} else {
		fmt.Printf("\033[92m你赢了, Score+10\033[0m")
	}
	time.Sleep(1500 * time.Millisecond)
	return nil
}

func exitRoom(c *models.TcpClient) {
	req := &entity.TransfeData{
		Cmd:    enum.ExitRoomPacket,
		Token:  c.Token,
		RoomID: c.RoomID,
	}
	ral.Ral(c, req)
}

func PlayGame(c *models.TcpClient) error {
	myCards := []string{}
	for {
		if len(myCards) == 0 {
			req := &entity.TransfeData{
				Cmd:    enum.InitCardPacket,
				Token:  c.Token,
				RoomID: c.RoomID,
			}
			initCardInfo, err := ral.Ral(c, req)
			if err != nil {
				exitRoom(c)
				return err
			}
			myCards = initCardInfo.Data.([]string)
		}

		utils.Cle()
		fmt.Print(view.PlayGameViewHead)

		fmt.Print("你的分数：")
		req := &entity.TransfeData{
			Cmd:    enum.GetScorePacket,
			Token:  c.Token,
			RoomID: c.RoomID,
		}
		scoreInfo, err := ral.Ral(c, req)
		if err != nil {
			exitRoom(c)
			return err
		}
		fmt.Println(scoreInfo.Data.(int))

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
			var opt string
			for {
				fmt.Print("请输入：")
				scanner := bufio.NewScanner(os.Stdin)
				scanner.Scan()
				opt = scanner.Text()
				if opt == "0" || opt == "1" || opt == "2" || opt == "3" {
					break
				}
			}
			if opt == "0" {
				req := &entity.TransfeData{
					Cmd:    enum.AskCardsPactet,
					Token:  c.Token,
					RoomID: c.RoomID,
				}
				cardInfo, err := ral.Ral(c, req)
				if err != nil {
					exitRoom(c)
					return err
				}
				card := cardInfo.Data.(string)
				myCards = append(myCards, card)

			} else if opt == "1" {
				stopFlag = true

			} else if opt == "2" {

			} else if opt == "3" {
				exitRoom(c)
				return nil
			}
		}

		if stopFlag {
			fmt.Println("停牌")
			point := utils.CalcPoint(myCards)
			err := waitResult(c, point)
			if err != nil {
				exitRoom(c)
				return err
			}
			myCards = []string{}
		}
	}
}
