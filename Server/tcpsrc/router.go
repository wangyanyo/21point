package tcpsrc

import (
	"context"
	"log"

	"github.com/wangyanyo/21point/Server/models"
	"github.com/wangyanyo/21point/Server/service"
	"github.com/wangyanyo/21point/common/entity"
	"github.com/wangyanyo/21point/common/enum"
)

func Router(ctx context.Context, req *entity.TransfeData, client *models.ClientUser) {
	var resp *entity.TransfeData
	ser := service.GetServer()
	switch req.Cmd {
	case enum.RegisterPacket: //注册
		resp = ser.RegisterHandler(ctx, req)

	case enum.LoginPacket: //登录
		resp = ser.LoginHandler(ctx, req)

	case enum.GetScorePacket: //获取分数
		resp = ser.GetSocreHandler(ctx, req)

	case enum.SearchPacket: //搜索玩家信息

	case enum.RankListPactet: //查看积分排行榜

	case enum.UserCountPacket: //玩家数量

	case enum.MatchPacket: //匹配

	case enum.MatchOffPacket: //取消匹配

	case enum.InitCardPacket: //洗牌

	case enum.AskCardsPactet: //要牌

	case enum.StopCardsPactet: //停牌

	case enum.GameResultPacket: //游戏结果

	case enum.ChatPacket: //聊天包

	case enum.GameSayPacket: //发起聊天

	case enum.ExitRoomPacket: //退出房间

	}

	resp.Cmd = req.Cmd
	_, err := client.Connection.Write(resp.Byte())
	if err != nil {
		log.Println("[写数据出错]", client.Connection.RemoteAddr().String(), client.Connection)
	}
}
