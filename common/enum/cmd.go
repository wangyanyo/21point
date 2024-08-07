package enum

import "runtime"

type Command string

const (
	LoginPacket       Command = "Login"       //登录
	RegisterPacket    Command = "Register"    //注册
	GetScorePacket    Command = "GetScore"    //获取分数
	SearchPacket      Command = "Search"      //搜索玩家排名分数
	RankListPactet    Command = "RankList"    //查看积分排行榜
	UserCountPacket   Command = "UserCount"   //玩家数量
	MatchPacket       Command = "Match"       //匹配
	MatchOffPacket    Command = "MatchOff"    //取消匹配
	AskCardsPactet    Command = "AskCards"    //要牌
	GameResultPacket  Command = "GameResult"  //游戏结果
	ChatPacket        Command = "Chat"        //发送聊天
	PullMessagePacket Command = "PullMessage" //拉模式获取对方聊天
	ExitRoomPacket    Command = "ExitRoom"    //退出房间
)

const (
	SYS_TYPE = runtime.GOOS
)
