package timer

import (
	"testing"
	"time"
)

func TestTimer1(t *testing.T) {
	t1 := time.NewTimer(time.Second * 2)
	<-t1.C
	t.Log("timer 1 expired")

	t2 := time.NewTimer(time.Second)
	go func() {
		<-t2.C
		t.Log("timer 2 expired")
	}()
	stop := t2.Stop()
	if stop {
		t.Log("timer 2 stopped")
	}
}
