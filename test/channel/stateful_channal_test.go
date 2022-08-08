package channel

import (
	"fmt"
	"math/rand"
	"runtime"
	atomic "sync/atomic"
	"testing"
	"time"
)

type ReadOp struct {
	key  int
	resp chan int
}

type WriteOp struct {
	key  int
	val  int
	resp chan bool
}

func TestStatefulChannel(t *testing.T) {
	var ops int64
	reads := make(chan *ReadOp)
	writes := make(chan *WriteOp)

	go func() {
		state := make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	for r := 0; r < 100; r++ {
		go func() {
			for {
				read := &ReadOp{
					key:  rand.Intn(5),
					resp: make(chan int),
				}
				reads <- read
				<-read.resp

				atomic.AddInt64(&ops, 1)

				runtime.Gosched()
			}
		}()
	}

	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := &WriteOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool),
				}
				writes <- write
				<-write.resp

				atomic.AddInt64(&ops, 1)

				runtime.Gosched()
			}
		}()

		time.Sleep(time.Second)

		opsFinal := atomic.LoadInt64(&ops)
		fmt.Println("ops:", opsFinal)
	}
}
