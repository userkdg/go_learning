package time

import (
	"fmt"
	"testing"
	"time"
)

var p = fmt.Println

func TestTime(t *testing.T) {
	now := time.Now()
	p(now)
	p(now.Format("2006-01-02 15:04:05"))

	then := time.Date(2022, 8, 8, 14, 00, 00, 00000000, time.Local)
	p(then)

	diff := now.Sub(then)
	p(diff)

	p(then.Add(diff))
	p(then.Add(-diff))
}
