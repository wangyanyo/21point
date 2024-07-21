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
	mutex     sync.Mutex
	TimeFlag1 time.Time
	TimeFlag2 time.Time
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
	}
	return r.Player1
}

func (r *Room) GetOtherPoint(name string) int {
	if name == r.Player1 {
		return r.Point2
	}
	return r.Point1
}

func (r *Room) JudgeOtherTimeOut(name string) bool {
	if name == r.Player1 {
		return time.Since(r.TimeFlag1).Seconds() > 10.0
	}
	return time.Since(r.TimeFlag2).Seconds() > 10.0
}

func (r *Room) SetTimeFlag(name string) {
	if name == r.Player1 {
		r.TimeFlag1 = time.Now()
	}
	r.TimeFlag2 = time.Now()
}

func (r *Room) SetPoint(name string, point int) {
	if name == r.Player1 {
		r.Point1 = point
	}
	r.Point2 = point
}
