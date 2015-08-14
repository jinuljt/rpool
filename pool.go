package rpool

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
	count := 0
	max := cap(rp.ch)

	for {
		rp.ch <- true
		count += 1
		if max == count {
			return
		}

	}
}
