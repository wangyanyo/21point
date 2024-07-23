package game

import (
	"fmt"

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

func Read(c *models.TcpClient, CmdChan chan *entity.TransfeData) {
	data, _ := ral.Read(c)
	CmdChan <- data
}

func Match(c *models.TcpClient) error {
	utils.Cle()
	fmt.Print(view.MatchView)

	req := &entity.TransfeData{
		Cmd:    enum.MatchPacket,
		Token:  c.Token,
		RoomID: c.RoomID,
	}
	if err := ral.SendRequest(c, req); err != nil {
		return err
	}

	ch := make(chan struct{})
	CmdChan := make(chan *entity.TransfeData)
	go checkQuit(ch)
	go Read(c, CmdChan)

	flag := false
	select {
	case roomInfo := <-CmdChan:
		if !flag {
			ch <- struct{}{}
		}
		err := ral.CheckPacket(roomInfo, req)
		if err != nil {
			myerror.PrintError(err)
			return err
		}
		utils.PrintMessage("匹配成功，正在进入房间...")
		c.RoomID = roomInfo.Data.(int)
		if err := PlayGame(c); err != nil {
			return err
		}
		return nil

	case <-ch:
		flag = true
		req := &entity.TransfeData{
			Cmd:    enum.MatchOffPacket,
			Token:  c.Token,
			RoomID: c.RoomID,
		}
		if err := ral.SendRequest(c, req); err != nil {
			return err
		}
	}
	return nil
}
