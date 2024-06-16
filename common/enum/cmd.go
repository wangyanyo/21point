package enum

import "runtime"

type Command string

const (
	HeartPacket      Command = "Heart"      //心跳
	LoginPacket      Command = "Login"      //登录
	RegisterPacket   Command = "Register"   //注册
	GetScorePacket   Command = "GetScore"   //获取分数
	SearchPacket     Command = "Search"     //搜索玩家信息
	RankListPactet   Command = "RankList"   //查看积分排行榜
	UserCountPacket  Command = "UserCount"  //玩家数量
	MatchPacket      Command = "Match"      //匹配
	MatchOffPacket   Command = "MatchOff"   //取消匹配
	EnterRoomPacket  Command = "EnterRoom"  //进入房间
	InitCardPacket   Command = "InitCard"   //洗牌
	AskCardsPactet   Command = "AskCards"   //要牌
	StopCardsPactet  Command = "StopCards"  //停牌
	ExitRoomPacket   Command = "ExitRoom"   //退出房间
	GameResultPacket Command = "GameResult" //游戏结果
	ChatPacket       Command = "Chat"       //聊天包
	GameSayPacket    Command = "GameSay"    //发起聊天
	ContinuePacket   Command = "Continue"   //继续游戏
	OutGamePacket    Command = "OutGame"    //离开游戏
)

const (
	SYS_TYPE = runtime.GOOS
)

var StateMap = map[int]string{
	0: "未开始",
	1: "游戏中",
}
