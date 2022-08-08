package os

import (
	"fmt"
	"os"
	"testing"
)

// go build xx
// ./os_exit_test
// echo $?
// 3
func TestOsExit(t *testing.T) {
	defer fmt.Println("!")
	os.Exit(3)
}
