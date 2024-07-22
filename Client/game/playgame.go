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
	otherPoint := resultInfo.Data.(int)
	flag := utils.CheckGameResult(point, otherPoint)
	if flag == -1 {
		fmt.Printf("\033[91m你输了, 对方点数: %d, Score-10\033[0m", otherPoint)
	} else if flag == 0 {
		fmt.Printf("\033[93m平局\033[0m")
	} else {
		fmt.Printf("\033[92m你赢了, 对方点数: %d, Score+10\033[0m", otherPoint)
	}
	time.Sleep(1500 * time.Millisecond)
	return nil
}

func exitRoom(c *models.TcpClient, flag int) {
	req := &entity.TransfeData{
		Cmd:    enum.ExitRoomPacket,
		Token:  c.Token,
		RoomID: c.RoomID,
		Data:   flag,
	}
	ral.Ral(c, req)
	c.RoomID = 0
	c.Count = 0
	c.ChatMsg = make([]*entity.ChatData, 0)
}

func askCards(c *models.TcpClient) (string, error) {
	req := &entity.TransfeData{
		Cmd:    enum.AskCardsPactet,
		Token:  c.Token,
		RoomID: c.RoomID,
	}
	cardInfo, err := ral.Ral(c, req)
	if err != nil {
		return "", err
	}
	return cardInfo.Data.(string), nil
}

func askContinue() bool {
	fmt.Println("是否继续？")
	fmt.Println("操作: 0 继续   1 退出")
	opt := utils.GetOpt("请输入: ", 1)
	if opt == "0" {
		return true
	} else {
		return false
	}
}

func printChatView(c *models.TcpClient) {
	utils.SetPos(2, 54)
	fmt.Print(view.MessageView)

	for i := 1; i <= 30; i++ {
		utils.SetPos(i, 52)
		fmt.Print("|")
	}

	for i, v := range c.ChatMsg {
		utils.SetPos(i+3, 54)
		if v.Flag == 1 {
			fmt.Print("Me: ")
		} else if v.Flag == 2 {
			fmt.Print("Other: ")
		}
		fmt.Print(v.Msg)
	}

	utils.SetPos(1, 1)
}

func addMeesage(c *models.TcpClient, data *entity.ChatData) {
	c.ChatMsg = append(c.ChatMsg, data)
	utils.SetPos(len(c.ChatMsg)+2, 54)
	if data.Flag == 1 {
		fmt.Print("Me: ")
	} else if data.Flag == 2 {
		fmt.Print("Other: ")
	}
	fmt.Print(data.Msg)
}

func chat(c *models.TcpClient) error {
	fmt.Print("请输入消息: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	addMeesage(c, &entity.ChatData{
		Flag: 1,
		Msg:  text,
	})
	req := &entity.TransfeData{
		Cmd:    enum.ChatPacket,
		Token:  c.Token,
		RoomID: c.RoomID,
		Data:   text,
	}
	_, err := ral.Ral(c, req)
	return err
}

func pullMessage(c *models.TcpClient) {
	for {
		req := &entity.TransfeData{
			Cmd:    enum.ChatPacket,
			Token:  c.Token,
			RoomID: c.RoomID,
			Data:   c.Count,
		}
		resp, err := ral.Ral(c, req)
		if err != nil {

		}
		msgInfo := resp.Data.([]string)
		for _, v := range msgInfo {
			addMeesage(c, &entity.ChatData{
				Flag: 2,
				Msg:  v,
			})
		}
		c.Count += len(msgInfo)

		time.Sleep(1 * time.Second)
	}
}

func printHead() {
	fmt.Print(view.PlayGameViewHead)
	utils.SetPos(2, 54)
	fmt.Print(view.MessageView)
	for i := 1; i <= 30; i++ {
		utils.SetPos(i, 52)
		fmt.Print("|")
	}
}

func flushGame() {
	for i := 2; i <= 20; i++ {
		utils.SetPos(i, 1)
		fmt.Printf("%*c", 50, ' ')
	}
	utils.SetPos(2, 1)
}

func flushMessage() {
	for i := 2; i <= 20; i++ {
		utils.SetPos(i, 54)
		fmt.Printf("%*c", 100, ' ')
	}
	utils.SetPos(2, 54)
}

func printGame(c *models.TcpClient, myCards []string) error {
	flushGame()
	fmt.Print("你的分数：")
	req := &entity.TransfeData{
		Cmd:    enum.GetScorePacket,
		Token:  c.Token,
		RoomID: c.RoomID,
	}
	scoreInfo, err := ral.Ral(c, req)
	if err != nil {
		return err
	}
	fmt.Println(scoreInfo.Data.(int))

	fmt.Print("你的牌：")
	for _, card := range myCards {
		fmt.Printf("%s  ", card)
	}
	fmt.Printf("\n你的点数：%d\n\n", utils.CalcPoint(myCards))
	return nil
}

func PrintMessage(c *models.TcpClient) {
	for i, v := range c.ChatMsg {
		utils.SetPos(i+3, 54)
		if v.Flag == 1 {
			fmt.Print("Me: ")
		} else if v.Flag == 2 {
			fmt.Print("Other: ")
		}
		fmt.Print(v.Msg)
	}

	utils.SetPos(1, 1)
}

func PlayGame(c *models.TcpClient) error {
	go pullMessage(c)

	myCards := []string{}
	for {
		if len(myCards) == 0 {
			for i := 1; i <= 2; i++ {
				card, err := askCards(c)
				if err != nil {
					exitRoom(c, 1)
					return err
				}
				myCards = append(myCards, card)
			}
		}

		utils.Cle()
		printChatView(c)
		fmt.Print(view.PlayGameViewHead)

		fmt.Print("你的分数：")
		req := &entity.TransfeData{
			Cmd:    enum.GetScorePacket,
			Token:  c.Token,
			RoomID: c.RoomID,
		}
		scoreInfo, err := ral.Ral(c, req)
		if err != nil {
			exitRoom(c, 1)
			return err
		}
		fmt.Println(scoreInfo.Data.(int))

		fmt.Print("你的牌：")
		for _, card := range myCards {
			fmt.Printf("%s  ", card)
		}
		fmt.Printf("\n你的点数：%d\n\n", utils.CalcPoint(myCards))

		stopFlag := false
		if point := utils.CalcPoint(myCards); point >= 21 {
			stopFlag = true
		}

		if !stopFlag {
			fmt.Print(view.PlayGameViewTail)
			opt := utils.GetOpt("请输入: ", 3)
			if opt == "0" {
				card, err := askCards(c)
				if err != nil {
					exitRoom(c, 1)
					return err
				}
				myCards = append(myCards, card)
				continue

			} else if opt == "1" {
				stopFlag = true

			} else if opt == "2" {
				err := chat(c)
				if err != nil {
					exitRoom(c, 1)
					return err
				}
				if len(c.ChatMsg) > 20 {
					flushMessage(c)
				}

			} else if opt == "3" {
				exitRoom(c, 2)
				return nil
			}
		}

		if stopFlag {
			fmt.Println("停牌")
			point := utils.CalcPoint(myCards)
			err := waitResult(c, point)
			if err != nil {
				exitRoom(c, 1)
				return err
			}
			jud := askContinue()
			if !jud {
				exitRoom(c, 1)
				return nil
			}
			myCards = []string{}
		}
	}
}
