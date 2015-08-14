package rpool

import "time"

type RPool struct {
	ch chan bool
}

func NewRPool(max int) *RPool {
	return &RPool{
		ch: make(chan bool, max),
	}
}

func (rp *RPool) Add() {
	rp.ch <- false
}

func (rp *RPool) Done() {
	<-rp.ch
}

func (rp *RPool) Wait() {
	for {
		if len(rp.ch) == 0 {
			return
		}
		time.Sleep(1 * time.Millisecond)
	}
}
