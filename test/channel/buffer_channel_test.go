package channel

import (
	"fmt"
	"testing"
)

func TestBuffer(t *testing.T) {
	msg := make(chan string, 2)

	msg <- "1"
	msg <- "2"

	close(msg)

	fmt.Println(<-msg)

	if m, ok := <-msg; ok {
		fmt.Println(m)
	} else {
		fmt.Println("closed")
	}

	if m, ok := <-msg; ok {
		fmt.Println(m)
	} else {
		fmt.Println("closed")
	}
}
