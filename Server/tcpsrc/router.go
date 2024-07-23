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
		resp = ser.SearchHandler(ctx, req)

	case enum.RankListPactet: //查看积分排行榜
		resp = ser.RankListHandler(ctx, req)

	case enum.UserCountPacket: //玩家数量
		resp = ser.UserCountHandler(ctx, req)

	case enum.MatchPacket: //匹配
		resp = ser.MatchHandler(ctx, req)

	case enum.MatchOffPacket: //取消匹配
		resp = ser.MatchOffHandler(ctx, req)

	case enum.AskCardsPactet: //要牌
		resp = ser.AskCardsHandler(ctx, req)

	case enum.GameResultPacket: //游戏结果
		resp = ser.GameResultHandler(ctx, req)

	case enum.ChatPacket: //发送聊天
		resp = ser.ChatHandler(ctx, req)

	case enum.PullMessagePacket: //拉模式获取对方聊天
		resp = ser.PullMessageHandler(ctx, req)

	case enum.ExitRoomPacket: //退出房间
		resp = ser.ExitRoomHandler(ctx, req)

	}

	resp.Cmd = req.Cmd
	_, err := client.Connection.Write(resp.Byte())
	if err != nil {
		log.Println("[写数据出错]", client.Connection.RemoteAddr().String(), client.Connection)
	}
}
