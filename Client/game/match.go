package game

import (
	"fmt"

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
		myerror.Reconnect(err, 1)
		return err
	}

	ch := make(chan int)
	go checkQuit(ch)

	select {
	case roomInfo := <-c.CmdChan:
		ch <- 1
		err := myerror.CheckPacket(roomInfo, enum.MatchPacket)
		if err != nil {
			return err
		}
		roomId := roomInfo.Data.(int)
		if err := EnterRoom(c, roomId); err != nil {
			return err
		}
		return nil

	case <-ch:
		if err := MatchOff(c); err != nil {
			return err
		}
		return nil
	}
}
