package tcpsrc

import (
	"context"
	"log"

	"github.com/wangyanyo/21point/Server/controller"
	"github.com/wangyanyo/21point/Server/models"
	"github.com/wangyanyo/21point/common/entity"
	"github.com/wangyanyo/21point/common/enum"
)

func Router(ctx context.Context, req *entity.TransfeData, client *models.ClientUser) {
	var res *entity.TransfeData
	switch req.Cmd {
	case enum.HeartPacket:
		controller.HeartHandle(ctx, req, client)

	case enum.RegisterPacket:

	case enum.LoginPacket:

	}

	if req.Cmd != enum.HeartPacket {
		_, err := client.Connection.Write(res.Byte())
		if err != nil {
			log.Println("断开连接", client.Connection.RemoteAddr(), client.Connection, client.Token)
			client.CloseChan <- struct{}{}
			return
		}
	}
}
