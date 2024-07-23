package models

import (
	"crypto/rand"
	"math/big"
	"sync"
	"time"
)

var RoomExitMsgMap map[string]string

type Room struct {
	Player1   string
	Player2   string
	Point1    int
	Point2    int
	Cards     []int
	Flag      bool
	TimeFlag1 time.Time
	TimeFlag2 time.Time
	mutex     sync.Mutex
	initFlag  int
	MsgSet1   []string
	MsgSet2   []string
}

func (r *Room) Exist(name string) bool {
	return name == r.Player1 || name == r.Player2
}

func (r *Room) Init() {
	r.Cards = make([]int, 13)
	for i := 0; i < 13; i++ {
		r.Cards[i] = 4
	}
	r.Point1 = 0
	r.Point2 = 0
	r.TimeFlag1 = time.Now()
	r.TimeFlag2 = time.Now()
	r.initFlag = 0
}

func (r *Room) GetCard() int {
	for {
		n, _ := rand.Int(rand.Reader, big.NewInt(1e9+7))
		t := int(n.Int64()) % 13
		r.mutex.Lock()
		if r.Cards[t] == 0 {
			r.mutex.Unlock()
			continue
		}
		r.Cards[t]--
		r.mutex.Unlock()
		return t
	}
}

func (r *Room) GetOtherPlayer(name string) string {
	if name == r.Player1 {
		return r.Player2
	} else {
		return r.Player1
	}
}

func (r *Room) GetOtherPoint(name string) int {
	if name == r.Player1 {
		return r.Point2
	} else {
		return r.Point1
	}
}

func (r *Room) JudgeOtherTimeOut(name string) bool {
	if name == r.Player1 {
		return time.Since(r.TimeFlag2).Seconds() > 10.0
	} else {
		return time.Since(r.TimeFlag1).Seconds() > 10.0
	}
}

func (r *Room) SetTimeFlag(name string) {
	if name == r.Player1 {
		r.TimeFlag1 = time.Now()
	} else {
		r.TimeFlag2 = time.Now()
	}
}

func (r *Room) SetPoint(name string, point int) {
	if name == r.Player1 {
		r.Point1 = point
	} else {
		r.Point2 = point
	}
}

func (r *Room) CallInit() {
	r.mutex.Lock()
	r.initFlag++
	if r.initFlag == 2 {
		r.Init()
	}
	r.mutex.Unlock()
}

func (r *Room) AddMsg(name string, msg string) {
	if name == r.Player1 {
		r.MsgSet1 = append(r.MsgSet1, msg)
	} else {
		r.MsgSet2 = append(r.MsgSet2, msg)
	}
}

func (r *Room) GetOtherMsg(name string, count int) []string {
	if name == r.Player1 {
		return r.MsgSet2[count:]
	} else {
		return r.MsgSet1[count:]
	}
}
