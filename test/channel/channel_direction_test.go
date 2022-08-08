package channel

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}

//pings接收不发送、pongs接收：func pong(pings chan<- string, pongs chan<- string) msg := <- pings 会编译错误
//ping发送、pong接收
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func ping(pings chan<- string, msg string) {
	pings <- msg
}
