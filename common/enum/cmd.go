package enum

import "runtime"

type Command string

const (
	HeartPacket     Command = "Heart_Packet"     //心跳
	LoginPacket     Command = "Login_Packet"     //登录
	RegisterPacket  Command = "Register_Packet"  //注册
	RankListPactet  Command = "RankList_Packet"  //查看积分排行榜
	MatchPacket     Command = "Match_Packet"     //匹配
	MatchOffPacket  Command = "MatchOff_Packet"  //取消匹配
	AskCardsPactet  Command = "AskCards_Pactet"  //要牌
	StopCardsPactet Command = "StopCards_Pactet" //停牌
	ChatPacket      Command = "Chat_Packet"      //聊天
	GameSayPacket   Command = "GameSay_Packet"   //发起聊天
	ContinuePacket  Command = "Continue_Packet"  //继续游戏
	OutGamePacket   Command = "OutGame_Packet"   //离开游戏
	WinGamePacket   Command = "WinGame_Packet"   //21点获胜
	LoseGamePacket  Command = "LoseGame_Packet"  //爆牌失败

)

const (
	SYS_TYPE = runtime.GOOS
)

var StateMap = map[int]string{
	0: "未开始",
	1: "游戏中",
}
