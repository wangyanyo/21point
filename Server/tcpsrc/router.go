package tcpsrc

import (
	"github.com/wangyanyo/21point/Sever/controller"
	"github.com/wangyanyo/21point/common/entity"
	"github.com/wangyanyo/21point/common/enum"
)

func Router(req *entity.TransfeData) {
	switch req.Cmd {
	case enum.HeartPacket:
		controller.HeartHandle()

	case enum.RegisterPacket:

	case enum.LoginPacket:

	}
}
