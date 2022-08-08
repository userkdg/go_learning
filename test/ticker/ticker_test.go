package ticker

import (
	"testing"
	"time"
)

func TestTicker1(t *testing.T) {
	t1 := time.NewTicker(time.Millisecond * 500)
	go func() {
		for c := range t1.C {
			t.Logf("tick at %s", c)
		}
	}()
	time.Sleep(time.Millisecond * 1600)
	t1.Stop()
	t.Log("ticker stopped")
}
