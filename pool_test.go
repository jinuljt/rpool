package rpool

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

var r = rand.New(rand.NewSource(99))

func sleep(rp *RPool, index int) {
	time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
	fmt.Println("sleep index:", index)
	rp.Done()
}

func TestRPool(t *testing.T) {
	rp := NewRPool(100)
	for i := 0; i < 1000; i++ {
		rp.Add()
		go sleep(rp, i)
	}
	time.Sleep(1 * time.Second)
	rp.Wait()
}
