package channel

import (
	"fmt"
	"testing"
	"time"
)

func TestSelect(t *testing.T) {
	c1 := make(chan string)
	c2 := make(chan string)
	done := make(chan bool)

	go c1Func(c1)
	go c2Func(c2)
	go doneFunc(done)

	// golang跳出for select 循环
	//通常在for循环中，使用break可以跳出循环，但是注意在go语言中，for select配合时，break并不能跳出循环。
	for {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		case d := <-done:
			fmt.Println("finished", d)
			//return
			goto Loop
		}
	}
Loop:
	fmt.Println("done!")
}

func doneFunc(done chan bool) {
	//time.Sleep(time.Second * 3)
	time.Sleep(time.Second * 1)
	done <- true
}

func c1Func(c1 chan string) {
	time.Sleep(time.Second * 2)
	c1 <- "one1"
}

func c2Func(c2 chan string) {
	time.Sleep(time.Second * 1)
	c2 <- "one2"
}
