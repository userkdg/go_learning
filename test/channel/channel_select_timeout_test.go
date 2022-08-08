package channel

import (
	"testing"
	"time"
)

func TestTimeout(t *testing.T) {
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "result 1"
	}()
	select {
	case m1 := <-c1:
		println(m1)
	case <-time.After(time.Second * 1):
		println("timeout 1")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 1)
		c2 <- "result 2"
	}()
	select {
	case m2 := <-c2:
		println(m2)
	case <-time.After(time.Second * 2):
		println("timeout 2")
	}

}
