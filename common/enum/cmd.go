package enum

import "runtime"

type Command string

const (
	LoginPacket      Command = "Login"      //登录
	RegisterPacket   Command = "Register"   //注册
	GetScorePacket   Command = "GetScore"   //获取分数
	SearchPacket     Command = "Search"     //搜索玩家信息
	RankListPactet   Command = "RankList"   //查看积分排行榜
	UserCountPacket  Command = "UserCount"  //玩家数量
	MatchPacket      Command = "Match"      //匹配
	MatchOffPacket   Command = "MatchOff"   //取消匹配
	InitCardPacket   Command = "InitCard"   //洗牌
	AskCardsPactet   Command = "AskCards"   //要牌
	StopCardsPactet  Command = "StopCards"  //停牌
	GameResultPacket Command = "GameResult" //游戏结果
	ChatPacket       Command = "Chat"       //聊天包
	GameSayPacket    Command = "GameSay"    //发起聊天
	ExitRoomPacket   Command = "ExitRoom"   //退出房间
)

const (
	SYS_TYPE = runtime.GOOS
)
