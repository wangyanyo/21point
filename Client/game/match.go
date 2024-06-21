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

func checkQuit(ch chan struct{}) {
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
				ch <- struct{}{}
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
		utils.PrintMessage("断线重连...")
		ral.Connect()
		time.Sleep(1 * time.Second)
		return err
	}

	ch := make(chan struct{})
	go checkQuit(ch)

	flag := false
	select {
	case roomInfo := <-c.CmdChan:
		if !flag {
			ch <- struct{}{}
		}
		err := myerror.CheckPacket(roomInfo, enum.MatchPacket)
		if err != nil {
			if flag && err.Error() == (string(enum.MatchPacket)+"--RequestError") {
				utils.PrintMessage("退出匹配成功！")
				return nil
			} else {
				utils.PrintMessage("匹配失败！")
				return err
			}
		}
		if flag {
			utils.PrintMessage("退出匹配失败，正在进入房间...")
		} else {
			utils.PrintMessage("匹配成功，正在进入房间...")
		}
		roomId := roomInfo.Data.(int)
		if err := EnterRoom(c, roomId); err != nil {
			return err
		}
		return nil

	case <-ch:
		flag = true
		if err := ral.SendRequest(c, enum.MatchOffPacket, c.Token, ""); err != nil {
			return err
		}
	}
	return nil
}
