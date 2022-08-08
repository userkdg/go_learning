package channel

import (
	"fmt"
	"testing"
)

func TestRangeChannel(t *testing.T) {
	queue := make(chan string, 2)

	queue <- "1"
	queue <- "2"
	queue <- "3"

	for elem := range queue {
		fmt.Println(elem)
	}
	close(queue)
}
