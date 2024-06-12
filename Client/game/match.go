package game

import (
	"fmt"
	"log"
	"time"

	"github.com/eiannone/keyboard"
	"github.com/wangyanyo/21point/Client/models"
	"github.com/wangyanyo/21point/Client/view"
	"github.com/wangyanyo/21point/common/entity"
	"github.com/wangyanyo/21point/common/enum"
	"github.com/wangyanyo/21point/common/myerror"
	"github.com/wangyanyo/21point/common/utils"
)

func checkQuit(ch chan int) {
	keysEvents, err := keyboard.GetKeys(10)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	for {
		select {
		case event := <-keysEvents:
			if event.Rune == 'q' {
				ch <- 1
				return
			}
		case <-ch:
			return
		}

	}
}

func Match(c *models.TcpClient) error {
	utils.Cle()
	fmt.Print(view.MatchView)

	if _, err := c.Send(entity.NewTransfeData(enum.MatchPacket, c.Token, 0)); err != nil {
		myerror.Reconnect(err)
		return err
	}

	ch := make(chan int)
	go checkQuit(ch)

	select {
	case roomInfo := <-c.CmdChan:
		ch <- 1
		log.Println("匹配成功", roomInfo)
		if roomInfo.Cmd == enum.MatchPacket {
			if roomInfo.Code == 1 {
				roomId := roomInfo.Data.(int)
				if err := PlayGame(c, roomId); err != nil {
					return err
				}
				return nil
			} else {
				utils.PrintMessage("匹配出错")
				return myerror.New("匹配出错")
			}
		} else {
			utils.PrintMessage("匹配包异常")
			return myerror.New("匹配包异常")
		}

	case <-ch:
		retry := 10
		for retry > 0 {
			if _, err := c.Send(entity.NewTransfeData(enum.MatchOffPacket, c.Token, 0)); err != nil {
				myerror.Reconnect(err)
				continue
			}
			matchOffInfo := <-c.CmdChan
			if matchOffInfo.Cmd != enum.MatchOffPacket {
				if matchOffInfo.Code == 1 {

				} else {

				}
			} else {

			}
		}
		if retry == 0 {

		}
		utils.PrintMessage("退出匹配成功！")
		time.Sleep(1 * time.Second)
		return nil
	}
}
