package runtime

import (
	"runtime"
	"testing"
)

func Test1(t *testing.T) {
	runtime.Gosched()
}
