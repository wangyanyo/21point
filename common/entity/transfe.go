package entity

import (
	"bytes"
	"encoding/gob"
	"log"
	"time"

	"github.com/wangyanyo/21point/common/enum"
)

func init() {
	gob.Register(User{})
	gob.Register(ChatData{})
	gob.Register(UserScoreInfo{})
}

// 任何类型都可以实现空接口，因此可以用空接口代表任何一个类型，类似所有类型的基类
// Code返回是否有异常，Data返回想要的结果，当然也要具体情况具体分析
type TransfeData struct {
	Cmd       enum.Command //指令
	Timestamp int64
	Token     string      //识别客户端身份
	Data      interface{} //传输的数据
	Msg       string      //传输消息
	Code      int         //传输Code
	RoomID    int         //房间号
}

// 数据压缩
func (t *TransfeData) Byte() []byte {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(t)
	if err != nil {
		log.Fatal("encode error", err)
	}
	return buffer.Bytes()
}

// 数据压缩的另一种方法
func NewTransfeData(cmd enum.Command, token string, data interface{}) []byte {
	tra := &TransfeData{
		Cmd:       cmd,
		Timestamp: time.Now().Unix(),
		Token:     token,
		Data:      data,
	}
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(tra)
	if err != nil {
		log.Fatal("encode error", err)
	}
	return buffer.Bytes()
}

func TransfeDataDecoder(data []byte) *TransfeData {
	buffer := bytes.NewBuffer(data)
	decoder := gob.NewDecoder(buffer)
	tra := &TransfeData{}
	err := decoder.Decode(&tra)
	if err != nil {
		log.Fatal("decode error", err)
	}
	return tra
}

// 用户
type User struct {
	Name     string
	Password string
}

// 聊天
type ChatData struct {
	Flag int
	Msg  string
}

// 玩家分数信息
type UserScoreInfo struct {
	Rank  int
	Name  string
	Score int
}
