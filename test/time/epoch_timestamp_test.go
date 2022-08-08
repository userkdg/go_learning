package time

import (
	"fmt"
	"testing"
	"time"
)

// 秒数，毫秒数，或者微
func TestEpoch(t *testing.T) {
	now := time.Now()
	secs := now.Unix()
	milli := now.UnixMilli()
	nano := now.UnixNano()
	fmt.Println(now)

	fmt.Println(secs)
	fmt.Println(nano / 1000000) // 毫秒：手动计算从纳秒转化一下
	fmt.Println(milli)          // 毫秒
	fmt.Println(nano)

	fmt.Println(time.Unix(secs, 0))
	fmt.Println(time.Unix(0, nano))
}
