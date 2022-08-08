package channel

import (
	"fmt"
	"testing"
	"time"
)

func TestChannelSync(t *testing.T) {
	done := make(chan bool, 1)
	go worker(done)

	<-done
}

func worker(done chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}
