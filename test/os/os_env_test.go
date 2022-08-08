package os

import (
	"fmt"
	"os"
	"testing"
)

func TestOs(t *testing.T) {
	for _, e := range os.Environ() {
		//pair := strings.Split(e, "=")
		fmt.Println(e)
	}
}
