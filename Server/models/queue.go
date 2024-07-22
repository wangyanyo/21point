package models

import (
	"errors"
	"sync"
)

type Queue struct {
	a     []int
	l     int
	r     int
	cnt   int
	size  int
	mutex sync.Mutex
}

func (q *Queue) Init(x int) {
	q.size = x
	q.a = make([]int, q.size)
}

func (q *Queue) Push(x int) error {
	q.mutex.Lock()
	q.cnt++
	if q.cnt > q.size {
		q.mutex.Unlock()
		panic(errors.New("queue overflow"))
	}
	q.r = (q.r + 1) % q.size
	q.a[q.r] = x
	q.mutex.Unlock()
	return nil
}

func (q *Queue) Pop() (int, error) {
	q.mutex.Lock()
	q.cnt--
	if q.cnt < 0 {
		q.mutex.Unlock()
		panic(errors.New("queue is empty"))
	}
	t := q.a[q.l]
	q.l = (q.l + 1) % q.size
	q.mutex.Unlock()
	return t, nil
}

func (q *Queue) Empty() bool {
	q.mutex.Lock()
	flag := bool(q.cnt == 0)
	q.mutex.Unlock()
	return flag
}
