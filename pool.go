package rpool

import (
	"sync"
	"time"
)

type RPool struct {
	max     int
	current int
	mutex   sync.Mutex
}

func NewRPool(max int) *RPool {
	return &RPool{max: max}
}

func (rp *RPool) Add() {
	for {
		rp.mutex.Lock()
		if rp.current < rp.max {
			rp.current += 1
			rp.mutex.Unlock()
			break
		}
		rp.mutex.Unlock()
		time.Sleep(1 * time.Nanosecond)
	}
}

func (rp *RPool) Done() {
	rp.mutex.Lock()
	rp.current -= 1
	rp.mutex.Unlock()
}

func (rp *RPool) Wait() {
	for {
		rp.mutex.Lock()
		if rp.current == 0 {
			rp.mutex.Unlock()
			break
		}
		rp.mutex.Unlock()
		time.Sleep(1 * time.Nanosecond)
	}
}
