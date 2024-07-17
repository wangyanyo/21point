package game

import (
	"fmt"
	"log"

	"github.com/wangyanyo/21point/Client/models"
	"github.com/wangyanyo/21point/Client/ral"
	"github.com/wangyanyo/21point/Client/view"
	"github.com/wangyanyo/21point/common/entity"
	"github.com/wangyanyo/21point/common/enum"
	"github.com/wangyanyo/21point/common/utils"
)

func EnterRoom(c *models.TcpClient, roomID int) error {
	utils.Cle()
	fmt.Print(view.EnterRoomView)

	req := &entity.TransfeData{
		Cmd:    enum.EnterRoomPacket,
		Token:  c.Token,
		RoomID: c.RoomID,
	}
	if _, err := ral.Ral(c, req); err != nil {
		return err
	}

	fmt.Println("进入房间成功！")
	log.Println("进入房间", c.Token, roomID)
	if err := PlayGame(c); err != nil {
		return err
	}
	return nil
}
