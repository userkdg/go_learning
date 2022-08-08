package channel

import "testing"

func TestNonBlocking(t *testing.T) {
	b1 := make(chan string)
	select {
	case msg1 := <-b1:
		println(msg1)
	default:
		println("no msg received!")
	}
}
