package channel

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	msg := make(chan string)

	go func() {
		msg <- "ping"
	}()

	m := <-msg
	fmt.Println(m)
}
