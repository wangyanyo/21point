package game

import (
	"fmt"
	"time"

	"github.com/eiannone/keyboard"
	"github.com/wangyanyo/21point/Client/models"
	"github.com/wangyanyo/21point/Client/ral"
	"github.com/wangyanyo/21point/Client/view"
	"github.com/wangyanyo/21point/common/entity"
	"github.com/wangyanyo/21point/common/enum"
	"github.com/wangyanyo/21point/common/myerror"
	"github.com/wangyanyo/21point/common/utils"
)

func printError(c *models.TcpClient, err error) {
	c.PrintMutex.Lock()
	utils.SetPos(15, 1)
	myerror.PrintError(err)
	c.PrintMutex.Unlock()
}

func waitResult(c *models.TcpClient, point int) error {
	req := &entity.TransfeData{
		Cmd:    enum.GameResultPacket,
		Token:  c.Token,
		RoomID: c.RoomID,
		Data:   point,
	}
	resultInfo, err := ral.Ral(c, req)
	if err != nil {
		printError(c, err)
		return err
	}
	otherPoint := resultInfo.Data.(int)
	flag := utils.CheckGameResult(point, otherPoint)

	c.PrintMutex.Lock()
	utils.SetPos(10, 1)
	if flag == -1 {
		fmt.Printf("\033[91m你输了, 对方点数: %d, Score-10\033[0m", otherPoint)
	} else if flag == 0 {
		fmt.Printf("\033[93m平局\033[0m")
	} else {
		fmt.Printf("\033[92m你赢了, 对方点数: %d, Score+10\033[0m", otherPoint)
	}
	c.PrintMutex.Unlock()

	time.Sleep(1500 * time.Millisecond)
	return nil
}

