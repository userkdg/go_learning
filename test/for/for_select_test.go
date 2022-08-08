package _for

import (
	"fmt"
	"testing"
	"time"
)

// 使用break lable 和 goto lable 都能跳出for循环；不同之处在于：break标签只能用于for循环，且标签位于for循环前面，goto是指跳转到指定标签处
func TestForSelectReturn(t *testing.T) {
	i := 0
	for {
		select {
		case <-time.After(time.Second * time.Duration(2)):
			i++
			if i == 5 {
				fmt.Println("跳出for循环")
				return
			}
		}
		fmt.Println("for循环内 i=", i)
	}
	// 不可达！由于for select结束用return，可以把for select走独立协程
	fmt.Println("for循环外")
}

func TestForSelectReturnReachable(t *testing.T) {
	done := make(chan bool)
	go func() {
		i := 0
		for {
			select {
			case <-time.After(time.Second * time.Duration(2)):
				i++
				if i == 5 {
					fmt.Println("跳出for循环")
					done <- true
					return // 跳出for
				}
			}
			fmt.Println("for循环内 i=", i)
		}
	}()
	<-done
	// 不可达！由于for select结束用return，可以把for select走独立协程
	fmt.Println("for循环外")
}

func TestForSelectBreak(t *testing.T) {
	i := 0
Loop:
	for {
		select {
		case <-time.After(time.Second * time.Duration(2)):
			i++
			if i == 5 {
				fmt.Println("跳出for循环")
				break Loop
			}
		}
		fmt.Println("for循环内 i=", i)
	}

	fmt.Println("for循环外")
}

func TestForSelectGoto(t *testing.T) {
	i := 0
	for {
		select {
		case <-time.After(time.Second * time.Duration(2)):
			i++
			if i == 5 {
				fmt.Println("跳出for循环")
				goto Loop
			}
		}
		fmt.Println("for循环内 i=", i)
	}
Loop:
	fmt.Println("for循环外")
}