func exitRoom(c *models.TcpClient, flag int) {
	if c.ExitFlag {
		c.ExitFlag = true
		return
	}
	req := &entity.TransfeData{
		Cmd:    enum.ExitRoomPacket,
		Token:  c.Token,
		RoomID: c.RoomID,
		Data:   flag,
	}
	_, err := ral.Ral(c, req)
	if err != nil {
		printError(c, err)
	}
	c.RoomID = 0
	c.Count = 0
	c.ChatMsg = make([]*entity.ChatData, 0)
	c.ExitFlag = false
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

func askContinue(c *models.TcpClient, keysEvents <-chan keyboard.KeyEvent) bool {
	c.PrintMutex.Lock()
	utils.SetPos(11, 1)
	fmt.Println("是否继续？")
	fmt.Println("操作: 0 继续   1 退出")
	fmt.Print("请输入: ")
	c.PrintMutex.Unlock()
	opt := input(c, keysEvents, 13, 9)
	if opt == "0" {
		return true
	} else {
		return false
	}
}

func addMeesage(c *models.TcpClient, data *entity.ChatData) {
	c.AddMsgMutex.Lock()
	c.ChatMsg = append(c.ChatMsg, data)
	if len(c.ChatMsg) > 15 {
		c.ChatMsg = c.ChatMsg[len(c.ChatMsg)-10:]
	}
	c.AddMsgMutex.Unlock()
}

func chat(c *models.TcpClient, keysEvents <-chan keyboard.KeyEvent) error {
	c.PrintMutex.Lock()
	utils.SetPos(9, 1)
	fmt.Print("请输入消息: ")
	c.PrintMutex.Unlock()

	text := input(c, keysEvents, 9, 13)
	if len(text) == 0 {
		return nil
	}
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
	if err != nil {
		printError(c, err)
		return err
	}
	return nil
}

func pullMessage(c *models.TcpClient) {
	for {
		if c.ExitFlag {
			c.ExitFlag = false
			return
		}

		req := &entity.TransfeData{
			Cmd:    enum.ChatPacket,
			Token:  c.Token,
			RoomID: c.RoomID,
			Data:   c.Count,
		}
		resp, err := ral.Ral(c, req)
		if err != nil {
			printError(c, err)
			exitRoom(c, 1)
			return
		}
		msgInfo := resp.Data.([]string)
		for _, v := range msgInfo {
			addMeesage(c, &entity.ChatData{
				Flag: 2,
				Msg:  v,
			})
		}
		c.Count += len(msgInfo)
		if len(msgInfo) > 0 {
			printMessage(c)
		}

		time.Sleep(1 * time.Second)
	}
}

func printHead(c *models.TcpClient) {
	c.PrintMutex.Lock()
	fmt.Print(view.PlayGameViewHead)
	utils.SetPos(2, 54)
	fmt.Print(view.MessageView)
	for i := 1; i <= 30; i++ {
		utils.SetPos(i, 52)
		fmt.Print("|")
	}
	c.PrintMutex.Unlock()
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
	c.PrintMutex.Lock()
	flushGame()
	fmt.Print("你当前的分数是: ")
	req := &entity.TransfeData{
		Cmd:    enum.GetScorePacket,
		Token:  c.Token,
		RoomID: c.RoomID,
	}
	scoreInfo, err := ral.Ral(c, req)
	if err != nil {
		myerror.PrintError(err)
		c.PrintMutex.Unlock()
		return err
	}
	fmt.Println(scoreInfo.Data.(int))

	fmt.Print("你的牌: ")
	for _, card := range myCards {
		fmt.Printf("%s  ", card)
	}
	fmt.Printf("\n点数: %d\n\n", utils.CalcPoint(myCards))
	c.PrintMutex.Unlock()
	return nil
}

func printMessage(c *models.TcpClient) {
	c.PrintMutex.Lock()
	flushMessage()
	c.AddMsgMutex.Lock()
	for i, v := range c.ChatMsg {
		utils.SetPos(i+3, 54)
		if v.Flag == 1 {
			fmt.Print("Me: ")
		} else if v.Flag == 2 {
			fmt.Print("Other: ")
		}
		fmt.Print(v.Msg)
	}
	c.AddMsgMutex.Lock()
	c.PrintMutex.Unlock()
}

func input(c *models.TcpClient, keysEvents <-chan keyboard.KeyEvent, row int, cls int) string {
	var str []rune
	for {
		event := <-keysEvents
		c.PrintMutex.Lock()

		if event.Key == keyboard.KeySpace {
			utils.SetPos(row, cls)
			fmt.Print(" ")
			cls += utils.RealLength(" ")
			str = append(str, ' ')
		} else if event.Key == keyboard.KeyBackspace {
			if cls > 1 {
				t := str[len(str)-1]
				str = str[:len(str)-1]
				len := utils.RealLength(string(t))
				cls -= len
				utils.SetPos(row, cls)
				fmt.Printf("%*c", len, ' ')
				utils.SetPos(row, cls)
			}
		} else if event.Key == keyboard.KeyEnter {
			fmt.Printf("\n")
			break
		} else {
			utils.SetPos(row, cls)
			c := string(event.Rune)
			fmt.Print(c)
			cls += utils.RealLength(c)
			str = append(str, event.Rune)
		}

		c.PrintMutex.Unlock()
	}
	return string(str)
}

func printTail(c *models.TcpClient) {
	c.PrintMutex.Lock()

	utils.SetPos(7, 1)
	fmt.Print(view.PlayGameViewTail)
	fmt.Print("请输入: ")

	c.PrintMutex.Unlock()
}

func printStopCard(c *models.TcpClient) {
	c.PrintMutex.Lock()

	utils.SetPos(9, 1)
	fmt.Println("停牌")

	c.PrintMutex.Unlock()
}

func PlayGame(c *models.TcpClient) error {
	utils.Cle()
	printHead(c)

	keysEvents, err := keyboard.GetKeys(15)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	myCards := []string{}

	go pullMessage(c)
	for {
		if c.ExitFlag {
			c.ExitFlag = false
			return nil
		}

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

		err := printGame(c, myCards)
		if err != nil {
			exitRoom(c, 1)
			return err
		}
		printMessage(c)

		stopFlag := false
		if point := utils.CalcPoint(myCards); point >= 21 {
			stopFlag = true
		}

		if !stopFlag {
			printTail(c)
			opt := input(c, keysEvents, 8, 9)
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
				err := chat(c, keysEvents)
				if err != nil {
					exitRoom(c, 1)
					return err
				}
				printMessage(c)
			} else if opt == "3" {
				exitRoom(c, 2)
				return nil
			} else {
				continue
			}
		}

		if stopFlag {
			printStopCard(c)
			point := utils.CalcPoint(myCards)
			err := waitResult(c, point)
			if err != nil {
				exitRoom(c, 1)
				return err
			}
			jud := askContinue(c, keysEvents)
			if !jud {
				exitRoom(c, 1)
				return nil
			}
			myCards = []string{}
		}
	}
}
